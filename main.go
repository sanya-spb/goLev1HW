package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"log"

	"github.com/sanya-spb/goLev1HW/utils"
	"github.com/sanya-spb/goLev1HW/utils/version"
)

type APP struct {
	Conf    utils.Config
	Version version.AppVersion
}

// var conf *utils.Config = new(utils.Config)
var MyApp *APP = new(APP)

func main() {
	MyApp.Version = *version.Version
	confPathPtr := flag.String("config", utils.GetEnv("APP_CONFIG", "no-config"), "Path to configuration file (*.toml|*.yaml)")
	confDebugPtr := flag.Bool("debug", utils.GetEnvBool("APP_DEBUG", false), "Output verbose debug information")
	flag.Parse()
	MyApp.Conf = *utils.LoadConfig(*confPathPtr, *confDebugPtr)
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
