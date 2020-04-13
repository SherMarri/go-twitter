package migrations

import (
	_ "github.com/go-sql-driver/mysql" // SQL Driver
	"github.com/jinzhu/gorm"
)

// Tweet Model representing a post in db
type Tweet struct {
	gorm.Model
	Content string `gorm:"NOT_NULL"`
	UserID  uint   `gorm:"NOT_NULL"`
}

// Tweets202009041738 Migration object
var Tweets202009041738 = Migration{
	Name: "Tweets202009041738",
	Up: func(db *gorm.DB) error {
		db.CreateTable(&Tweet{})
		db.Model(&Tweet{}).AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT")
		return nil
	},
	Down: func(db *gorm.DB) error {
		db.DropTable(&Follower{})
		return nil
	},
}
