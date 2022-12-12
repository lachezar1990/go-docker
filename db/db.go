package db

import (
	"database/sql"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func SetupDB() error {
	err := createDBIfNeeded()
	if err != nil {
		return err
	}

	dsn := "root:1234567@tcp(db)/docker_test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return err
	}

	db.AutoMigrate(&Task{})

	return nil
}

func createDBIfNeeded() error {
	dsn := "root:1234567@tcp(db)/"
	db, err := sql.Open("mysql", dsn)

	if err != nil {
		return err
	}

	defer db.Close()

	_, err = db.Exec("CREATE DATABASE  IF NOT EXISTS `docker_test` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;")
	if err != nil {
		return err
	}

	_, err = db.Exec("USE `docker_test`")
	if err != nil {
		return err
	}

	return nil
}
