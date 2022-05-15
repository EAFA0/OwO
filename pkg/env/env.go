package env

// 枚举值, 标记当前所处环境
type Env string

const DEV Env = "dev"

// Local 确定本机所处环境
var Local = env

func env() Env {
	return DEV
}
