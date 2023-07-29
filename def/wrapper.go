package def

import (
	cfgDef "github.com/requiemofthesouls/config/def"
	"github.com/requiemofthesouls/container"

	"github.com/requiemofthesouls/logger"
)

const DIWrapper = "logger.wrapper"

type Wrapper = logger.Wrapper

func init() {
	container.Register(func(builder *container.Builder, params map[string]interface{}) error {
		return builder.Add(container.Def{
			Name: DIWrapper,
			Build: func(container container.Container) (_ interface{}, err error) {
				var cfg cfgDef.Wrapper
				if err = container.Fill(cfgDef.DIWrapper, &cfg); err != nil {
					return nil, err
				}

				var loggerCfg logger.Config
				if err = cfg.UnmarshalKey("logger", &loggerCfg); err != nil {
					return nil, err
				}

				serviceName := cfg.GetString("service")
				if serviceName == "" {
					serviceName = "system"
				}

				return logger.New(
					loggerCfg,
					[]logger.Field{
						logger.String(logger.KeyServiceName, serviceName),
					},
				)
			},
			Close: func(obj interface{}) error {
				_ = obj.(logger.Wrapper).Sync()
				return nil
			},
		})
	})
}
