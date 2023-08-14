package models

import "time"

type Config struct {
	Host string
	Port string
}

type Note struct {
	Id      int
	Content string
	Date    time.Time
}
