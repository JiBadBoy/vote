package setting

import (
	"log"
	"os"
	"time"

	"github.com/go-ini/ini"
)

type App struct {
	JwtSecret       string
	PageSize        int
	RuntimeRootPath string

	LogSavePath string
	LogSaveName string
	LogFileExt  string
	TimeFormat  string
}

var AppSetting = &App{}

type Server struct {
	RunMode      string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

var ServerSetting = &Server{}

type Database struct {
	User     string
	Password string
	Host     string
	Port     int
	Name     string
}

var DatabaseSetting = &Database{}

var cfg *ini.File

func Setup() {

	var err error
	dir, _ := os.Getwd()

	cfg, err = ini.Load(dir + "/conf/app.ini")
	if err != nil {
		log.Fatalf("Fail to parse"+dir+" '/conf/app.ini': %v", err)
	}

	mapTo("app", AppSetting)
	mapTo("server", ServerSetting)
	mapTo("database", DatabaseSetting)

	ServerSetting.ReadTimeout = ServerSetting.ReadTimeout * time.Second
	ServerSetting.WriteTimeout = ServerSetting.ReadTimeout * time.Second

}

func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo %s err: %v", section, err)
	}
}
