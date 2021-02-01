package models

import (
	"fmt"

	"gorm.io/gorm"
)

// Migrations func
func Migrations(db *gorm.DB) {
	var (
		checkUsers bool
		checkDogs  bool
		checkJokes bool
	)

	db.Migrator().DropTable(&Users{})
	db.Migrator().DropTable(&Dogs{})
	db.Migrator().DropTable(&Jokes{})

	checkUsers = db.Migrator().HasTable(&Users{})
	if !checkUsers {
		db.Migrator().CreateTable(&Users{})
		fmt.Println("Users table created")
	}

	checkJokes = db.Migrator().HasTable(&Jokes{})
	if !checkJokes {
		db.Migrator().CreateTable(&Jokes{})
		fmt.Println("Jokes table created")
	}

	checkDogs = db.Migrator().HasTable(&Dogs{})
	if !checkDogs {
		db.Migrator().CreateTable(&Dogs{})
		fmt.Println("Dogs table created")
	}
}
