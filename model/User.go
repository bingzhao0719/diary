package model

import (
	"diary/dao"
	"fmt"
	"time"
)

type User struct {
	Id        int    `gorm:"column id;primary_key" json:"id"`
	Name      string `gorm:"column name;" json:"name"`
	Phone     string `gorm:"column phone;" json:"phone"`
	Password  string `gorm:"column password;" json:"password"`
	CreatedAt string `gorm:"column created_at;" json:"createAt"`
}

func init() {
	//QueryUserById(1)
	//QueryUserByPhone("15915838227")

}

func AddUser(user *User) {
	db := dao.DB
	user.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	db.Create(user)
}
func QueryUserById(id int) *User {
	return QueryUserByWhere("id=?", id)
}
func QueryUserByPhone(phone string) *User {
	return QueryUserByWhere("phone=?", phone)
}
func QueryUserByWhere(where ...interface{}) *User {
	fmt.Println("queryUser where:", where)
	db := dao.DB
	var user User
	err := db.First(&user, where...).Error
	if err != nil {
		fmt.Println("QueryUserByWhere err:", err, user)
		return nil
	}
	fmt.Println("QueryUserByWhere user:", user)
	return &user
}

func DeleteUserById(id int) *User {
	db := dao.DB
	var user User
	db.Delete(&user, id)
	fmt.Println("DeleteUserById user:", user)
	return &user
}

func QueryUserListByWhere(where ...interface{}) *[]User {
	fmt.Println("QueryUserListByWhere where:", where)
	db := dao.DB
	var user []User
	err := db.Find(&user, where...).Error
	if err != nil {
		fmt.Println("QueryUserListByWhere err:", err, user)
		return nil
	}
	fmt.Println("QueryUserListByWhere user:", user)
	return &user
}
