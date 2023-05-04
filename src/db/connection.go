package db

import (
	"database/sql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

type Connection struct {
	DB *gorm.DB
}

func (c *Connection) Init(db *sql.DB) *Connection {
	var err error
	c.DB, err = gorm.Open(mysql.New(mysql.Config{
		Conn: db,
	}), &gorm.Config{})
	if err != nil {
		log.Println("Error: ", err.Error())
		return nil
	}

	err = c.DB.AutoMigrate(ModelList...)
	if err != nil {
		log.Println("Error: ", err.Error())
		return nil
	}
	log.Println("migrated!")
	return c
}
