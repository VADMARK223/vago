package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

type Answer struct {
	Text    string `json:"text"`
	Correct bool   `json:"correct"`
}

type Question struct {
	TopicID int      `json:"topic_id"`
	Text    string   `json:"text"`
	Answers []Answer `json:"answers"`
}

const (
	input  = "data/questions.json"
	output = "db/04_questions.sql"
)

func main() {
	log.Println("Start generation.")
	data, err := os.ReadFile(input)
	if err != nil {
		log.Fatal(err)
	}

	var questions []Question
	if err := json.Unmarshal(data, &questions); err != nil {
		log.Fatal(err)
	}

	out, err := os.Create(output)
	if err != nil {
		log.Fatal(err)
	}
	defer func(out *os.File) {
		err := out.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(out)

	_, _ = fmt.Fprintln(out, "TRUNCATE TABLE questions RESTART IDENTITY CASCADE;")
	_, _ = fmt.Fprintln(out, "BEGIN;")

	questionID := 1
	for _, q := range questions {
		_, _ = fmt.Fprintf(out, "INSERT INTO questions (id, topic_id, text) VALUES (%d, %d, '%s');\n", questionID, q.TopicID, escape(q.Text))

		for _, ans := range q.Answers {
			correct := "false"
			if ans.Correct {
				correct = "true"
			}

			_, _ = fmt.Fprintf(out, "INSERT INTO answers (question_id, text, is_correct) VALUES (%d, '%s', %s);\n", questionID, escape(ans.Text), correct)
		}

		_, _ = fmt.Fprintln(out)
		questionID++
	}

	_, _ = fmt.Fprintln(out, "COMMIT;")
	log.Println("Generated in " + output)
	log.Println("Questions loaded", len(questions))

}

func escape(s string) string {
	return strings.ReplaceAll(s, "'", "''")
}
