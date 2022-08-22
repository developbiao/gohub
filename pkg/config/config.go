package config

import (
	"github.com/spf13/cast"
	viperlib "github.com/spf13/viper" // 自定义包名，避免与内置 viper
	"gohub/pkg/helpers"
	"os"
)

// viper instance
var viper *viperlib.Viper

// ConfigFunc dynamic
type ConfigFunc func() map[string]interface{}

// ConfigFuncs ConfigFunc loading on array, after generate loadConfig dynamic
var ConfigFuncs map[string]ConfigFunc

func init() {
	// 1. Initialization Viperlib
	viper = viperlib.New()
	// 2. Config type support ["json", "toml", "yaml", "yaml", "properties"
	// "props", "prop", "env", "dotenv"]
	viper.SetConfigType("env")
	// 3. Environment config relative path relative to main.go
	viper.AddConfigPath(".")
	// 4. Set environment prefix differentiate go environment
	viper.SetEnvPrefix("appenv")
	// 5. Read environment (support flags)
	viper.AutomaticEnv()

	ConfigFuncs = make(map[string]ConfigFunc)
}

// InitConfig initialization config information
func InitConfig(env string) {
	// 1. Load config environment
	loadEnv(env)
	// 2. Register config information
	loadConfig()
}

// loadEnv
func loadConfig() {
	for name, fn := range ConfigFuncs {
		viper.Set(name, fn())
	}
}

// loadEnv
func loadEnv(envSuffix string) {
	// Default load .env file, if give --env=name otherwise load .env.name file name
	envPath := ".env"
	if len(envSuffix) > 0 {
		filepath := ".env." + envSuffix
		if _, err := os.Stat(filepath); err == nil {
			envPath = filepath
		}
	}

	// load env
	viper.SetConfigName(envPath)
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	// Watch .env file change reload
	viper.WatchConfig()
}

// Env Read env support provider default value
func Env(envName string, defaultValue ...interface{}) interface{} {
	if len(defaultValue) > 0 {
		return internalGet(envName, defaultValue[0])
	}
	return internalGet(envName)
}

// Add config
func Add(name string, configFn ConfigFunc) {
	ConfigFuncs[name] = configFn
}

// Get config
// First parameter allow use "dot" e.g: app.name
// Second parameter allow default value
func Get(path string, defaultValue ...interface{}) string {
	return GetString(path, defaultValue)
}

// internalGet
func internalGet(path string, defaultValue ...interface{}) interface{} {
	// if config not exists read default value
	if !viper.IsSet(path) || helpers.Empty(viper.Get(path)) {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		return nil
	}
	return viper.Get(path)
}

// GetString get string type config  value
func GetString(path string, defaultValue ...interface{}) string {
	return cast.ToString(internalGet(path, defaultValue...))
}

// GetInt get int type config  value
func GetInt(path string, defaultValue ...interface{}) int {
	return cast.ToInt(internalGet(path, defaultValue...))
}

// GetInt64 get int64 type config value
func GetInt64(path string, defaultValue ...interface{}) int64 {
	return cast.ToInt64(internalGet(path, defaultValue...))
}

// GetUint get uint type config value
func GetUint(path string, defaultValue ...interface{}) uint {
	return cast.ToUint(internalGet(path, defaultValue...))
}

// GetFloat64 get float64 type config value
func GetFloat64(path string, defaultValue ...interface{}) float64 {
	return cast.ToFloat64(internalGet(path, defaultValue...))
}

// GetBool get bool type config value
func GetBool(path string, defaultValue ...interface{}) bool {
	return cast.ToBool(internalGet(path, defaultValue...))
}

// GetStringMapString get structure data
func GetStringMapString(path string) map[string]string {
	return viper.GetStringMapString(path)
}
