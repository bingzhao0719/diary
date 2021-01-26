package main

import (
	"diary/router"
	"fmt"
)

func init() {
	fmt.Println("main init")
}

func main() {
	//dao.InitMySql()
	engine := router.InitRouter()
	engine.Run()

}
