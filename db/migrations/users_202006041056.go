package migrations

import (
	_ "github.com/go-sql-driver/mysql" // SQL Driver
	"github.com/jinzhu/gorm"
)

// User Model representing
type User struct {
	gorm.Model
	Username string `gorm:"NOT_NULL; UNIQUE"`
	Password string `gorm:"NOT_NULL"`
	Name     string `gorm:"NOT_NULL"`
	IsActive bool   `gorm:"NOT_NULL; DEFAULT:true"`
}

// Users202006041056 Migration object
var Users202006041056 = Migration{
	Name: "Users202006041056",
	Up: func(db *gorm.DB) error {
		db.CreateTable(&User{})
		return nil
	},
	Down: func(db *gorm.DB) error {
		db.DropTable(&User{})
		return nil
	},
}
