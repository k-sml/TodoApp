package models

import (
	"crypto/sha1"
	// "golang.org/x/crypto/bcrypt"
	"fmt"
	"log"
	"database/sql"
	"gostudy/application/config"

	"github.com/google/uuid"

	_ "github.com/mattn/go-sqlite3"
)

var Db *sql.DB

var err error

const (
	tableNameUser = "users"
	tableNameTodo = "todos"
	tableNameSession = "sessions"
)

func init() {
	Db, err = sql.Open(config.Config.SQLDriver, config.Config.DbName)  //ドライバーの名前, DBの名前
	if err != nil {
		log.Fatalln(err)
	}

	cmdU := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		uuid STRING NOT NULL UNIQUE,   
		name STRING,
		email STRING,
		password STRING,
		created_at DATETIME
	)`, tableNameUser)
	//uuidはユーザーを識別するためのもの

	Db.Exec(cmdU) //データベースへのSQLクエリを実行するために使用される

	cmdT := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		content TEXT,
		user_id INTEGER,
		created_at DATETIME
	)`, tableNameTodo)
	Db.Exec(cmdT)

	cmdS := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		uuid STRING NOT NULL UNIQUE,
		email STRING,
		user_id INTEGER,
		created_at DATETIME
	)`, tableNameSession)
	Db.Exec(cmdS)
}

func createUUID() (uuidobj uuid.UUID) {
	uuidobj, _ = uuid.NewUUID()   //完全な一意性はない
	return uuidobj
}

//sha1は推奨されていない
func Encrypt(plaintext string) (cryptext string) {
	cryptext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
	return cryptext
}

//bcryptを使用する(少し計算コスト高い)
// func Encrypt(plaintext string) string {
// 	hash, _ := bcrypt.GenerateFromPassword([]byte(plaintext), bcrypt.DefaultCost)
// 	return string(hash)
// }