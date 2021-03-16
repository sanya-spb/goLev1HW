package utils

import (
	"encoding/json"
	"errors"
	"flag"
	"log"
	"net"
	"net/url"
	"os"
	"strconv"

	"github.com/komkom/toml"
)

const (
	_ = iota
	ERROR
	WARN
	NOTICE
)

//TODO: db_url (https://github.com/xo/dburl)
type Config struct {
	Debug    bool
	My_url   string
	Database struct {
		Host string
		Port int
		User string
		Pass string
		Ssl  bool
	}
	Server struct {
		Bind      []string
		Port      int
		Log_level int
	}
}

func LoadConfig(cfgFile string) *Config {
	var result *Config = new(Config)

	// Default values
	result.Debug = false
	result.Database.Host = "127.0.0.1"
	result.Database.Port = 5432
	result.Database.Ssl = false
	result.Server.Bind = []string{"0.0.0.0"}
	result.Server.Port = 8888
	result.Server.Log_level = WARN

	// Если есть файл конфига
	if _, err := os.Stat(cfgFile); !os.IsNotExist(err) {
		// Берем конфиг из файла и цепляем к структуре
		if f, err := os.Open(cfgFile); err != nil {
			log.Fatal(err)
		} else {
			defer f.Close()
			if err := json.NewDecoder(toml.New(f)).Decode(&result); err != nil {
				log.Fatal(err)
			}
		}
	}

	// Теперь посмотрим в ENV и params (приоритет тут будет выше)
	// BUG: такая схема не заработала:
	// result.Debug = *flag.Bool("debug", getEnvBool("DEBUG", result.Debug), "Output verbose debug information")
	// WORKAROUND: сначала возьмем из ENV
	result.Debug = getEnvBool("DEBUG", result.Debug)
	result.Server.Port = getEnvInt("SERVER_PORT", result.Server.Port)
	result.My_url = getEnv("MY_URL", result.My_url)
	// WORKAROUND: теперь возьмем из params
	flag.BoolVar(&result.Debug, "debug", result.Debug, "Output verbose debug information")
	flag.IntVar(&result.Server.Port, "srv-port", result.Server.Port, "Server listen port")
	flag.StringVar(&result.My_url, "my-url", result.My_url, "testing URL")
	flag.Parse()

	// Теперь все проверим и пофиксим бред, по возможности
	if err := testConfig(result); err != nil {
		log.Fatal(err)
	}

	return result
}

// Check if host is valid IPv4 address (хотя на регулярках быстрее..)
func IsIPv4Net(host string) bool {
	return net.ParseIP(host) != nil
}

// Check if URL is valid (так себе проверка, но велосипед изобретать не хочу, лучше уж потом на коннекте ошибки ловить..)
func IsURL(str string) bool {
	if _, err := url.ParseRequestURI(str); err != nil {
		return false
	}
	return true
}

// test for errors and fix
func testConfig(conf *Config) error {
	// check My_url
	if conf.My_url != "" && !IsURL(conf.My_url) {
		err := errors.New("Invalid my_url: " + conf.My_url)
		return err
	}
	// check Database.Host
	if !IsIPv4Net(conf.Database.Host) {
		err := errors.New("Invalid Database.Host: " + conf.Database.Host)
		return err
	}
	// check Server.Bind
	for _, host := range conf.Server.Bind {
		if !IsIPv4Net(host) {
			err := errors.New("Invalid Server.Host: " + host)
			return err
		}
	}
	// fix Server.Log_level
	if conf.Server.Log_level > NOTICE {
		conf.Server.Log_level = NOTICE
	} else if conf.Server.Log_level < ERROR {
		conf.Server.Log_level = ERROR
	}
	return nil
}

func getEnv(key string, defaultVal string) string {
	if envVal, ok := os.LookupEnv(key); ok {
		return envVal
	}
	return defaultVal
}

func getEnvBool(key string, defaultVal bool) bool {
	if envVal, ok := os.LookupEnv(key); ok {
		envBool, err := strconv.ParseBool(envVal)
		if err == nil {
			return envBool
		}
	}
	return defaultVal
}

func getEnvInt(key string, defaultVal int) int {
	if envVal, ok := os.LookupEnv(key); ok {
		envInt, err := strconv.ParseInt(envVal, 10, 64)
		if err == nil {
			return int(envInt)
		}
	}
	return defaultVal
}
