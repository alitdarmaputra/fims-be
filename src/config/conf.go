package config

import (
	"log"
	"strings"

	"github.com/spf13/viper"
)

type Api struct {
	Host                  string   `json:"host"                         mapstructure:"APP_HOST"`
	Port                  int      `json:"port"                         mapstructure:"APP_PORT"`
	Env                   string   `json:"env"                          mapstructure:"ENV"`
	JWTSecretKey          string   `json:"-"                            mapstructure:"JWT_SECRET_KEY"`
	JWTExpiredTime        int      `json:"jwt_expired_time"             mapstructure:"JWT_EXPIRED"`
	ResetTokenExpiredTime int      `json:"reset_token_expiredbali_time" mapstructure:"RESET_TOKEN_EXPIRED"`
	Database              Database `json:"database"`
	SMTP                  SMTP     `json:"smtp"`
	Figma                 Figma    `json:"figma"`
}

type Database struct {
	Host     string `json:"host"     mapstructure:"DATABASE_HOST"`
	Port     int    `json:"port"     mapstructure:"DATABASE_PORT"`
	Username string `json:"username" mapstructure:"DATABASE_USERNAME"`
	Password string `json:"password" mapstructure:"DATABASE_PASSWORD"`
	Schema   string `json:"schema"   mapstructure:"DATABASE_SCHEMA"`
	Loc      string `json:"loc"      mapstructure:"DATABASE_LOC"`
}

type SMTP struct {
	ClientOrigin string `json:"client_origin" mapstructure:"CLIENT_ORIGIN"`
	EmailFrom    string `json:"from"          mapstructure:"EMAIL_FROM"`
	Host         string `json:"smtp_host"     mapstructure:"SMTP_HOST"`
	Port         int    `json:"smtp_port"     mapstructure:"SMTP_PORT"`
	Username     string `json:"smtp_username" mapstructure:"SMTP_USERNAME"`
	Password     string `json:"smtp_password" mapstructure:"SMTP_PASSWORD"`
}

type Figma struct {
	FigmaToken      string `json:"figma_token"        mapstructure:"FIGMA_TOKEN"`
	FigmaBaseUrl    string `json:"figma_base_url"     mapstructure:"FIGMA_BASE_URL"`
	FigmaApiBaseUrl string `json:"figma_api_base_url" mapstructure:"FIGMA_API_BASE_URL"`
}

func LoadConfigAPI(path string) *Api {
	if path := strings.TrimSpace(path); path == "" {
		path = "."
	}

	viper.AddConfigPath(path)
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		log.Println("read config failed:", err.Error())
	}

	viper.SetDefault("ENV", "development")
	viper.SetDefault("APP_PORT", 8001)
	viper.SetDefault("APP_HOST", "127.0.0.1")

	api := &Api{}

	viper.Unmarshal(api)
	viper.Unmarshal(&api.Database)
	viper.Unmarshal(&api.SMTP)
	viper.Unmarshal(&api.Figma)

	return api
}
