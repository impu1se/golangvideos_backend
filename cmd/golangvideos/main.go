package main

import (
	"github.com/go-pg/pg"
	"github.com/komly/golangvideos_backend/server"
	"github.com/komly/golangvideos_backend/service/videos"
	"log"
)

func main() {
	db := pg.Connect(&pg.Options{
		User:     "postgres",
		Password: "postgres",
		Database: "golangvideos",
	})
	defer db.Close()
	vs := videos.NewService(db)

	s := server.New(vs)
	if err := s.Start(); err != nil {
		log.Fatalf("can't start server: %s", err)
	}

}
