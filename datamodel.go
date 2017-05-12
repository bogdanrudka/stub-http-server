package main

import (
	"path/filepath"
	"encoding/json"
	"strings"
	"io/ioutil"
	"fmt"
)

type Configuration struct {
	Settings Settings
	Mappings []Mapping
}

func (c Configuration) String() string {
	return fmt.Sprintf("{\n\tsettings: %s,\n \tmappings: %s\n}", c.Settings, c.Mappings)
}

type Settings struct {
	BaseURL string
	Port    int
}

func (s Settings) String() string {
	return fmt.Sprintf("{\n\t\tbaseURL: %s,\n\t\tport: %d\n\t}", s.BaseURL, s.Port)
}

type Mapping struct {
	Request  Request
	Response Response
}

type Request struct {
	Method string
	Path   string
}

type Response struct {
	Status int
	Body   interface{}
	Header map[string]string
}

func InitConfig() *Configuration {
	cpath := filepath.Join("config.json")
	file, err := ioutil.ReadFile(cpath)
	if err != nil {
		fmt.Printf("Error while opening file %s", cpath)
		return &Configuration{}
	}
	dec := json.NewDecoder(strings.NewReader(string(file)))
	var conf *Configuration
	err = dec.Decode(&conf)

	if err != nil {
		fmt.Printf("Error while parsing json %v", err)
		return &Configuration{}
	}
	return conf
}

//func main() {
//	var conf Configuration
//	for _, m := range conf.Mappings {
//		http.HandleFunc(m.Request.Path)
//	}
//}
