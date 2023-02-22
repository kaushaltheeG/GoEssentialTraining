package main

import (
	"fmt"
	"log"
	"os" //operation system package 

	// "github.com/pkg/errors"
)

type Config struct {

}

func readConfig(path string) (*Config, error) {
	file, err := os.Open(path)

	if err != nil {
		//return nil, //errors.Wrap(err, "can't open configuration file") //errors.Wrap come from the pkg errors package 
	}
	
	defer file.Close()

	cfg := &Config{}

	return cfg, nil 
}

func setupLogging() {
	out, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE| os.O_WRONLY, 06) //last arg (06) is wrong 

	if err != nil {
		return 
	}

	log.SetOutput(out)
}

func main() {
	setupLogging()
	cfg, err := readConfig("/path/to/config.toml")

	if err != nil {
		//fmt.Fprint(os.Stderr, "error: %s\n", err) => comes back when pkg/errors is able to be properly installed 
		log.Printf("error: %+v", err) //%+v to print with log package which logs the stack trace when fails
		os.Exit(1) //exit program with a non zero value 
	}

	fmt.Println(cfg) 

}