package main

import (
	"diary/dao"
	"diary/model"
	"fmt"
)

func init() {
	fmt.Println("main init")
}

func main() {
	dao.InitMySql()
	//engine := router.InitRouter()
	//engine.Run()
	db := dao.DB
	db = db.Exec("update user set phone=15915838227 where id = ?",2)

	fmt.Println("更新了",db.RowsAffected,"条数据")
	var users []model.User
	db.Raw("select * from user").Find(&users)
	fmt.Println("users",users)

}
