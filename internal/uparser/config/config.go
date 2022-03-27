package config

import (
	"errors"
	"time"
)

type (
	Config struct {
		Environment string
		Logger      LoggerConfig
		FileStorage S3Config
		HTTP        HTTPConfig
		JWT         JWTConfig
		Postgresql  PostgresqlConfig
		Admin       AdminConfig
		Redis       RedisConfig
		Proxy       ProxyConfig
		VPN         VPNConfig
	}

	LoggerConfig struct {
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
		ID       string
		Login    string
		Password string
	}

	RedisConfig struct {
		Host string
		Port string
	}

	ProxyConfig struct {
		Host string
		Port string
	}

	VPNConfig struct {
		Host     string
		Port     string
		Login    string
		Password string
	}
)

func Init(uip *string, configPath *string) (*Config, error) {

	// UIP
	if len(*uip) >= 0 {
		cfg, err := uipSetup(uip)
		if err != nil {
			return nil, err
		}
		return cfg, nil
	}

	// JSON Config
	if len(*uip) >= 0 && len(*configPath) > 0 {
		cfg, err := getJSONConfig(configPath)
		if err != nil {
			return nil, err
		}
		return cfg, nil
	}

	// Else return error
	return nil, errors.New("config keys error")
}

// Load config via UIP protocol
func uipSetup(uip *string) (*Config, error) {
	//var cfg Config

	// TODO uip get config
	// func AddSecureRemoteProvider (viper)

	return nil, errors.New("uip not available")

	//return &cfg, nil
}

// Load config via JSON file
func getJSONConfig(configPath *string) (*Config, error) {
	//var cfg Config

	// TODO JSON config load

	//s := strings.Split(*configPath, "/")
	//viper.SetConfigName(s[len(s)-1])
	// не включать расширение
	// TODO заменить везде configPath на configName и фиксировать папку конфигураций

	return nil, errors.New("son config not available")

	//return &cfg, nil
}
