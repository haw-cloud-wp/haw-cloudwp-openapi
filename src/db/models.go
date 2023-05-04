package db

import (
	"time"
)

var ModelList = []any{
	Bucket{},
	User{},
}

type Bucket struct {
	Name      string `gorm:"primaryKey"`
	CreatedAt time.Time
	User      []*User `gorm:"many2many:user_bucket;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

type User struct {
	Auth0ID string    `gorm:"primaryKey"`
	Access  []*Bucket `gorm:"many2many:user_bucket;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
