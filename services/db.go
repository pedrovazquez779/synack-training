package services

import "database/sql"

type IDBService interface {
	GetDB() (*sql.DB, error)
}

type DBService struct{}

func (d *DBService) GetDB() (*sql.DB, error) {
	return sql.Open("mysql", `root:root1234@tcp(localhost:3306)/video_club`)
}
