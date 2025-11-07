package config

import (
	"log/slog"
	"os"
	"time"

	"go.yaml.in/yaml/v2"
)

var con *config

type config struct {
	Address      string        `yaml:"address" default:"127.0.0.1:8085"` // 应用监听地址
	ReadTimeout  time.Duration `yaml:"read_timeout" default:"60s"`       // 应用请求超时时间
	WriteTimeout time.Duration `yaml:"write_timeout" default:"60"`       // 应用响应超时时间
	LogLevel     string        `yaml:"log_level" default:"info"`         // 日志输出级别
	LogFormat    string        `yaml:"log_format" default:"text"`        // 日志输出格式
	SecretKey    string        `yaml:"secret_key" default:"1adnfdjkfa"`  // token 密钥
	ExpiredTime  time.Duration `yaml:"expired_time" default:"60m"`       // token 过期时间
	DBUrl        string        `yaml:"db_url"`                           // 数据库连接地址
	DBIdleConn   int           `yaml:"db_idle_conn" default:"20"`        // 数据最大空闲连接数
	DBOpenConn   int           `yaml:"db_open_conn" default:"60"`        // 数据库最大打开连接数
	DBLifetime   time.Duration `yaml:"db_life_time" default:"60m"`       // 数据库连接最大存活时间
	DBIdleTime   time.Duration `yaml:"db_idle_time" default:"60m"`       // 数据库最大空闲连接时间
}

// 获取配置项
func Get() *config {
	if con != nil {
		return con
	}
	err := load()
	if err != nil {
		os.Exit(-3)
	}
	return con
}

// 加载配置文件
func load() error {
	data, err := os.ReadFile("config.yaml")
	if err != nil {
		slog.Error("加载配置文件失败", slog.String("Error", err.Error()))
		return err
	}

	err = yaml.Unmarshal(data, &con)
	if err != nil {
		slog.Error("解析配置文件失败", slog.String("Error", err.Error()))
		return err
	}
	return nil
}
