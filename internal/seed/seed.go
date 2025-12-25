package seed

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
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
	batchSize         = 100
	dataFileQuestions = "data/questions.json"
	dataFileTopics    = "data/topics.json"
)

func AddQuestion(ctx context.Context, dsn string, question Question) error {
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

	return SyncQuestions(ctx, dsn)
}

func SyncQuestions(ctx context.Context, dsn string) error {
	start := time.Now()
	fmt.Printf("➡️ \033[93m%s: \033[92m%v\033[0m\n", "Start seed questions", dsn)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	db = db.WithContext(ctx)

	data, err := os.ReadFile(dataFileQuestions)
	if err != nil {
		return err
	}

	var questions []Question
	if err := json.Unmarshal(data, &questions); err != nil {
		return err
	}

	/*if err := truncateTables(db); err != nil {
		return err
	}*/

	for i := 0; i < len(questions); i += batchSize {
		fmt.Printf("➡️ \033[93m%s: \033[92m%v\033[0m\n", "Batch", i)
		if err := ctx.Err(); err != nil {
			return err
		}

		end := min(i+batchSize, len(questions))
		if err := insertBatch(ctx, db, questions[i:end]); err != nil {
			return err
		}
	}

	fmt.Printf("➡️ \033[93m%s: \033[92m%v\033[0m\n", "End seed questions", time.Since(start))
	return nil
}

/*func truncateTables(db *gorm.DB) error {
	if err := db.Exec(
		"TRUNCATE TABLE answers RESTART IDENTITY CASCADE",
	).Error; err != nil {
		return err
	}

	if err := db.Exec(
		"TRUNCATE TABLE questions RESTART IDENTITY CASCADE",
	).Error; err != nil {
		return err
	}

	return nil
}*/

type upsertResult struct {
	ID       int
	Inserted bool
}

var res upsertResult

func insertBatch(ctx context.Context, db *gorm.DB, batch []Question) error {
	tx := db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	// если что-то пойдёт не так — откатываем
	committed := false
	defer func() {
		if !committed {
			tx.Rollback()
		}
	}()

	for _, q := range batch {
		// 🔴 точка отмены ВНУТРИ батча
		if err := ctx.Err(); err != nil {
			return err
		}

		var upsertRes upsertResult
		if err := tx.Raw(
			`INSERT INTO questions (topic_id, text, code, explanation, content_hash)
			 VALUES (?, ?, ?, ?, ?)
			 ON CONFLICT (content_hash) DO UPDATE
			 SET explanation = EXCLUDED.explanation
			 RETURNING id, (xmax = 0) AS inserted`,
			q.TopicID, q.Text, q.Code, q.Explanation, questionHash(q.TopicID, q.Text, q.Code),
		).Scan(&upsertRes).Error; err != nil {
			return err
		}

		if !upsertRes.Inserted {
			continue
		}

		for _, a := range q.Answers {
			if err := tx.Exec(
				`INSERT INTO answers (question_id, text, is_correct)
				 VALUES (?, ?, ?)`,
				upsertRes.ID, a.Text, a.Correct,
			).Error; err != nil {
				return err
			}
		}
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	committed = true
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

func questionHash(topicID int, text, code string) string {
	h := sha256.New()

	h.Write([]byte(strconv.Itoa(topicID)))
	h.Write([]byte{'|'})
	h.Write([]byte(text))
	h.Write([]byte{'|'})
	h.Write([]byte(code))

	return hex.EncodeToString(h.Sum(nil))
}
