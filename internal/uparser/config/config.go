package config

import (
	"errors"
	"time"
)

type (
	Config struct {
		Environment string
		FileStorage S3Config
		HTTP        HTTPConfig
		JWT         JWTConfig
		Postgresql  PostgresqlConfig
		Admin       AdminConfig
		Redis       RedisConfig
		Proxy       ProxyConfig
	}

	S3Config struct {
		Access        string
		Secret        string
		SysBucket     string
		ContentBucket string
	}

	HTTPConfig struct {
		Port string
		Host string
	}

	JWTConfig struct {
		AccessTokenTTL  time.Duration
		RefreshTokenTTL time.Duration
		PrivateKey      string
		PublicKey       string
	}

	PostgresqlConfig struct {
		Host     string
		User     string
		Password string
	}

	AdminConfig struct {
		Login    string
		Password string
		ID       string
	}

	RedisConfig struct {
		Host string
		Port string
	}

	ProxyConfig struct {
		Host string
		Port string
	}
)

func Init(uip string, configPath string) (*Config, error) {

	if len(uip) >= 0 {
		pCfg, err := uipSetup(uip)
		if err != nil {
			return nil, err
		}

		return pCfg, nil
	}

	if len(uip) >= 0 && len(configPath) > 0 {
		// TODO viper set config
		//return pCfg, nil
	}

	return nil, errors.New("config keys not set")
}

func uipSetup(uip string) (*Config, error) {
	//var cfg Config

	// TODO uip get config

	return nil, errors.New("uip not available")

	//return &cfg, nil
}
