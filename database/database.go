package database

import (
	"database/sql"
	"go-sessions/types"

	"github.com/go-sql-driver/mysql"
)

func Connect() *sql.DB {
	config := mysql.Config{
		User: "session",
		Passwd: "password",
		DBName: "go_sessions",
		AllowNativePasswords: true,
	}

	sql, _ := sql.Open("mysql", config.FormatDSN())

	return sql
}

func AuthenticateUser(conn *sql.DB, username, password string) types.User {
	stmt, _ := conn.Prepare("SELECT * FROM user WHERE username = ? AND password = ?")

	var user types.User
	_ = stmt.QueryRow(username, password).Scan(&user.ID, &user.Username, &user.Password, &user.Role)

	return user
}

func GetUserByID(conn *sql.DB, userID int) types.User {
	stmt, _ := conn.Prepare("SELECT * FROM user WHERE id = ?")

	var user types.User
	_ = stmt.QueryRow(userID).Scan(&user.ID, &user.Username, &user.Password, &user.Role)

	return user
}
