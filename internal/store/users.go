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
	Username string `json:"username"`
	Password string `json:"password"`
}

type Tokens struct {
	Id    int64  `json:"user_id"`
	Name  string `json:"name"`
	Token string `json:"token"`
}
type UsersAPI struct {
	db *sql.DB
}

func (f *UsersAPI) Login(ctx *gin.Context, Logins *UsersLogin) (*Tokens, error) {
	query := `SELECT id, username, name, password FROM users WHERE username = $1`
	users := PostUsers{}
	err := f.db.QueryRow(
		query,
		Logins.Username,
	).Scan(
		&users.Id,
		&users.Username,
		&users.Name,
		&users.Password,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusUnauthorized, gin.H{"Error": "There was an error "})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"Error": "Database error"})
		}
		return nil, err

	}

	err = utils.HashValidation(Logins.Password, users.Password)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Error": "Password Wrong "})
		return nil, err
	}

	if users.Id <= 0 {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Error": "User id not found"})
	}

	token, err := utils.GenerateToken(users.Id, users.Username)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Token Cannot be generate"})
		return nil, err
	}

	LoginQuery := `UPDATE users SET veriflogin = $1 WHERE id = $2`

	_, nil := f.db.ExecContext(ctx,
		LoginQuery,
		"True",
		users.Id,
	)

	tokens := Tokens{
		Id:    users.Id,
		Name:  users.Name,
		Token: token,
	}

	return &tokens, nil
}

type StoreLogout struct {
	Username string `json:"username"`
}
type UsersLogout struct {
	Id       int64  `json:"user_id"`
	Username string `json:"username"`
}

// type RevokeToken struct {
// 	Token      string `json:"token"`
// 	User_id    int64  `json:"user_id"`
// 	Expires_at string `json:"expires_at"`
// }

// func (f *UsersAPI) Logout(ctx *gin.Context) error {
// 	tokenString := strings.TrimPrefix(ctx.GetHeader("Authorization"), "Bearer")

// 	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
// 		return []byte(os.Getenv("JWT_SECRET")), nil
// 	})

// 	if err != nil {
// 		return ctx.AbortWithError(http.StatusUnauthorized, fmt.Errorf("invalid token"))
// 	}

// 	claims := token.Claims.(jwt.MapClaims)

// 	expUnix, ok := claims["exp"].(float64)
// 	if !ok {
// 		return ctx.AbortWithError(http.StatusBadRequest, fmt.Errorf("invalid exp claim"))
// 	}

// 	exp := time.Unix(int64(expUnix), 0)

// 	uidFloat, ok := claims["user_id"].(float64)

// 	if !ok {
// 		return ctx.AbortWithError(http.StatusBadRequest, fmt.Errorf("invalid token"))
// 	}

// 	userID := int(uidFloat)

// 	query := `INSERT INTO revok_token (token, user_id, expires_at) VALUES ($1, $2, $3)`
// 	_, err = f.db.ExecContext(
// 		ctx,
// 		query,
// 		tokenString,
// 		userID,
// 		exp,
// 	)

// 	if err != nil {
// 		return ctx.AbortWithError(http.StatusInternalServerError, err)
// 	}

// 	ctx.JSON(http.StatusOK, gin.H{"Message": "Token Revoked SuccessFully"})

// 	return nil

// }

func (f *UsersAPI) Logout(ctx *gin.Context, StoreLogout *StoreLogout) error {

	// tokenString := strings.TrimPrefix(ctx.GetHeader("Authorization"), "Bearer")

	// token, _ := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
	// 	return []byte(os.Getenv("JWT_SECRET")), nil
	// })

	// claim := token.Claims.(jwt.MapClaims)
	// exp := time.Unix(int64(claim["exp"].(float64)), 0)

	// query := `INSERT INTO revoked_tokens (token, expires_at) VALUES ($1, $2)`

	// _, err := f.db.ExecContext(
	// 	ctx,
	// 	query,
	// 	tokenString,
	// 	exp,
	// )

	// if err != nil {
	// 	ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed To Revoke Token"})
	// }

	// return nil

	query := `SELECT id, username FROM users WHERE username = $1`
	UserLogout := UsersLogout{}
	err := f.db.QueryRow(
		query,
		StoreLogout.Username,
	).Scan(
		&UserLogout.Id,
		&UserLogout.Username,
	)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "There was an error with Queryrow LogOut"})
		return err
	}

	queryLogout := `UPDATE users SET veriflogin = $1 WHERE id = $2`

	_, err = f.db.ExecContext(
		ctx,
		queryLogout,
		"False",
		UserLogout.Id,
	)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "There was an error with Query LogOut"})
		return err
	}

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

// func RevokToken(ctx *gin.Context, tokenstring string) error {
// 	var revoked bool

// 	var f *sql.DB

// 	query := `SELECT EXISTS(SELECT 1 FROM revok_token WHERE token = $1)`

// 	err := f.QueryRowContext(
// 		ctx,
// 		query,
// 		tokenstring,
// 	)

// 	if err != nil {
// 		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to check token revocation"})
// 		return nil
// 	}

// 	if revoked {
// 		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token has been revoked"})
// 		return nil
// 	}

// 	return nil

// }
