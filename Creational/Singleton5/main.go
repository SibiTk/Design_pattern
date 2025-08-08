package config

import (
	"fmt"
	"sync"
)

type Config struct {
	AppName string
	Port    int
	Env     string
}

var instance *Config
var lock = &sync.Mutex{}

func GetConfig() *Config {
	if instance == nil {
		
		lock.Lock()
		defer lock.Unlock()
	
		if instance == nil {
			fmt.Println("Loading config (mutex version)...")
			instance = &Config{
				AppName: "PaymentsService",
				Port:    8080,
				Env:     "production",
			}
		}
	}
	return instance
}

func main() {
	cfg1 := 																																																																																																																																																																																																																																																																																																																																																																																																																																																																											                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         																																																				   		                                                                 GetConfig()
	fmt.Println("App Name:", cfg1.AppName)

	cfg2 := config.GetConfig()
	fmt.Println("Env:", cfg2.Env)

	if cfg1 == cfg2 {
		fmt.Println("cfg1 and cfg2 are the same (Singleton works!)")
	}
}