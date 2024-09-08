package user

import (
	"time"

	"github.com/google/uuid"
	"github.com/rayhan889/rss-aggr/internal/database"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	Email     string	`json:"email"`
	Name      string	`json:"name"`
	Password  string	`json:"password"`
	CreatedAt time.Time	`json:"created_at"`
	UpdatedAt time.Time	`json:"updated_at"`
}

func HandleUserToUserCustomModel(dbUser database.User) User {
	return User {
		ID: dbUser.ID,
		Email: dbUser.Email,
		Name: dbUser.Name,
		Password: dbUser.Password,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
	}
}