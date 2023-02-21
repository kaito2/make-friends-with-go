package main

import (
	"log"
	"os"
	"path"
	"time"
)

func main() {
	log.Println("=== Start ===")

	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	log.Println("Working Directory: ", wd)
	dataDirectory := path.Join(wd, "/data")
	log.Println("Data Directory: ", dataDirectory)

	am := NewAuditManager(
		3,
		dataDirectory,
	)

	if err := am.AddRecord("Tom", time.Now()); err != nil {
		log.Fatalf("failed to AddRecord: %v", err)
	}
}
