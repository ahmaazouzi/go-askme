package repository

import (
	"time"

	"github.com/bashmohandes/go-askme/internal/domain"
)

type questionsRepo struct {
}

// LoadQuestions loads the specified user's set of questions
func (r *questionsRepo) LoadQuestions(userID models.UniqueID) []models.Question {
	result := []models.Question{
		models.Question{
			UserEntity: models.UserEntity{
				Entity: models.Entity{
					CreatedOn: time.Now(),
					ID:        models.NewUniqueID(),
				},
				CreatedBy: models.User{
					Entity: models.Entity{
						CreatedOn: time.Now(),
						ID:        models.NewUniqueID(),
					},
					Email: "Bashmohandes@live.com",
					Name:  "Mohamed Elsherif",
				},
			},
			Text: "Who Are You?",
		},
		models.Question{
			UserEntity: models.UserEntity{
				Entity: models.Entity{
					CreatedOn: time.Now(),
					ID:        models.NewUniqueID(),
				},
				CreatedBy: models.User{
					Entity: models.Entity{
						CreatedOn: time.Now(),
						ID:        models.NewUniqueID(),
					},
					Email: "Bashmohandes@live.com",
					Name:  "Mohamed Elsherif",
				},
			},
			Text: "Where Are You?",
		},
	}

	return result
}

// NewQuestionRepository creates a new repo object
func NewQuestionRepository() models.QuestionRepository {
	return &questionsRepo{}
}