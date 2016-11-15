package config

import v "github.com/spf13/viper"

// Load config from file and/or the environment.
func InitConfig() error {
	v.SetConfigName(".tekleader")
	v.SetConfigType("yaml")
	v.AddConfigPath("$HOME")
	v.AddConfigPath("/etc/")

	v.BindEnv("tek_timeout")
	v.BindEnv("tek_authkey")

	return v.ReadInConfig()
}
