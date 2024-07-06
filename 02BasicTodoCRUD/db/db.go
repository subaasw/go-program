package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

const TABLE_NAME = "Todo"

type myDB struct {
	conn *sql.DB
}

type Todo struct {
	id                 int
	title, description string
}

func (todo Todo) String() string {
	return fmt.Sprintf("%d.\nTitle: %s,\nDescription: %s\n", todo.id, todo.title, todo.description)
}

func goDotEnvVariable(key string) string {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

// database connection
func DBConn() *myDB {
	dbDriver := "mysql"
	dbUser := goDotEnvVariable("DBUSER")
	dbName := goDotEnvVariable("DBNAME")
	dbPass := goDotEnvVariable("DBPASS")

	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return &myDB{conn: db}
}

func (db *myDB) Close() error {
	return db.conn.Close()
}

func (db *myDB) MaybeCreateTable() {

	query := `
		CREATE TABLE IF NOT EXISTS Todo (
			id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
			title TEXT NOT NULL,
			description TEXT
		);`

	_, err := db.conn.Exec(query)

	if err != nil {
		panic(err.Error())
	}

	println()
	fmt.Println("| Table created Successfully |")
	println()
}

func (db *myDB) GetById(todoId int) int {
	var todo Todo

	query := fmt.Sprintf("SELECT * FROM %s WHERE id = ?", TABLE_NAME)
	if err := db.conn.QueryRow(query, todoId).Scan(&todo.id, &todo.title, &todo.description); err != nil {
		log.Fatal(err)
	}

	fmt.Println(todo.String())
	return todo.id
}

func (db *myDB) GetAll() {
	query := fmt.Sprintf("SELECT * FROM %s", TABLE_NAME)

	rows, err1 := db.conn.Query(query)

	if err1 != nil {
		panic(err1)
	}
	defer rows.Close()

	for rows.Next() {
		var todo Todo

		err2 := rows.Scan(&todo.id, &todo.title, &todo.description)
		if err2 != nil {
			panic(err2)
		}

		fmt.Println(todo.String())
	}

	if err3 := rows.Err(); err3 != nil {
		panic(err3)
	}
}

func (db *myDB) AddTodo(title string, description string) int64 {
	query := fmt.Sprintf("INSERT INTO %s (title, description) VALUES (?, ?)", TABLE_NAME)

	result, err := db.conn.Exec(query, title, description)

	if err != nil {
		panic(err)
	}

	rowsAffected, _ := result.RowsAffected()

	if rowsAffected > 0 {
		fmt.Println(rowsAffected, "Added Successfully!")
	}
	return rowsAffected
}

func (db *myDB) UpdateTodo(id int, title string, description string) {
	query := fmt.Sprintf("UPDATE %s SET title = ? , description = ? WHERE id = ?", TABLE_NAME)

	result, err := db.conn.Exec(query, title, description, id)

	if err != nil {
		panic(err.Error())
	}

	rA, _ := result.RowsAffected()

	if rA > 0 {
		fmt.Println("Updated Successfully")
	} else {
		fmt.Println("Something went wrong!")
	}
}

func (db *myDB) Remove(id int) {
	query := fmt.Sprintf("DELETE from %s where id = ?", TABLE_NAME)
	res, err := db.conn.Exec(query, id)

	if err != nil {
		fmt.Println("Something went wrong!")
	}

	if rF, _ := res.RowsAffected(); rF > 0 {
		fmt.Println("Successfully Removed!")
	}
}
