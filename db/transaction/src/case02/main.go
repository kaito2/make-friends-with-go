package main

import (
	"context"
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// 2つの読み込みロックを取るトランザクションを発行し、ロックすることを確認
func main() {
	const (
		targetUserID = "3"
	)

	ctx1 := context.Background()
	tx(ctx1, "a", targetUserID)
	ctx2 := context.Background()
	tx(ctx2, "b", targetUserID)
}

func tx(ctx context.Context, txID, targetUserID string) {
	log.Println("start tx: ", txID)
	defer log.Println("end tx: ", txID)

	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:33306)/db")
	if err != nil {
		// ここでのエラーは想定していない
		panic(err)
	}

	// 新しいコンテキストを作成し、3秒のタイムアウトを設定する
	ctx, _ = context.WithTimeout(ctx, 3*time.Second)
	// NOTE: cancel を読んでしまうと tx 関数終了時に tx がロールバックされてしまう
	// defer cancel()

	tx, err := db.BeginTx(ctx, nil)
	// tx, err := db.Begin()
	if err != nil {
		// ここでのエラーは想定していない
		panic(err)
	}

	// 何かクエリを発行する
	log.Println("start query")
	_, err = tx.Query(
		"SELECT * FROM user WHERE id=? FOR UPDATE;", // ロックを取る
		// "SELECT * FROM user WHERE id=?;", // ロックを取らない
		targetUserID,
	)
	if err != nil {
		// ここでのエラーは想定していない
		panic(err)
	}
	log.Println("end query")
}
