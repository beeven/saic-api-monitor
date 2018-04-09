package main

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// A struct represent configurations.
// api: api addresses
// interval: check interval in ms.
type configuration struct {
	API struct {
		BaseUrl           string `yaml:"baseUrl"`
		LoginPath         string `yaml:"loginPath"`
		SearchCompanyPath string `yaml:"searchCompanyPath"`
	}
	Login struct {
		Account  string `url:"account"`
		Password string `url:"password"`
	}
	Interval int
	Notify   struct {
		Url     string   `yaml:"url"`
		Mobiles []string `yaml:"mobiles"`
		Content string   `yaml:"content"`
	}
}

// FromFile Make configuration from file
func (c *configuration) fromFile(path string) error {
	bs, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(bs, c)
	if err != nil {
		return err
	}
	return nil
}
