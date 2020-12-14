package global

import (
	"fmt"
	"github.com/prometheus/common/log"
	"github.com/spf13/viper"
	"io/ioutil"
	"os"
	"strings"
)

// 数据库配置项
var cfgDatabase *viper.Viper

// 载入配置文件
func Setup(path string) {
	viper.SetConfigFile(path)
	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(fmt.Sprintf("Read config file fail: %s", err.Error()))
	}

	// Replace environment variables

	//log.Info("-----------------------------------------------")
	//// 返回字节
	//log.Info(content)
	//// 将字节转化为 string
	//log.Info(string(content))
	//log.Info(os.ExpandEnv(string(content)))
	//// 返回类型 *Reader
	//log.Info(strings.NewReader(os.ExpandEnv(string(content))))
	//log.Info("-----------------------------------------------")

	err = viper.ReadConfig(strings.NewReader(os.ExpandEnv(string(content))))
	cfgDatabase = viper.Sub("database")
	if cfgDatabase == nil {
		panic("No found mysql in the configuration")
	}
	DatabaseConfig = InitDatabase(cfgDatabase)

}
