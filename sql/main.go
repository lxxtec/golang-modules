package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" // 匿名引入驱动
)

// 初始化连接
func InitDivider() (db *sql.DB) {
	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", "root", "", "127.0.0.1", "3306", "testsql")
	db, err := sql.Open("mysql", url)
	if err != nil {
		fmt.Println(err)
		return
	}
	return db
}
// 插入
func TestInsert(db *sql.DB) {
	sql := `INSERT INTO t_user_info(user_id, username, account_id, account_name, gender, phone, avatar) values(?,?,?,?,?,?,?)`
	_, err := db.Exec(sql, "1", "测试用户", "10", "测试账号", 1, "18732132132", "https://moose-plus.oss-cn-shenzhen.aliyuncs.com/avatar.png")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("添加成功")
}
// 更新
func TestUpdate(db *sql.DB) {
	sql := `UPDATE t_user_info SET username = ? WHERE user_id = ?`
	result, err := db.Exec(sql, "修改名字", "1")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)
}
// 删除
func TestDelete(db *sql.DB) {
	sql := `DELETE FROM t_user_info WHERE user_id = ?`
	result, err := db.Exec(sql, "1")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)
}
type UserInfo struct {
	UserId   string
	UserName string
	Avatar   string
}
// 查询
func TestQueryRow(db *sql.DB) {
	// 查询 SQL
	sql := "select user_id, avatar, username from t_user_info"

	var userInfo UserInfo
	err := db.QueryRow(sql).Scan(&userInfo.UserId, &userInfo.UserName, &userInfo.Avatar)
	if err != nil {
		fmt.Println("sacn error :: ", err)
		return
	}
	fmt.Println(userInfo)
}
// 查询多条数据
func TestQuery(db *sql.DB) {
	sql := "select user_id, avatar, username from t_user_info where user_id IN(?,?)"
	rows, err := db.Query(sql, "785919644501544960", "790883082524954600")
	if err != nil {
		fmt.Println(err)
		return
	}

	for rows.Next() {
		var userInfo UserInfo
		err = rows.Scan(&userInfo.UserId, &userInfo.UserName, &userInfo.Avatar)
		if err != nil {
			fmt.Println("发送错误")
			return
		}
		fmt.Println(userInfo)
	}
}


func main() {

	db := InitDivider()

	defer db.Close()

	//TestInsert(db)
	//TestUpdate(db)
	//TestQueryRow(db)
	//TestQuery(db)
	//TestDelete(db)
}
