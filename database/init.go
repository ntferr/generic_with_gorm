package database

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitGorm() *gorm.DB {

	dsn := fmt.Sprintf(
		"host=%v user=%v password=%v dbname=%v port=%v sslmode=disable",
		"127.0.0.1",
		"postgres",
		"123456",
		"generics",
		"5432",
	)

	dialector := postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	})

	db, err := gorm.Open(dialector)
	if err != nil {
		fmt.Fprintf(os.Stdout, "failed to open connection: %s", err)
	}

	return db
}
