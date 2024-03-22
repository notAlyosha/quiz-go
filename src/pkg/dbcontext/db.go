// Package dbcontext provides DB transaction support for transactions tha span method calls of multiple
// repositories and services.
package dbcontext

import (
	dbx "github.com/go-ozzo/ozzo-dbx"
	_ "github.com/go-sql-driver/mysql"
	"github.com/notAlyosha/quiz-go/internal/config"
)

var db *dbx.DB

func CreateDBConnection() error {
	if db == nil {
		c := config.LoadConfig()
		connectionString := c.DBUserName + ":" + c.DBUserPassword + "@tcp(" + c.DBHost + ":" + c.DBPort + ")/" + c.DBName

		// Connect to the database
		dbInstance, err := dbx.Open("mysql", connectionString)
		if err != nil {
			return err
		}

		// Set the maximum number of idle database connections
		dbInstance.DB().SetMaxIdleConns(10)

		db = dbInstance
		return nil
	}
	return nil
}

func GetDB() *dbx.DB {
	if db != nil {
		return db
	}

	return nil
}
