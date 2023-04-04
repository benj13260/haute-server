package config

import (
	"fmt"

	"github.com/spf13/viper"
)

var (
	Instance *CommonConf
)

type Stage struct {
	Stage string `env:"STAGE,default=dev"`
}

type ServerConfig struct {
	HealthPort int `env:"HEALTHPORT,default=9999"`
	Port       int `env:"PORT"`
}

type DbConfig struct {
	Dbname                string `env:"DATABASE_NAME,required"`
	Host                  string `env:"DATABASE_HOST,required"`
	Port                  string `env:"DATABASE_PORT,required"`
	User                  string `env:"DATABASE_USER,required"`
	Password              string `env:"DATABASE_PASSWORD,required"`
	Lang                  string `env:"DATABASE_LANG,default=postgresql"`
	Sslmode               string `env:"sslmode,default=postgresql"`
	TimeZone              string `env:"TimeZone,default=postgresql"`
	ConcurrentProtected   bool   `env:"ConcurrentProtected,default=false"`
	MaxConcurrentCalls    int32  `env:"MaxConcurrentCalls,default=10"`
	MaxConcurrentRequests int    `env:"MaxConcurrentRequests,default=200"`
}

func (dc *DbConfig) URL(appName string) string {
	return fmt.Sprintf("%s://%s:%s@%s:%s/%s?application_name=%s",
		dc.Lang, dc.User, dc.Password, dc.Host, dc.Port, dc.Dbname, appName)
}

type RedisConfig struct {
	Protocol    string `env:"REDIS_PROTOCOL,default=tcp"`
	Server      string `env:"REDIS_SERVER"`
	IdleThreads int    `env:"REDIS_THREAD,default=50"`
	IdleTimeout int    `env:"REDIS_TIMEOUT,default=240"`
	URL         string `env:"REDIS_URL"`
}

type RedisCacheConfig struct {
	*RedisConfig `env:",prefix=CACHE_"`
}

type RedisSessionConfig struct {
	*RedisConfig `env:",prefix=SESSION_"`
}

type BrokerConfig struct {
	Kind                 string `env:"BROKER_KIND,default=rabbitmq"` // pubsub or rabbitmq
	RabbitURL            string `env:"RABBITMQ_URL"`
	RabbitChannel        string `env:"RABBITMQ_CHANNEL"`
	RabbitConsumer       string `env:"RABBITMQ_CONSUMER"`
	PubSubProjectID      string `env:"PUBSUB_PROJECT_ID"`
	PubSubSubscriptionID string `env:"PUBSUB_SUBSCRIPTION_ID"`
}

type CommonConf struct {
	Database   DbConfig
	CacheRedis *RedisConfig `env:",prefix=CACHE_"`
	AuthRedis  *RedisConfig `env:",prefix=SESSION_"` // TODO pass to prefix = auth_
	Server     ServerConfig
	Broker     BrokerConfig

	Stage           string `env:"STAGE,default=dev"`
	Testing         bool   `env:"TESTING,default=false"`
	RealtimeEnabled bool   `env:"REALTIME_ENABLED"`
}

func init() {
	LoadConfig(".")
}

func LoadConfig(path string) (config CommonConf, err error) {
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("yaml")   // REQUIRED if the config file does not have the extension in the name

	viper.AddConfigPath(".")   // optionally look for config in the working directory
	err = viper.ReadInConfig() // Find and read the config file
	if err != nil {            // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	err = viper.Unmarshal(&config)
	Instance = &config
	return
}
