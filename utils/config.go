package utils

import (
	"log"
	"os"

	gap "github.com/muesli/go-app-paths"
	"github.com/spf13/viper"
)

func IsPlayerNameMissing() bool {
	name := viper.GetString(PlayerNameKey)
	return name == ""
}

func IsAuthKeyMissing() bool {
	key := viper.GetString(AuthKey)
	return key == ""
}

func createConfigDir(path string) error {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return os.Mkdir(path, 0o770)
		}
		return err
	}
	return nil
}

func SetupPath() string {
	scope := gap.NewScope(gap.User, "league-predictions")
	dirs, err := scope.ConfigDirs()
	if err != nil {
		log.Fatal(err)
	}
	var configPath string
	if len(dirs) > 0 {
		configPath = dirs[0]
	} else {
		configPath, _ = os.UserHomeDir()
	}
	if err := createConfigDir(configPath); err != nil {
		log.Fatal(err)
	}
	return configPath
}

func InitConfig() {
	configPath := SetupPath()
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(configPath)
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			viper.Set(PlayerNameKey, "")
			if err := viper.SafeWriteConfig(); err != nil {
				log.Println("Error creating config file:", err)
			}
		} else {
			log.Println("Error reading config file:", err)
		}
	}
}

func isValidConfigKey(key string) bool {
	switch key {
	case PlayerNameKey, AuthKey:
		return true
	}
	return false
}

func SaveConfig(key string, value string) {
	if !isValidConfigKey(key) {
		log.Fatalln("Invalid config key:", key)
	}
	viper.Set(key, value)
	if err := viper.WriteConfig(); err != nil {
		log.Println("Error saving config:", err.Error())
	}
}
