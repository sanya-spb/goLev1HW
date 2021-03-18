package utils

import (
	"encoding/json"
	"errors"
	"log"
	"net"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/komkom/toml"
	"gopkg.in/yaml.v3"
)

const (
	_ = iota
	ERROR
	WARN
	NOTICE
)

//TODO: db_url (https://github.com/xo/dburl)
type Config struct {
	Debug    bool   `toml:"debug" yaml:"debug" json:"debug"`
	MyUrl    string `toml:"my_url" yaml:"my_url" json:"my_url"`
	Database struct {
		Host string `toml:"host" yaml:"host" json:"host"`
		Port int    `toml:"port" yaml:"port" json:"port"`
		User string `toml:"user" yaml:"user" json:"user"`
		Pass string `toml:"pass" yaml:"pass" json:"pass"`
		Ssl  bool   `toml:"ssl" yaml:"ssl" json:"ssl"`
	} `toml:"database" yaml:"database" json:"database"`
	Server struct {
		Bind     []string `toml:"bind" yaml:"bind" json:"bind"`
		Port     int      `toml:"port" yaml:"port" json:"port"`
		LogLevel int      `toml:"log_level" yaml:"log_level" json:"log_level"`
	} `toml:"server" yaml:"server" json:"server"`
}

func LoadConfig(cfgFile string, debug bool) *Config {
	var result *Config = new(Config)

	// Default values
	// result.Debug = false		// с этим параметром отдельная история ниже
	result.Database.Host = "127.0.0.1"
	result.Database.Port = 5432
	result.Database.Ssl = false
	result.Server.Bind = []string{"0.0.0.0"}
	result.Server.Port = 8888
	result.Server.LogLevel = WARN

	// Если есть файл конфига
	if _, err := os.Stat(cfgFile); !os.IsNotExist(err) {
		// Берем конфиг из файла и цепляем к структуре
		if f, err := os.Open(cfgFile); err != nil {
			log.Fatal(err)
		} else {
			defer f.Close()

			switch ext := strings.ToLower(filepath.Ext(cfgFile)); ext {
			case ".toml":
				if err := json.NewDecoder(toml.New(f)).Decode(&result); err != nil {
					log.Fatal(err)
				}
				// toml.Unmarshal([]byte(""), result)
			case ".yaml", ".yml":
				if err := yaml.NewDecoder(f).Decode(&result); err != nil {
					log.Fatal(err)
				}
			case ".json":
				if err := json.NewDecoder(f).Decode(&result); err != nil {
					log.Fatal(err)
				}
			default:
				log.Fatalf("Неизвестный формат файла: %s", cfgFile)
			}
		}
	} else if cfgFile == "no-config" {
		log.Println("Работаем без файла конфигурации (на дефолтных настройках)")
	} else {
		log.Fatalf("Файл конфигурации отсутсвует: %s", cfgFile)
	}

	// debug достанем из ENV и params (приоритет тут будет выше)
	result.Debug = debug

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
	if conf.MyUrl != "" && !IsURL(conf.MyUrl) {
		err := errors.New("Invalid my_url: " + conf.MyUrl)
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
	if conf.Server.LogLevel > NOTICE {
		conf.Server.LogLevel = NOTICE
	} else if conf.Server.LogLevel < ERROR {
		conf.Server.LogLevel = ERROR
	}
	return nil
}

func GetEnv(key string, defaultVal string) string {
	if envVal, ok := os.LookupEnv(key); ok {
		return envVal
	}
	return defaultVal
}

func GetEnvBool(key string, defaultVal bool) bool {
	if envVal, ok := os.LookupEnv(key); ok {
		envBool, err := strconv.ParseBool(envVal)
		if err == nil {
			return envBool
		}
	}
	return defaultVal
}

// UNUSED:
// func getEnvInt(key string, defaultVal int) int {
// 	if envVal, ok := os.LookupEnv(key); ok {
// 		envInt, err := strconv.ParseInt(envVal, 10, 64)
// 		if err == nil {
// 			return int(envInt)
// 		}
// 	}
// 	return defaultVal
// }
