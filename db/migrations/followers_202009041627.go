package migrations

import (
	_ "github.com/go-sql-driver/mysql" // SQL Driver
	"github.com/jinzhu/gorm"
)

// Follower Model representing a follower and following user
type Follower struct {
	gorm.Model
	UserID     uint
	FollowerID uint
}

// Followers202009041627 Migration object
var Followers202009041627 = Migration{
	Name: "Followers202009041627",
	Up: func(db *gorm.DB) error {
		db.CreateTable(&Follower{})
		db.Model(&Follower{}).AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT")
		db.Model(&Follower{}).AddForeignKey("follower_id", "users(id)", "RESTRICT", "RESTRICT")
		db.Model(&Follower{}).AddUniqueIndex("idx_follower_id_user_id", "follower_id", "user_id")
		return nil
	},
	Down: func(db *gorm.DB) error {
		db.DropTable(&Follower{})
		return nil
	},
}
