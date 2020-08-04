package util

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func InitialiseFlags() {
	pflag.String("token", "", "bot token")

	pflag.String("colors.image_endpoint", "en0.io/-/sheepcolors", "image endpoint for colored sheep generation")

	pflag.String("db.host", "localhost", "database hostname")
	pflag.Int("db.port", 3306, "database port")
	pflag.String("db.username", "root", "database username")
	pflag.String("db.password", "", "database password")
	pflag.String("db.database", "sheep", "database name")
}

func InitialiseConfig() {
	viper.SetEnvPrefix("SHEEP")

	viper.SetConfigName("bot")

	viper.AddConfigPath("/etc/sheep")
	viper.AddConfigPath("$HOME/sheep")
	viper.AddConfigPath("./config")
	viper.AddConfigPath(".")

	pflag.Parse()

	if err := viper.BindPFlags(pflag.CommandLine); err != nil {
		panic(err)
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
		} else {
			// Config file was found but another error was produced
			panic(err)
		}
	} else {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}

	// watch config file for changes
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})

}
