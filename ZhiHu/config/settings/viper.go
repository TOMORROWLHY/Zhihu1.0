package settings

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func InitViper() (err error) {
	viper.SetConfigName("config")   // 配置文件名称(无扩展名)
	viper.SetConfigType("yaml")     // 指定配置文件类型
	viper.AddConfigPath("./config") // 指定查找配置文件的路径（这里使用相对路径）
	err = viper.ReadInConfig()      // 查找并读取配置文件
	if err != nil {
		fmt.Printf("viper.ReadInConfig failed, err:%#v", err)
		// 处理读取配置文件的错误
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
		return
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件发生修改......")
	})
	return err
}
