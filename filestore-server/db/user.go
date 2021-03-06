package db

import (
	"fmt"
	mydb "go_project/filestore-server/db/mysql"
)

// UserSignUp : 通过用户名及密码完成user表的注册操作
func UserSignUp(username string, passwd string) bool {
	stmt, err := mydb.DBConn().Prepare(
		"insert ignore into tbl_user (`user_name`, `user_pwd`) values (?,?)")
	if err != nil {
		fmt.Println("Failed to insert, err:" + err.Error())
		return false
	}
	defer stmt.Close()

	ret, err := stmt.Exec(username, passwd)
	if err != nil {
		fmt.Println("Failed to insert, err:" + err.Error())
		return false
	}
	if rowsAffected, err := ret.RowsAffected(); nil == err && rowsAffected > 0 {
		return true
	}
	return false
}

// UserSignIn : 判断密码是否一致
func UserSignIn(username string, encpwd string) bool {
	stmt, err := mydb.DBConn().Prepare("select * from tbl_user where user_name=? limit 1")
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	rows, err := stmt.Query(username)
	if err != nil {
		fmt.Println(err.Error())
		return false
	} else if rows == nil {
		fmt.Println("username not found:" + username)
		return false
	}

	pRows := mydb.ParseRows(rows)
	//对比password是否一致
	if len(pRows) > 0 && string(pRows[0]["user_pwd"].([]byte)) == encpwd {
		return true
	}
	return false
}

// UpdateToken : 刷新用户登录的token
func UpdateToken(username string, token string) bool {
	stmt, err := mydb.DBConn().Prepare(
		"replace into tbl_user_token (`user_name`, `user_token`) values (?,?)")
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	defer stmt.Close()

	_, err = stmt.Exec(username, token)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	return true
}

type User struct {
	Username     string
	Email        string
	Phone        string
	SignUpAt     string
	LastActiveAt string
	Status       int
}

func GetUserInfo(username string) (User, error) {
	userInfo := User{}
	stmt, err := mydb.DBConn().Prepare(
		"select user_name, signup_at from tbl_user where user_name=? limit 1")
	if err != nil {
		fmt.Println(err.Error())
		return userInfo, err
	}
	defer stmt.Close()

	// 执行查询的操作
	err = stmt.QueryRow(username).Scan(&userInfo.Username, &userInfo.SignUpAt)
	if err != nil {
		return userInfo, err
	}
	return userInfo, nil
}
