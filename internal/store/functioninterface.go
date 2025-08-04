package store

import (
	"database/sql"
	"errors"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

var (
	ErrNotFound = errors.New("record not found")
)

type FunctionStore struct {
	Stuff interface {
		CreateStuff(*gin.Context, *PostStuff) error
		DeleteStuff(*gin.Context, []int64) ([]DeletedStuff, error)
		GetDataStuff(*gin.Context) ([]PostStuff, error)
		EditStuff(*gin.Context, *PostStuff) error
	}
	Users interface {
		CreateUsers(*gin.Context, *PostUsers) error
		Login(*gin.Context, *UsersLogin) (*Tokens, error)
		Logout(*gin.Context, *StoreLogout) error
	}

	Role interface {
		Role(*gin.Context) ([]Role, error)
	}

	Transaction interface {
		Cart(*gin.Context, *Transaction) error
	}
}

func FunctionStorage(db *sql.DB) FunctionStore {
	return FunctionStore{
		Stuff:       &StuffApi{db},
		Users:       &UsersAPI{db},
		Role:        &RoleAPI{db},
		Transaction: &TransactionAPI{db},
	}
}
