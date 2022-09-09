package config

import (
	"fmt"
	"github.com/go-ini/ini"
	"log"
	"micro_api/micro_common/utils"
	"micro_api/models"
)

var (
	Cfg       *ini.File
	IsDebug   bool
	ConfigRes = models.Config{}

	ClientInfo = make(map[string]models.ClientInfo, 0)

	DcProduct = "dcProduct"
	DcUser    = "dcUser"
)

func InitConfig() {

	// 读取配置文件
	Cfg, _ = ini.Load("./config/app.ini")
	debug := Cfg.Section("").Key("RUN_MOD").MustString("")
	if debug == "debug" {
		IsDebug = true
	}

	//读取内部配置
	serverSection, err := Cfg.GetSection("server")
	if err != nil {
		log.Fatal("Fail to load section 'server': ", err)
	}
	mServer := &models.Server{
		HttpPort:  serverSection.Key("HTTP_PORT").MustInt(8080),
		PprofPort: serverSection.Key("PPROF_PORT").MustInt(8081),
	}

	ConfigRes.Ser = mServer

	mysqlSection, err := Cfg.GetSection("mysql")
	if err != nil {
		log.Fatal("Fail to load section 'mysql': ", err)
	}
	mysql := &models.MysqlInI{
		Addr: mysqlSection.Key("ADDR").MustString("localhost"),
		Port: mysqlSection.Key("PORT").MustInt(13306),
		Db:   mysqlSection.Key("DB").MustInt(0),
		Pwd:  mysqlSection.Key("PWD").MustString(""),
		User: mysqlSection.Key("USER").MustString(""),
	}
	ConfigRes.Mysql = mysql

	redisSection, err := Cfg.GetSection("redis")
	if err != nil {
		log.Fatal("Fail to load section 'redis': ", err)
	}
	redis := &models.RedisInI{
		Addr: redisSection.Key("ADDR").MustString("localhost"),
		Port: redisSection.Key("PORT").MustInt(6370),
		Db:   redisSection.Key("DB").MustInt(0),
		Pwd:  redisSection.Key("PWD").MustString(""),
	}
	ConfigRes.Redis = redis

	fmt.Println("读取的配置文件: ", utils.JsonToString(ConfigRes))

	//获取dc_product的id和端口
	dcProductSection, err := Cfg.GetSection(DcProduct)
	if err != nil {
		log.Fatal("Fail to load section 'mysql': ", err)
	}
	dcProductInfo := models.ClientInfo{
		Addr: dcProductSection.Key("ADDR").MustString("127.0.0.1"),
		Port: dcProductSection.Key("PORT").MustString("8802"),
	}

	ClientInfo[DcProduct] = dcProductInfo

	// 读取dc_user的ip和端口
	dcUserSection, err := Cfg.GetSection(DcUser)
	if err != nil {
		log.Fatal("Fail to load section 'mysql': ", err)
	}
	userInfo := models.ClientInfo{
		Addr: dcUserSection.Key("ADDR").MustString("127.0.0.1"),
		Port: dcUserSection.Key("PORT").MustString("8803"),
	}
	ClientInfo[DcUser] = userInfo
}
