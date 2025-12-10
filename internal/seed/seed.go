package seed

import (
	"encoding/json"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Answer struct {
	Text    string `json:"text"`
	Correct bool   `json:"correct"`
}

type Question struct {
	TopicID     int      `json:"topic_id"`
	Text        string   `json:"text"`
	Code        string   `json:"code"`
	Explanation string   `json:"explanation"`
	Answers     []Answer `json:"answers"`
}

const (
	input = "data/questions.json"
)

func Run() error {
	log.Println("Start generation.")
	dsn := "postgres://vadmark:5125341@localhost:5432/vagodb?sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	data, err := os.ReadFile(input)
	if err != nil {
		return err
	}

	var questions []Question
	if err := json.Unmarshal(data, &questions); err != nil {
		return err
	}

	db.Exec("TRUNCATE TABLE answers RESTART IDENTITY CASCADE")
	db.Exec("TRUNCATE TABLE questions RESTART IDENTITY CASCADE")

	for _, q := range questions {
		var id int
		errInsert := db.Raw(
			`INSERT INTO questions (topic_id, text, code, explanation) 
             VALUES (?, ?, ?, ?) RETURNING id`,
			q.TopicID, q.Text, q.Code, q.Explanation,
		).Scan(&id).Error

		if errInsert != nil {
			return errInsert
		}

		for _, a := range q.Answers {
			db.Exec(
				`INSERT INTO answers (question_id, text, is_correct) VALUES (?, ?, ?)`,
				id, a.Text, a.Correct,
			)
		}
	}

	log.Println("Seeding complete.")
	return nil
}
