package store

import (
	"database/sql"
	"net/http"

	"github.com/Razor4456/FoundationBackEnd/utils"
	"github.com/gin-gonic/gin"
)

type PostUsers struct {
	Id         int64  `json:"id"`
	Email      string `json:"email"`
	Username   string `json:"username"`
	Name       string `json:"name"`
	Password   string `json:"password"`
	Role       string `json:"role"`
	VerifLogin string `json:"veriflogin"`
	CreatedAt  string `json:"created_at"`
}

type UsersLogin struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type UsersAPI struct {
	db *sql.DB
}

func (f *UsersAPI) Login(ctx *gin.Context, Logins *UsersLogin) error {
	query := `SELECT id, username, password FROM users WHERE username = $1`
	users := PostUsers{}
	err := f.db.QueryRow(
		query,
		Logins.Username,
	).Scan(
		&users.Id,
		&users.Username,
		&users.Password,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusUnauthorized, gin.H{"Error": "There was an error "})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"Error": "Database error"})
		}
		return err

	}
	err = utils.HashValidation(Logins.Password, users.Password)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Error": "Password Wrong "})
		return err
	}

	LoginQuery := `UPDATE users SET veriflogin = $1 WHERE id = $2`

	_, nil := f.db.ExecContext(ctx,
		LoginQuery,
		"True",
		users.Id,
	)

	return nil
}

func (f *UsersAPI) CreateUsers(ctx *gin.Context, PostUsers *PostUsers) error {
	query := `INSERT INTO users(email, username, name, password, role, veriflogin) VALUES ($1,$2,$3,$4,$5,$6) RETURNING id, created_at`

	err := f.db.QueryRowContext(
		ctx,
		query,
		PostUsers.Email,
		PostUsers.Username,
		PostUsers.Name,
		PostUsers.Password,
		PostUsers.Role,
		PostUsers.VerifLogin,
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
