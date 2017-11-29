package config

import (
	"sync"
	"fmt"
	"io/ioutil"
	"encoding/json"
)

type configFile struct {
	Token string `json:token`
	Weeb string `json:weeb`
}

const (
	CONFIG_FILE = "config.json"
)

var (
	config *configFile
	mutex = &sync.Mutex{}
)

func getConfig() configFile {
	mutex.Lock()
	defer mutex.Unlock()

	if config == nil { //check if config is "nil"
		//load config
		loadConfig()
	}

	return *config
}

func loadConfig(){
	fmt.Println("Loading Config...")
	raw, err := ioutil.ReadFile(CONFIG_FILE)
	if err != nil {
		fmt.Println("Couldn't find Config File!", err)
		panic(err)
	}
	json.Unmarshal(raw, &config)
}

func Get() configFile{
	return getConfig()
}
