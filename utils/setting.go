package utils

import (
	"fmt"

	"gopkg.in/ini.v1"
)

var (
	//AppMode AppMode
	AppMode string
	//HttpPort HttpPort
	HttpPort string
	JwtKey   string

	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassWord string
	DbName     string

	AccessKey string
	SecretKey string
	Bucket string
	Qiniuserver string
)

func init() {
	file, err := ini.Load("config/config.ini")
	if err != nil {
		fmt.Println("配置文件读取错误，请检查文件路径：", err)
	}

	loadServer(file)
	loadDatabase(file)
	loadQiniu(file)
}

func loadServer(file *ini.File) {
	AppMode = file.Section("server").Key("AppMode").MustString("debug")
	HttpPort = file.Section("server").Key("HttpPort").MustString(":3000")
	JwtKey = file.Section("server").Key("JwtKey").MustString("uuahdfuaohsdq")

}

func loadDatabase(file *ini.File) {
	Db = file.Section("database").Key("Db").MustString("mysql")
	DbHost = file.Section("database").Key("DbHost").MustString("localhost")
	DbPort = file.Section("database").Key("DbPort").MustString("3306")
	DbUser = file.Section("database").Key("DbUser").MustString("root")
	DbPassWord = file.Section("database").Key("DbPassWord").MustString("huangpeng123")
	DbName = file.Section("database").Key("DbName").MustString("ginblog")
}

func loadQiniu(file *ini.File)  {
	AccessKey = file.Section("qiniu").Key("AccessKey").String()
	SecretKey = file.Section("qiniu").Key("SecretKey").String()
	Bucket = file.Section("qiniu").Key("Bucket").String()
	Qiniuserver = file.Section("qiniu").Key("Qiniuserver").String()
}

//5421325246

//421325246