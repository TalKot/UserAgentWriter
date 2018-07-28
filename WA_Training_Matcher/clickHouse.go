package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/kshvakov/clickhouse"
	_ "github.com/lib/pq"
)

type eventStructer struct {
	Date      string `json:"date"`
	Referer   string `json:"referer"`
	EventType string `json:"eventType"`
	UserAgent string `json:"userAgent"`
}

var items = []eventStructer{}

func saveMessage(msg string) {
	userRequestData := eventStructer{}
	err := json.Unmarshal([]byte(msg), &userRequestData)
	if err != nil {
		fmt.Println(err)
	}

	connect, err := sql.Open("clickhouse", c.ConnectionDB)
	if err != nil {
		log.Fatal(err)
	}
	if err := connect.Ping(); err != nil {
		if exception, ok := err.(*clickhouse.Exception); ok {
			fmt.Printf("[%d] %s \n%s\n", exception.Code, exception.Message, exception.StackTrace)
		} else {
			fmt.Println(err)
		}
		return
	}
	defer connect.Close()

	tx, err := connect.Begin()
	if err != nil {
		log.Fatal(err)
	}

	stmt, err := tx.Prepare("INSERT INTO watest (date, referer, event, useragent) VALUES (?, ?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}

	if _, err := stmt.Exec(
		time.Now(),
		userRequestData.Referer,
		userRequestData.EventType,
		userRequestData.UserAgent,
	); err != nil {
		log.Fatal(err)
	}

	if err := tx.Commit(); err != nil {
		log.Fatal(err)
	}
}
