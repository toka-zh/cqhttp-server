package config

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
)

var Config struct {
	Port             string  `mapstructure:"port"`
	Version          string  `mapstructure:"version"`
	PixivRankUrl     string  `mapstructure:"pixiv_rank_url"`
	WhitelistPrivate []int64 `mapstructure:"whitelist_private"`
	WhitelistGroup   []int64 `mapstructure:"whitelist_group"`

	WhitelistPrivateFlg bool `mapstructure:"whitelist_private_con"`
	WhitelistGroupFlg   bool `mapstructure:"whitelist_group_con"`

	WhiteListPrivateMap map[int64]struct{}
	WhiteListGroupMap   map[int64]struct{}
}

// load config
func init() {
	viper.SetConfigFile("./config.yml")
	err := viper.ReadInConfig()
	if err != nil {
		return
	}

	// 监控并重新读取配置文件
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		log.Println("config file changed")
		err := viper.Unmarshal(&Config)
		if err != nil {
			return
		}
	})

	err = viper.Unmarshal(&Config)
	if err != nil {
		return
	}

	//sort.Ints()
	Config.WhiteListGroupMap = make(map[int64]struct{})
	Config.WhiteListPrivateMap = make(map[int64]struct{})
	for _, id := range Config.WhitelistGroup {
		Config.WhiteListGroupMap[id] = struct{}{}
	}
	for _, id := range Config.WhitelistPrivate {
		Config.WhiteListPrivateMap[id] = struct{}{}
	}
}
