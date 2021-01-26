package dao

import (
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func InitMySql() {
	CreateUser()
}
func CreateUser() {
	//sqlStr := `CREATE TABLE IF NOT EXISTS user2(
	//    id bigint primary key auto_increment not null,
	//	phone varchar(20),
	//	name varchar(20),
	//	password varchar(20),
	//	created_at varchar(20));`
	//if !db.HasTable(&model.User{}){
	//	if err := db.Set("gorm:table_options","ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&model.User{}).Error;err != nil{
	//		fmt.Println("CreateTable err:",err)
	//	}
	//}
}
func execDB(sql string, args ...interface{}) {

}
