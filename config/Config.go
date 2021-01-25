package config

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	AppName string   `json:"app_name"`
	AppMode string   `json:"app_mode"`
	AppHost string   `json:"app_host"`
	AppPort string   `json:"app_port"`
	Database DatabaseConfig `database`
}
type DatabaseConfig struct {
	Driver   string `json:"driver"`
	User     string `json:"user"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	DbName   string `json:"db_name"`
	Charset  string `json:"charset"`
	ShowSql  bool   `json:"show_sql"`
}

func init() {
	fmt.Println("config init")
	cfg,err := ParseConfig("./config/app.json")
	if err != nil{
		fmt.Println(err)
		return
	}
	fmt.Println(cfg.AppName,cfg.AppHost)
	fmt.Println(cfg.Database)
}
var _cfg *Config

func GetConfig() *Config  {
	return _cfg
}
func ParseConfig(path string) (*Config,error)  {
	file ,err := os.Open(path)
	if err != nil{
		panic(err)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	decoder := json.NewDecoder(reader)
	if err :=decoder.Decode(&_cfg);err != nil{
		return _cfg,err
	}
	fmt.Printf("ParseConfig:%p\n",_cfg)
	return _cfg,nil
}

