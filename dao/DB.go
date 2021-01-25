package dao

import (
	"diary/config"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func init() {
	fmt.Println("db init")
	cfg := config.GetConfig()
	database := cfg.Database
	conn := database.User + ":" + database.Password + "@tcp(" + database.Host +":"+database.Port+")/"+
		database.DbName+"?"+ "charset="+database.Charset
	fmt.Println(conn)
	db,err := gorm.Open(database.Driver,conn)
	//engine, err := gorm.Open("mysql", "root:root123456@tcp(127.0.0.1:3306)/diary?charset=utf8")
	if err != nil{
		fmt.Println("err:",err)
		return
	}
	db.SingularTable(true)
	db.LogMode(true)
	DB = db
	fmt.Println("数据库连接成功",db)

	//if !db.HasTable(&model.User{}){
	//	if err := db.Set("gorm:table_options","ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&model.User{}).Error;err != nil{
	//		fmt.Println("CreateTable err:",err)
	//		return
	//	}
	//}
	//插入一条数据
	//user := model.User{Name: "wubz",Password: "123456",CreatedTime: time.Now()}
	//db.Create(&user)
	//var firstUser model.User
	//db.First(&firstUser)
	//fmt.Println("firstUser:",firstUser)

	//deleteUser := model.User{Id: 2,Name: "wubz",Password: "123456"}
	//db.Delete(&deleteUser)
}
