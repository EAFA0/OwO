package conf

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"github.com/EAFA0/Toast/pkg/env"
	"github.com/EAFA0/Toast/pkg/log"
	"github.com/spf13/viper"
)

func Init() {
	BasicConfig()
	EnvRelatedConfig()
}

var v = viper.New()

var DefaultConfigPath = "config"
var ProjectPathSkipStep = 4

// BasicConfig 加载默认配置及环境无关的配置项
func BasicConfig() {
	if err := loadConfigFromFolder(projectConfigPath()); err != nil {
		log.Panic("load basic config failed", log.Err(err))
	}
}

// EnvRelatedConfig 加载环境相关的配置项
func EnvRelatedConfig() {
	envConfigPath, envVal := "env_config", env.Local()
	// 如果设置了环境相关的设置, 则合并环境相关配置
	if !IsSet(envConfigPath) {
		return
	}

	// 装载环境相关配置
	path := GetString(envConfigPath, string(envVal))
	if err := loadConfigFromFolder(path); err != nil {
		log.Panic("load env related config failed", log.Err(err))
	}
}

func IsSet(nodes ...string) bool {
	return v.IsSet(NodePath(nodes...))
}
func SetDefault(val interface{}, nodes ...string) {
	v.SetDefault(NodePath(nodes...), val)
}
func GetString(nodes ...string) string {
	return v.GetString(NodePath(nodes...))
}
func GetBool(nodes ...string) bool {
	return v.GetBool(NodePath(nodes...))
}
func GetInt(nodes ...string) int {
	return v.GetInt(NodePath(nodes...))
}
func GetInt64(nodes ...string) int64 {
	return v.GetInt64(NodePath(nodes...))
}
func GetUInt(nodes ...string) uint {
	return v.GetUint(NodePath(nodes...))
}
func GetUInt64(nodes ...string) uint64 {
	return v.GetUint64(NodePath(nodes...))
}
func GetFloat64(nodes ...string) float64 {
	return v.GetFloat64(NodePath(nodes...))
}
func GetDuration(nodes ...string) time.Duration {
	return v.GetDuration(NodePath(nodes...))
}
func GetIntSlice(nodes ...string) []int {
	return v.GetIntSlice(NodePath(nodes...))
}
func GetStringSlice(nodes ...string) []string {
	return v.GetStringSlice(NodePath(nodes...))
}
func GetStringMap(nodes ...string) map[string]interface{} {
	return v.GetStringMap(NodePath(nodes...))
}
func GetStringMapString(nodes ...string) map[string]string {
	return v.GetStringMapString(NodePath(nodes...))
}
func GetStringMapStringSlice(nodes ...string) map[string][]string {
	return v.GetStringMapStringSlice(NodePath(nodes...))
}

func MustGetString(nodes ...string) string          { return v.GetString(MustSet(nodes...)) }
func MustGetBool(nodes ...string) bool              { return v.GetBool(MustSet(nodes...)) }
func MustGetInt(nodes ...string) int                { return v.GetInt(MustSet(nodes...)) }
func MustGetInt64(nodes ...string) int64            { return v.GetInt64(MustSet(nodes...)) }
func MustGetUInt(nodes ...string) uint              { return v.GetUint(MustSet(nodes...)) }
func MustGetUInt64(nodes ...string) uint64          { return v.GetUint64(MustSet(nodes...)) }
func MustGetFloat64(nodes ...string) float64        { return v.GetFloat64(MustSet(nodes...)) }
func MustGetDuration(nodes ...string) time.Duration { return v.GetDuration(MustSet(nodes...)) }
func MustGetIntSlice(nodes ...string) []int         { return v.GetIntSlice(MustSet(nodes...)) }
func MustGetStringSlice(nodes ...string) []string   { return v.GetStringSlice(MustSet(nodes...)) }
func MustGetStringMap(nodes ...string) map[string]interface{} {
	return v.GetStringMap(MustSet(nodes...))
}
func MustGetStringMapString(nodes ...string) map[string]string {
	return v.GetStringMapString(MustSet(nodes...))
}
func MustGetStringMapStringSlice(nodes ...string) map[string][]string {
	return v.GetStringMapStringSlice(MustSet(nodes...))
}

// NodePath 拼接配置路径
func NodePath(nodes ...string) string {
	return strings.Join(nodes, ".")
}

// MustSet 拼接配置路径 & 配置不存在时 panic
func MustSet(nodes ...string) string {
	path := NodePath(nodes...)
	if v.IsSet(path) {
		return path
	}
	format := "the config item [%s] is miss"
	panic(fmt.Sprintf(format, path))
}

// projectConfigPath 获取项目配置文件夹路径
func projectConfigPath() string {
	_, file, _, ok := runtime.Caller(ProjectPathSkipStep)
	if !ok {
		return filepath.Join("../", DefaultConfigPath)
	}
	return filepath.Join(filepath.Dir(file), DefaultConfigPath)
}

// loadConfigFromFolder 加载指定目录下的所有配置文件, 跳过文件夹内容
func loadConfigFromFolder(path string) error {
	files, err := os.ReadDir(path)
	if err != nil {
		return err
	}

	// 遍历该目录, 获取目录下的所有文件名并逐个加载
	for _, item := range files {
		// 跳过目录对象
		if item.IsDir() {
			continue
		}
		// merge config
		filename := filepath.Join(path, item.Name())
		v.SetConfigFile(filename)
		if err := v.MergeInConfig(); err != nil {
			return err
		}
	}
	return nil
}
