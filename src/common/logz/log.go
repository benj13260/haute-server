package logz

import (
	"context"
	"time"

	"github.com/sethvargo/go-envconfig"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var LOGZ *zap.SugaredLogger

var Version = "dev"
var App = "go-dev"

type SlackExporterConfig struct {
	WebhookURL  string `env:"SLACK_WEBHOOK_URL"`
	Channel     string `env:"SLACK_CHANNEL,default=zlogz_token_watcher"`
	ChannelSlow string `env:"SLACK_CHANNEL_SLOW,default=zlogz_slow_queries"`
}

type LogExportConfig struct {
	WithSlack bool `env:"LOG_EXPORT_SLACK,default=false"`
	Slack     SlackExporterConfig
}

type LoggerConfig struct {
	Format string `env:"LOG_FMT,default=pretty-console"`
	Export LogExportConfig
}

func LoadCfg() LoggerConfig {
	var conf LoggerConfig
	envconfig.Process(context.TODO(), &conf)
	return conf
}

func RFC3339TimeEncoder(ts time.Time, encoder zapcore.PrimitiveArrayEncoder) {
	encoder.AppendString(ts.UTC().Format(time.RFC3339))
}

func init() {
	var err error
	var logger *zap.Logger
	var loggerConf = LoadCfg()

	if err = zap.RegisterEncoder("pretty-console",
		func(ec zapcore.EncoderConfig) (zapcore.Encoder, error) {
			return NewEncoder(1, true), nil
		}); err != nil {
		panic(err)
	}

	cfg := zap.NewProductionConfig()
	cfg.DisableStacktrace = true
	cfg.EncoderConfig.EncodeTime = RFC3339TimeEncoder
	cfg.Encoding = loggerConf.Format
	cfg.EncoderConfig.LevelKey = "severity"
	cfg.EncoderConfig.MessageKey = "message"
	cfg.EncoderConfig.TimeKey = "time"
	// cfg.InitialFields = map[string]interface{}{
	// 	"logging.googleapis.com/sourceLocation": struct {
	// 		File     string `json:"file"`
	// 		Line     string `json:"line"`
	// 		Function string `json:"function"`
	// 	}{},
	// }

	logger, err = cfg.Build()
	if err != nil {
		panic("failed to build logger: " + err.Error())
	}

	if loggerConf.Export.WithSlack {
		defErrLevel := zap.WarnLevel
		logger = logger.WithOptions(
			zap.Hooks(
				NewSlackHook(
					loggerConf.Export.Slack.WebhookURL,
					loggerConf.Export.Slack.Channel,
					loggerConf.Export.Slack.ChannelSlow,
					App+"_"+Version,
					&defErrLevel,
				).GetHook(),
			),
		)
	}
	if App != "go-dev" {
		logger = logger.Named(App)
	}
	LOGZ = logger.Sugar()
	defer LOGZ.Sync()
}
