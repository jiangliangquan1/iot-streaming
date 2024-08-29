package viperex

import (
	"fmt"
	"github.com/spf13/cast"
	"github.com/spf13/viper"
	"io"
	"regexp"
	"strings"
)

type Option struct {
	ConfigFile string
	BindEnv    bool
}

type ViperEx struct {
	*viper.Viper
}

// 创建对象
func New() *ViperEx {
	vex := new(ViperEx)
	vex.Viper = viper.New()
	return vex
}

func NewViperEx(option *Option) *ViperEx {
	v := New()
	v.SetConfigFile(option.ConfigFile)

	// 读取配置文件
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Sprintf("Error reading config file, %s", err))
	}

	if option.BindEnv {
		v.BindEnv()
	}

	return v
}

// 绑定环境变量
func (vex *ViperEx) BindEnv() {
	vex.SetEnvPrefix("")
	vex.AutomaticEnv()

	keys := vex.AllKeys()
	for _, key := range keys {
		value := vex.Viper.Get(key)

		//fmt.Printf("Key: %s, Value: %v\n", key, value)

		if stringValue, ok := value.(string); ok {
			re := regexp.MustCompile(`^\${(.*?)}`)
			match := re.FindStringSubmatch(stringValue)
			if len(match) > 1 {
				envConfig := match[1]
				if len(envConfig) > 0 {

					envName := strings.Split(envConfig, ":")[0]
					vex.Viper.BindEnv(key, envName)
					fmt.Printf("bind config_key: %s to env: %s\n", key, envName)
				}
			}
		}
	}
}

// 重写Get方法
func (vex *ViperEx) Get(key string) any {
	value := vex.Viper.Get(key)

	if stringValue, ok := value.(string); ok {
		re := regexp.MustCompile(`^\${(.*?)}`)
		match := re.FindStringSubmatch(stringValue)
		if len(match) > 1 {
			envConfig := match[1]
			if len(envConfig) > 0 && strings.Contains(envConfig, ":") {
				defaultValue := strings.Split(envConfig, ":")[1]
				return defaultValue
			}
		}
	}

	return value
}

func (vex *ViperEx) GetString(key string) string   { return cast.ToString(vex.Get(key)) }
func (vex *ViperEx) GetBool(key string) bool       { return cast.ToBool(vex.Get(key)) }
func (vex *ViperEx) GetInt(key string) int         { return cast.ToInt(vex.Get(key)) }
func (vex *ViperEx) GetInt32(key string) int32     { return cast.ToInt32(vex.Get(key)) }
func (vex *ViperEx) GetInt64(key string) int64     { return cast.ToInt64(vex.Get(key)) }
func (vex *ViperEx) ReadConfig(in io.Reader) error { return vex.Viper.ReadConfig(in) }
func (vex *ViperEx) SetConfigFile(in string)       { vex.Viper.SetConfigFile(in) }

var v *ViperEx

func init() {
	v = New()
}

func Get(key string) any            { return v.Get(key) }
func GetString(key string) string   { return v.GetString(key) }
func GetBool(key string) bool       { return v.GetBool(key) }
func GetInt(key string) int         { return v.GetInt(key) }
func GetInt32(key string) int32     { return v.GetInt32(key) }
func GetInt64(key string) int64     { return v.GetInt64(key) }
func SetConfigName(in string)       { v.SetConfigName(in) }
func SetConfigType(in string)       { v.SetConfigType(in) }
func AddConfigPath(in string)       { v.AddConfigPath(in) }
func ReadInConfig() error           { return v.ReadInConfig() }
func BindEnv()                      { v.BindEnv() }
func ReadConfig(in io.Reader) error { return v.ReadConfig(in) }
func SetConfigFile(in string)       { v.SetConfigFile(in) }
