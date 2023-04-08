package main

import (
	"database/sql"
	"errors"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// トランザクションのロールバックを確認
func main() {
	tx1()
}

func tx1() {
	db, _ := sql.Open("mysql", "root:root@tcp(127.0.0.1:33306)/db")

	tx, err := db.Begin()
	if err != nil {
		// ここでのエラーは想定していない
		panic(err)
	}

	// 何かクエリを発行する
	result, err := tx.Exec("INSERT INTO user (name) VALUES (?)", "Someone")
	if err != nil {
		// ここでのエラーは想定していない
		panic(err)
	}
	log.Println("result: ", result)

	// err = nil // コミットする
	err = errors.New("Something wrong!!!") // ロールバックする

	// 何かの処理で問題が起こったとき
	if err != nil {
		log.Println("Rollback!")
		log.Println("error: ", err)
		tx.Rollback()
	} else {
		log.Println("Commit!")
		tx.Commit()
	}
}
