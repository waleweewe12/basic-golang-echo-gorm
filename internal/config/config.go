package config

import (
	"fmt"
	"net/http"

	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Configs struct {
	vn         *viper.Viper
	ConfigPath string
	Postgres   Postgres    `mapstructure:"postgres"`
	HttpConfig HttpConfigs `mapstructure:"http_config"`
}

type HttpConfigs struct {
	MaxIdleConnections        int `mapstructure:"max_idle_connections"`
	MaxConnectionsPerHost     int `mapstructure:"max_connections_per_host"`
	MaxIdleConnectionsPerHost int `mapstructure:"max_idle_connections_per_host"`
	Timeout                   int `mapstructure:"time_out"`
}

type Postgres struct {
	Host     string `mapstructure:"host"`
	Database string `mapstructure:"database"`
	User     string `mapstructure:"user"`
	Port     string `mapstructure:"port"`
	Password string `mapstructure:"password"`
	Uri      string `mapstructure:"uri"`
}

var DB *gorm.DB

func (config *Configs) InitAllConfigs() error {

	name := "config"
	//log.Infof("config file using : %s", name)

	config.ConfigPath = "./internal/config"

	vn := viper.New()
	vn.AddConfigPath(config.ConfigPath)
	vn.SetConfigName(name)
	config.vn = vn

	if err := vn.ReadInConfig(); err != nil {
		return err
	}

	if err := config.vn.Unmarshal(&config); err != nil {
		return err
	}

	if err := config.Postgres.connectPostgres(); err != nil {
		return errors.Wrap(err, "connect postgres error")
	}

	t := http.DefaultTransport.(*http.Transport).Clone()
	t.MaxIdleConns = config.HttpConfig.MaxIdleConnections
	t.MaxConnsPerHost = config.HttpConfig.MaxIdleConnectionsPerHost
	t.MaxIdleConnsPerHost = config.HttpConfig.MaxIdleConnectionsPerHost

	//log.Infof("all config loaded : %#v", config)
	fmt.Println("init configs success")
	return nil
}

func (pg *Postgres) connectPostgres() error {

	var err error
	// dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Bangkok",
	// 	pg.Host,
	// 	pg.User,
	// 	pg.Password,
	// 	pg.Database, pg.Port)

	DB, err = gorm.Open(postgres.Open(pg.Uri), &gorm.Config{})

	if err != nil {
		fmt.Println(err)
	}
	return nil
}

func GetDBInstance() *gorm.DB {
	return DB
}
