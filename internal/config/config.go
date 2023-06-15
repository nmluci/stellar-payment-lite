package config

import (
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
)

type Config struct {
	ServiceName    string      `json:"serviceName"`
	ServiceAddress string      `json:"servicePort"`
	Environment    Environment `json:"environment"`

	BuildVer  string
	BuildTime string

	WorkerConfig  WorkerConfig  `json:"workerConfig"`
	MariaDBConfig MariaDBConfig `json:"mariaDBConfig"`
	MongoDBConfig MongoDBConfig `json:"mongoDBConfig"`
	RedisConfig   RedisConfig   `json:"redisConfig"`

	JWT_ISSUER         string
	JWT_AT_EXPIRATION  time.Duration
	JWT_RT_EXPIRATION  time.Duration
	JWT_SIGNING_METHOD jwt.SigningMethod
	JWT_SIGNATURE_KEY  []byte
}

const logTagConfig = "[Init Config]"

var config *Config

func Init(buildTime, buildVer string) {
	godotenv.Load("conf/.env")

	conf := Config{
		ServiceName:    os.Getenv("SERVICE_NAME"),
		ServiceAddress: os.Getenv("SERVICE_ADDR"),
		MariaDBConfig: MariaDBConfig{
			Address:  os.Getenv("MARIADB_ADDRESS"),
			Username: os.Getenv("MARIADB_USERNAME"),
			Password: os.Getenv("MARIADB_PASSWORD"),
			DBName:   os.Getenv("MARIADB_DBNAME"),
		},
		MongoDBConfig: MongoDBConfig{
			Address:  os.Getenv("MONGODB_ADDRESS"),
			Username: os.Getenv("MONGODB_USERNAME"),
			Password: os.Getenv("MONGODB_PASSWORD"),
			DBName:   os.Getenv("MONGODB_DBNAME"),
		},
		RedisConfig: RedisConfig{
			Address:  os.Getenv("REDIS_ADDRESS"),
			Port:     os.Getenv("REDIS_PORT"),
			Password: os.Getenv("REDIS_PASSWORD"),
		},
		WorkerConfig: WorkerConfig{},
		BuildVer:     buildVer,
		BuildTime:    buildTime,
	}

	if conf.ServiceName == "" {
		log.Fatalf("%s service name should not be empty", logTagConfig)
	}

	if conf.ServiceAddress == "" {
		log.Fatalf("%s service port should not be empty", logTagConfig)
	}

	if conf.MariaDBConfig.Address == "" || conf.MariaDBConfig.DBName == "" {
		log.Fatalf("%s address and db name cannot be empty", logTagConfig)
	}

	envString := os.Getenv("ENVIRONMENT")
	if envString != "dev" && envString != "prod" && envString != "local" {
		log.Fatalf("%s environment must be either local, dev or prod, found: %s", logTagConfig, envString)
	}

	conf.Environment = Environment(envString)

	conf.Environment = Environment(envString)
	conf.JWT_ISSUER = os.Getenv("JWT_ISSUER")
	conf.JWT_SIGNING_METHOD = jwt.SigningMethodHS256
	conf.JWT_SIGNATURE_KEY = []byte(os.Getenv("JWT_SIGNATURE_KEY"))
	conf.JWT_AT_EXPIRATION = time.Duration(2) * time.Hour
	conf.JWT_RT_EXPIRATION = time.Duration(24) * time.Hour

	config = &conf
}

func Get() (conf *Config) {
	conf = config
	return
}
