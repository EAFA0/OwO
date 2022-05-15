package register

import (
	"github.com/EAFA0/toast/pkg/conf"
	"github.com/EAFA0/toast/pkg/log"
)

var component = map[string]func() error{}

func Component(name string, initFunc func() error) {
	component[name] = initFunc
}

func Init() {
	// 加载配置
	conf.Init()
	// 初始化日志

	// 逐个加载 component
	for name, iniFunc := range component {
		log.Info("load component", log.String("name", name))
		if err := iniFunc(); err != nil {
			log.Error("component init failed", log.Err(err))
		}
	}
}
