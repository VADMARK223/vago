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
	Correct bool   `json:"correct,omitempty"`
}

type Question struct {
	TopicID     int      `json:"topic_id"`
	Text        string   `json:"text"`
	Code        string   `json:"code,omitempty"`
	Explanation string   `json:"explanation,omitempty"`
	Answers     []Answer `json:"answers"`
}

const dataFile = "data/questions.json"

//const dataFileTest = "data/questions_test.json"

func AddQuestion(question Question) error {
	data, err := os.ReadFile(dataFile)
	if err != nil {
		return err
	}

	var questions []Question
	if err := json.Unmarshal(data, &questions); err != nil {
		return err
	}

	questions = append(questions, question)

	out, err := json.MarshalIndent(questions, "", "  ")
	if err != nil {
		return err
	}

	if err := os.WriteFile(dataFile, out, 0644); err != nil {
		return err
	}

	log.Println("Question added.")

	return nil
}

func Run(dsn string) error {
	log.Println("Start generation: ", dsn)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	data, err := os.ReadFile(dataFile)
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
