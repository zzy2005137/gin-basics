//加载配置
package setting

import (
	"log"
	"time"

	"github.com/go-ini/ini"
)

type Server struct {
	RunMode      string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

var ServerSetting = &Server{}

type App struct {
	JwtSecret string
}

var AppSetting = &App{}

var cfg = &ini.File{}

func Setup() {
	var err error
	cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatal("load config file failed")
	}

	MapTo("server", ServerSetting)
	MapTo("app", AppSetting)

	ServerSetting.ReadTimeout = ServerSetting.ReadTimeout * time.Second
	ServerSetting.WriteTimeout = ServerSetting.WriteTimeout * time.Second

}

func MapTo(section string, v interface{}) {
	err := cfg.Section("server").MapTo(ServerSetting)
	if err != nil {
		log.Fatalf("failed to map to section : %s", section)
	}
}
