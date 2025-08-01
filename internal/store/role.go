package store

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Role struct {
	Id   int64  `json:"id"`
	Role string `json:"role"`
}

type RoleAPI struct {
	db *sql.DB
}

func (f *RoleAPI) Role(ctx *gin.Context) ([]Role, error) {
	query := `SELECT * FROM role`

	UserRole, err := f.db.Query(query)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Error": "Database query failed"})
		return nil, fmt.Errorf("database query failed: %w", err)
	}

	defer UserRole.Close()

	var DataRole []Role

	for UserRole.Next() {
		var DataRoleRows Role
		err := UserRole.Scan(
			&DataRoleRows.Id,
			&DataRoleRows.Role,
		)

		if DataRoleRows.Id <= 0 || DataRoleRows.Role == "" {
			ctx.JSON(http.StatusInternalServerError, gin.H{"Error": "Id Or Role Empty"})
			return nil, err
		}

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"Error": "Something Went Wrong With Data Role"})
			return nil, err
		}

		DataRole = append(DataRole, DataRoleRows)
	}

	if err = UserRole.Err(); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Message": "There was an error on Get Data Role"})
		return nil, err
	}

	return DataRole, nil

}
