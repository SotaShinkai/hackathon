package main

import (
	"crypto/rand"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/go-sql-driver/mysql"
	"github.com/oklog/ulid"
)

type UserResForHTTPGet struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type UserReqForHTTPPost struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type UserResForHTTPPost struct {
	Id string `json:"id"`
}

// ① GoプログラムからMySQLへ接続
var db *sql.DB

func init() {
	// ①-1
	user := os.Getenv("MYSQL_USER")
	password := os.Getenv("MYSQL_PASSWORD")
	database := os.Getenv("MYSQL_DATABASE")
	mysqlUser := user
	mysqlUserPwd := password
	mysqlDatabase := database
	mysqlHost := os.Getenv("MYSQL_HOST")

	// ①-2
	connStr := fmt.Sprintf("%s:%s@%s/%s", mysqlUser, mysqlUserPwd, mysqlHost, mysqlDatabase)
	_db, err := sql.Open("mysql", connStr)
	if err != nil {
		log.Fatalf("fail: sql.Open, %v\n", err)
	}
	// ①-3
	if err := _db.Ping(); err != nil {
		log.Fatalf("fail: _db.Ping, %v\n", err)
	}
	db = _db
}

// ② /userでリクエストされたらnameパラメーターと一致する名前を持つレコードをJSON形式で返す
func handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		// ②-1
		name := r.URL.Query().Get("name")
		if name == "" {
			log.Println("fail: name is empty")
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// ②-2
		rows, err := db.Query("SELECT id, name, age FROM user WHERE name = ?", name)
		if err != nil {
			log.Printf("fail: db.Query, %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// ②-3
		users := make([]UserResForHTTPGet, 0)
		for rows.Next() {
			var u UserResForHTTPGet
			if err := rows.Scan(&u.Id, &u.Name, &u.Age); err != nil {
				log.Printf("fail: rows.Scan, %v\n", err)

				if err := rows.Close(); err != nil { // 500を返して終了するが、その前にrowsのClose処理が必要
					log.Printf("fail: rows.Close(), %v\n", err)
				}
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			users = append(users, u)
		}

		// ②-4
		bytes, err := json.Marshal(users)
		if err != nil {
			log.Printf("fail: json.Marshal, %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(bytes)
	case http.MethodPost:
		decoder := json.NewDecoder(r.Body)
		var data UserReqForHTTPPost

		if err := decoder.Decode(&data); err != nil {
			log.Printf("fail: decoder.Decode, %v\n", err)
			http.Error(w, "Failed to decode JSON data", http.StatusInternalServerError)
			return
		}
		if data.Name == "" || len(data.Name) > 50 || data.Age < 20 || data.Age > 80 {
			log.Printf("fail: data.Name, data.Age is too large\n")
			http.Error(w, "JSON data error", http.StatusInternalServerError)
			return
		}

		id := ulid.MustNew(ulid.Now(), rand.Reader)
		_, err := db.Exec("INSERT INTO user VALUES (?, ?, ?)", id.String(), data.Name, data.Age)
		if err != nil {
			log.Printf("fail: db.Query, %v\n", err)
			http.Error(w, "Failed to insert data", http.StatusInternalServerError)
			return
		}

		idStruct := UserResForHTTPPost{Id: id.String()}
		bytes, err := json.Marshal(idStruct)
		if err != nil {
			log.Printf("fail: json.Marshal, %v\n", err)
			http.Error(w, "Failed to JSON data", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(bytes)
	default:
		log.Printf("fail: HTTP Method is %s\n", r.Method)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

func main() {
	// ② /userでリクエストされたらnameパラメーターと一致する名前を持つレコードをJSON形式で返す
	http.HandleFunc("/user", handler)

	// ③ Ctrl+CでHTTPサーバー停止時にDBをクローズする
	closeDBWithSysCall()

	// 8000番ポートでリクエストを待ち受ける
	log.Println("Listening...")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}

// ③ Ctrl+CでHTTPサーバー停止時にDBをクローズする
func closeDBWithSysCall() {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGTERM, syscall.SIGINT)
	go func() {
		s := <-sig
		log.Printf("received syscall, %v", s)

		if err := db.Close(); err != nil {
			log.Fatal(err)
		}
		log.Printf("success: db.Close()")
		os.Exit(0)
	}()
}
