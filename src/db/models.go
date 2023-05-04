package db

import (
	"gorm.io/gorm"
	"time"
)

var ModelList = []any{
	Bucket{},
	User{},
}

type Bucket struct {
	gorm.Model
	Name      string `gorm:"primaryKey"`
	CreatedAt time.Time
}

type User struct {
	gorm.Model
	Auth0ID string   `gorm:"primaryKey"`
	Access  []Bucket `gorm:"many2many:user_bucket;"`
}
