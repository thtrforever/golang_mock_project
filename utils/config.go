package utils

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type CfgServer struct {
	Port uint `mapstructure:"port"`
}

type CfgDatabase struct {
	DriverName string `mapstructure:"driver_name"`
	Host       string `mapstructure:"host"`
	Port       uint   `mapstructure:"port"`
	Username   string `mapstructure:"username"`
	Password   string `mapstructure:"password"`
	DbName     string `mapstructure:"db_name"`
	SslMode    string `mapstructure:"ssl_mode"`
	TimeZone   string `mapstructure:"time_zone"`
}

type Config struct {
	Server   CfgServer   `mapstructure:"server"`
	Database CfgDatabase `mapstructure:"database"`
}

type GrpcConfig struct {
	Server   GrpcCfgServer   `mapstructure:"server"`
	Database GrpcCfgDatabase `mapstructure:"database"`
}

type GrpcCfgServer struct {
	Port uint `mapstructure:"port"`
}

type GrpcCfgDatabase struct {
	DriverName string `mapstructure:"driver_name"`
	Host       string `mapstructure:"host"`
	Port       uint   `mapstructure:"port"`
	Username   string `mapstructure:"username"`
	Password   string `mapstructure:"password"`
	DbName     string `mapstructure:"db_name"`
	SslMode    string `mapstructure:"ssl_mode"`
	TimeZone   string `mapstructure:"time_zone"`
}

func LoadConfigs(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&config)
	fmt.Printf("config is: %v\n", config)
	return
}

func LoadGrpcConfig(path string) (config GrpcConfig, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("grpc_config")
	viper.SetConfigType("yml")
	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&config)
	fmt.Printf("config is: %v\n", config)
	return
}

func GetDatabaseSourceName(cfgDatabase CfgDatabase) string {
	databaseSourceName := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=%v TimeZone=%v",
		cfgDatabase.Host,
		cfgDatabase.Port,
		cfgDatabase.Username,
		cfgDatabase.Password,
		cfgDatabase.DbName,
		cfgDatabase.SslMode,
		cfgDatabase.TimeZone,
	)
	log.Printf("databaseSourceName=%v\n", databaseSourceName)
	return databaseSourceName
}

func GetDatabaseSourceNameForGrpc(cfgDatabase GrpcCfgDatabase) string {
	databaseSourceName := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=%v TimeZone=%v",
		cfgDatabase.Host,
		cfgDatabase.Port,
		cfgDatabase.Username,
		cfgDatabase.Password,
		cfgDatabase.DbName,
		cfgDatabase.SslMode,
		cfgDatabase.TimeZone,
	)
	log.Printf("databaseSourceName=%v\n", databaseSourceName)
	return databaseSourceName
}
