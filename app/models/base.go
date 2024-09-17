package models

import (
	"crypto/sha1"
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"log"
	"todo_app/config"

	_ "github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
)

var Db *sql.DB

var err error

const (
	tableNameUser    = "users"
	tableNameTodo    = "todos"
	tableNameSession = "sessions"
)

/* 初期化 */
func init() {

	// データベースに接続
	Db, err = sql.Open(config.Config.SQLDriver, config.Config.DbName)
	if err != nil {
		log.Fatal(err)
	}

	// ユーザテーブルの作成
	cmdU := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    uuid STRING NOT NULL UNIQUE,
    name STRING,
    email STRING,
    password STRING,
    created_at DATETIME)`, tableNameUser)
	Db.Exec(cmdU)

	// TODOテーブルの作成
	cmdT := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    content TEXT,
    user_id INTEGER,
    created_at DATETIME)`, tableNameTodo)
	Db.Exec(cmdT)

	// sessionテーブルの作成
	cmdS := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    uuid STRING NOT NULL UNIQUE,
    email STRING,
    user_id INTEGER,
    created_at DATETIME)`, tableNameSession)
	Db.Exec(cmdS)
}

/* uuid生成 */
func createUUID() (uuidobj uuid.UUID) {
	uuidobj, _ = uuid.NewUUID()
	return uuidobj
}

/* パスワード暗号化 */
func Encrypt(plaintext string) (cryptext string) {
	cryptext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
	return cryptext
}
