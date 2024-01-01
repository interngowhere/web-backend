package users

import (
	"github.com/interngowhere/web-backend/internal/database"
	"github.com/interngowhere/web-backend/internal/models"
)

func List(db *database.Database) ([]models.User, error) {
	users := []models.User{
		{
			ID:   1,
			Name: "CVWO",
		},
	}
	return users, nil
}
