package config

import (
	"github.com/BurntSushi/toml"
	"strings"
)

type Mysql struct {
	Host      string
	Port      int
	Database  string
	Username  string
	Password  string
	Charset   string
	ParseTime bool `toml:"parse_time"`
	Loc       string
}

type Redis struct {
	IP       string
	Port     int
	Database int
}

type Server struct {
	IP   string
	Port int
}

type Path struct {
	FfmpegPath       string `toml:"ffmpeg_path"`
	StaticSourcePath string `toml:"static_source_path"`
}

type Config struct {
	DB     Mysql `toml:"mysql"`
	RDB    Redis `toml:"redis"`
	Server `toml:"server"`
	Path   `toml:"path"`
}

var Info Config

func init() {
	if _, err := toml.DecodeFile("D:\\Go_WorkSpace\\src\\My_douyin\\config\\config.toml", &Info); err != nil {
		panic(err)
	}
	strings.Trim(Info.DB.Host, " ")
	strings.Trim(Info.RDB.IP, " ")
	strings.Trim(Info.Server.IP, " ")

}
