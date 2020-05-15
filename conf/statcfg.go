package conf

import (
	"log"
)

// StatCfg 统计 配置文件1
type StatCfg struct {
	BaseConf
	Total int //总访问量

}

// Stat 全局声明
var Stat StatCfg

func init() {
	log.Println("StatCfg init ENTER")
	Abt.BaseConf.conf = &Stat

}
