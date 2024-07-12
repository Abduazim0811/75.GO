package main

import (
	"log"

	"75.GO/api"
	"75.GO/internal/mongodb"
)

func main() {
	db, err := mongodb.NewStudent("mongodb://localhost:27017", "student", "students")
	if err != nil {
		log.Fatal(err)
	}
	api.Routes(db)
}
