package config

import (
	"github.com/spf13/viper"
	"log"
)

var Settings *Configuration

type Configuration struct {
	Server      ServerConfiguration
	Database    DatabaseConfiguration
	Jwt			JwtConfiguration
	Application ApplicationConfiguration
	Mail        MailConfiguration
	ExternalUrl ExternalUrlConfiguration
}

type ServerConfiguration struct {
	Port int
}

type DatabaseConfiguration struct {
	ConnectionUri string
}

type JwtConfiguration struct {
	TokenKey string
	TokenExp string
}

type ApplicationConfiguration struct {
	Name string
	Link string
	SendMails bool
}

type MailConfiguration struct {
	SmtpPort int
	SmtpPassword string
	SmtpUserName string
	SmtpServer string
	SenderEmail string
	SenderIdentity string
}

type ExternalUrlConfiguration struct {
	StockSymbols      string
	StockSymbolDetail string
}

func Init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	if err := viper.Unmarshal(&Settings); err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}
}