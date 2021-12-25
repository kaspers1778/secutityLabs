package internal

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Repository struct {
	db *gorm.DB
}

func NewRepo(db *sql.DB) *Repository {
	pgConf := postgres.Config{Conn: db}
	gormConf := &gorm.Config{Logger: logger.Default.LogMode(logger.Info)}
	pg := postgres.New(pgConf)

	gormDb, err := gorm.Open(pg, gormConf)
	if err != nil {
		panic(fmt.Errorf("gorm: can not open: %v", err))
	}
	return &Repository{gormDb}
}

func (r *Repository) Create( usr User) error {
	gormUser := toGORMUser(usr)
	if err := r.db.Create(gormUser).Error; err != nil {
		return fmt.Errorf("gorm: can not create: %v", err)
	}
	return nil
}

func (r *Repository) Get(name string) (*User, error) {
	gormUser := gormUser{}
	if result := r.db.Where("Name = ?", name).First(&gormUser); result.Error != nil {
		return nil, fmt.Errorf("gorm: can not find: %v", result.Error)
	}
	return &User{
		Name: gormUser.Name,
		PswHash: gormUser.Password,
		Number: Decrypt(gormUser.Phonenumber),
		Color: Decrypt(gormUser.Color),
	}, nil
}

type gormUser struct {
	Id       int `gorm:"primary_key, auto_increment;not_null"`
	Name     string
	Password []byte
	Phonenumber []byte
	Color 	  []byte
}

func toGORMUser(usr User) *gormUser {
	return &gormUser{
		Name:     usr.Name,
		Password: usr.PswHash,
		Phonenumber: Encrypt(usr.Number),
		Color: Encrypt(usr.Color),
	}
}
