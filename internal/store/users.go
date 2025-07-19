package store

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PostUsers struct {
	Id        int64  `json:"id"`
	Email     string `json:"email"`
	Username  string `json:"username"`
	Name      string `json:"name"`
	Password  string `json:"password"`
	Role      string `json:"role"`
	CreatedAt string `json:"created_at"`
}

type UsersAPI struct {
	db *sql.DB
}

func (f *UsersAPI) CreateUsers(ctx *gin.Context, PostUsers *PostUsers) error {
	query := `INSERT INTO users(email, username, name, password,role) VALUES ($1,$2,$3,$4,$5) RETURNING id, created_at`

	err := f.db.QueryRowContext(
		ctx,
		query,
		PostUsers.Email,
		PostUsers.Username,
		PostUsers.Name,
		PostUsers.Password,
		PostUsers.Role,
	).Scan(
		&PostUsers.Id,
		&PostUsers.CreatedAt,
	)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Error": "There was an error when insert CreateUsers"})
		ctx.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return nil
	}

	return nil
}
