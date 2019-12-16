package dbops

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

//创建用户
func AddUserCredential(loginName string, pwd string) error {
	stmt, err := dbConn.Prepare("INSERT INTO users (login_name, pwd) values (?, ?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(loginName, pwd)
	if err != nil {
		return err
	}
	defer stmt.Close()
	return nil
}

//通过loginName获取密码
func GetUserCredential(loginName string) (string, error) {
	stmt, err := dbConn.Prepare("SELECT pwd from users where login_name = ?")
	if err != nil {
		log.Printf("%s", err)
		return "", err
	}
	var pwd string
	err = stmt.QueryRow(loginName).Scan(&pwd)
	if err != nil && err != sql.ErrNoRows {
		return "", err
	}
	defer stmt.Close()
	return pwd, nil
}

//删除用户
func DeleteUser(loginName string, pwd string) error {
	stmt, err := dbConn.Prepare("DELETE FROM users where login_name = ? and pwd = ?")
	if err != nil {
		log.Printf("%s", err)
		return err
	}
	_, err = stmt.Exec(loginName, pwd)
	if err != nil {
		return err
	}
	defer stmt.Close()
	return nil
}
