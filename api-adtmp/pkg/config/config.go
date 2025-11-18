package config

import (
	"log/slog"
	"os"
	"sync"
	"time"

	"gopkg.in/yaml.v3"
)

var (
	con  *config
	once sync.Once
)

type config struct {
	Address      string        `yaml:"address"`       // 应用监听地址
	ReadTimeout  time.Duration `yaml:"read_timeout"`  // 应用请求超时时间
	WriteTimeout time.Duration `yaml:"write_timeout"` // 应用响应超时时间
	LogLevel     string        `yaml:"log_level"`     // 日志输出级别
	LogFormat    string        `yaml:"log_format"`    // 日志输出格式
	LogSource    bool          `yaml:"log_source"`    // 日志源文件是否显示
	SecretKey    string        `yaml:"secret_key"`    // token 密钥
	ExpiredTime  time.Duration `yaml:"expired_time"`  // token 过期时间
	ServerKey    string        `yaml:"server_key"`    // 应用证书key
	ServerCrt    string        `yaml:"server_crt"`    // 应用这书密钥
	DBUrl        string        `yaml:"db_url"`        // 数据库连接地址
	DBIdleConn   int           `yaml:"db_idle_conn"`  // 数据最大空闲连接数
	DBOpenConn   int           `yaml:"db_open_conn"`  // 数据库最大打开连接数
	DBLifetime   time.Duration `yaml:"db_life_time"`  // 数据库连接最大存活时间
	DBIdleTime   time.Duration `yaml:"db_idle_time"`  // 数据库最大空闲连接时间
}

// 默认配置
func defaultConfig() *config {
	return &config{
		Address:      ":8080",
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		LogLevel:     "info",
		LogFormat:    "text",
		LogSource:    false,
		SecretKey:    "",
		ExpiredTime:  24 * time.Hour,
		ServerKey:    "",
		ServerCrt:    "",
		DBUrl:        "",
		DBIdleConn:   10,
		DBOpenConn:   100,
		DBLifetime:   time.Hour,
		DBIdleTime:   30 * time.Minute,
	}
}

// 获取配置项
func Get() *config {
	once.Do(func() {
		// 首先加载默认配置
		con = defaultConfig()

		// 尝试加载配置文件，如果失败则使用默认配置
		if err := load(); err != nil {
			slog.Warn("使用默认配置启动应用", slog.String("Error", err.Error()))
		}

		// 验证关键配置
		validateConfig()
	})
	return con
}

// 加载配置文件
func load() error {
	configFile := "config.yaml"

	// 检查配置文件是否存在
	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		return &os.PathError{Op: "open", Path: configFile, Err: err}
	}

	data, err := os.ReadFile(configFile)
	if err != nil {
		return err
	}

	// 创建临时配置对象，只覆盖配置文件中存在的字段
	tempConfig := defaultConfig()
	if err := yaml.Unmarshal(data, tempConfig); err != nil {
		return err
	}

	// 使用配置文件中的值更新当前(默认)配置
	con = tempConfig
	// slog.Info("配置文件加载成功", slog.String("ConfigFile", configFile))
	return nil
}

// 验证配置
func validateConfig() {
	if con.SecretKey == "" {
		slog.Error("请先配置应用 secret_key 字段")
		os.Exit(-3)
	}

	if con.DBUrl == "" {
		// postgres://user:pass@localhost:5432/dbname?sslmode=disable
		slog.Error("请先配置数据库连接信息 db_url 字段")
		os.Exit(-4)
	}
}
