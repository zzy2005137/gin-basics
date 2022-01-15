//加载配置
package setting

import (
	"fmt"
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

var cfg = &ini.File{}

func Setup() {
	var err error
	cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatal("load config file failed")
	}

	MapTo("server", ServerSetting)

	ServerSetting.ReadTimeout = ServerSetting.ReadTimeout * time.Second
	ServerSetting.WriteTimeout = ServerSetting.WriteTimeout * time.Second

	fmt.Println(ServerSetting)
}

func MapTo(section string, v interface{}) {
	err := cfg.Section("server").MapTo(ServerSetting)
	if err != nil {
		log.Fatalf("failed to map to section : %s", section)
	}
}
