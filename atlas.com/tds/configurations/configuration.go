package configurations

import (
	"errors"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type Configuration struct {
	Topics []TopicConfiguration `yaml:"topics"`
}

type TopicConfiguration struct {
	Id   string `yaml:"id"`
	Name string `yaml:"name"`
}

func GetConfiguration() (*Configuration, error) {
	yamlFile, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
		return nil, err
	}

	c := &Configuration{}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
		return nil, err
	}

	return c, nil
}

func (c Configuration) GetTopicConfiguration(id string) (*TopicConfiguration, error) {
	for _, x := range c.Topics {
		if x.Id == id {
			return &x, nil
		}
	}
	return nil, errors.New(fmt.Sprintf("could not find by id: %s", id))
}
