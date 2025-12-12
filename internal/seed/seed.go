package seed

import (
	"encoding/json"
	"log"
	"os"
	"time"

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

type Topic struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

const (
	dataFileQuestions = "data/questions.json"
	dataFileTopics    = "data/topics.json"
)

func AddQuestion(question Question) error {
	data, err := os.ReadFile(dataFileQuestions)
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

	if err := os.WriteFile(dataFileQuestions, out, 0644); err != nil {
		return err
	}

	log.Println("Question added.")

	return nil
}

func Topics(dsn string) error {
	start := time.Now()
	log.Println("Start topics: ", dsn)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	data, err := os.ReadFile(dataFileTopics)
	if err != nil {
		return err
	}

	var topics []Topic
	if err := json.Unmarshal(data, &topics); err != nil {
		return err
	}

	db.Exec("TRUNCATE TABLE topics RESTART IDENTITY CASCADE")

	if err := db.Create(&topics).Error; err != nil {
		return err
	}

	elapsed := time.Since(start)
	log.Println("Topics seeding complete: ", elapsed)
	return nil
}

func SyncQuestions(dsn string) error {
	start := time.Now()
	log.Println("Start questions: ", dsn)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	data, err := os.ReadFile(dataFileQuestions)
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
	elapsed := time.Since(start)
	log.Println("Questions seeding complete:", elapsed)
	return nil
}
