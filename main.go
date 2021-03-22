package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"log"

	"github.com/sanya-spb/goLev1HW/utils/config"
	"github.com/sanya-spb/goLev1HW/utils/version"
)

type APP struct {
	Conf    config.Config
	Version version.AppVersion
}

// var conf *config.Config = new(config.Config)
var MyApp *APP = new(APP)

func main() {
	MyApp.Version = *version.Version
	confPathPtr := flag.String("config", config.GetEnv("APP_CONFIG", "no-config"), "Path to configuration file (*.toml|*.yaml)")
	confDebugPtr := flag.Bool("debug", config.GetEnvBool("APP_DEBUG", false), "Output verbose debug information")
	flag.Parse()
	MyApp.Conf = *config.LoadConfig(*confPathPtr, *confDebugPtr)
	if MyApp.Conf.Debug {
		if b, err := json.Marshal(MyApp.Conf); err == nil {
			var out bytes.Buffer
			if err := json.Indent(&out, b, "", "\t"); err == nil {
				fmt.Println(out.String())
			} else {
				// сюда по идее мы никогда не придем, мы же по сути из json делаем json
				log.Fatal("Что-то пошло не так..")
			}
		} else {
			// сюда по идее мы никогда не придем, мы же по сути из json делаем json
			log.Fatal("Что-то пошло не так..")
		}
	}
	fmt.Printf("version: %+v\n", MyApp.Version)
}
