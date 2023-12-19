package config

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type (
	AppConfig struct {
		Name    string `mapstructure:"name"`
		Version string `mapstructure:"version"`
	}
	DatabaseConfig struct {
		Name     string `mapstructure:"name"`
		Host     string `mapstructure:"host"`
		Port     string `mapstructure:"port"`
		Username string `mapstructure:"username"`
		Password string `mapstructure:"password"`
	}

	RESTConfig struct {
		Port     int   `mapstructure:"port"`
		BodySize int64 `mapstructure:"bodySize"`
		Debug    bool  `mapstructure:"debug"`
		Timeout  struct {
			Read  int64 `mapstructure:"read"`
			Write int64 `mapstructure:"write"`
			Idle  int64 `mapstructure:"idle"`
		}
	}

	RPCConfig struct {
		Port int `mapstructure:"port"`
	}

	ServerConfig struct {
		REST *RESTConfig `mapstructure:"rest"`
		RPC  *RPCConfig  `mapstructure:"rpc"`
	}

	LoggingConfig struct {
		Level  int    `mapstructure:"level"`
		Format string `mapstructure:"format"`
	}

	SecurityConfig struct {
		JWTSecret     string `mapstructure:"jwtsecret"`
		JWTExpiryTime int    `mapstructure:"jwtexpirytime"`
	}

	DriverConfig struct {
		App      *AppConfig      `mapstructure:"app"`
		Database *DatabaseConfig `mapstructure:"database"`
		Server   *ServerConfig   `mapstructure:"server"`
		Logging  *LoggingConfig  `mapstructure:"logging"`
		Security *SecurityConfig `mapstructure:"security"`
	}

	BootstrapConfig struct {
		RESTServer   *gin.Engine
		Log          *logrus.Logger
		SqlDB        *sql.DB
		DriverConfig *DriverConfig
	}
)
