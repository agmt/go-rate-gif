package main

import (
	"fmt"
	"log"
	"rategif/config"
	"rategif/services/openexchange"
	"rategif/services/rategif"

	"github.com/gin-gonic/gin"
	"github.com/peterhellberg/giphy"
	"github.com/spf13/viper"
)

func main() {
	var c config.Config
	viper.SetConfigFile("rategif.yml")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	if err := viper.Unmarshal(&c); err != nil {
		log.Fatalf("[Config][Error] cannot read config file: %+v\n", err)
	}

	fmt.Printf("OE Key: '%s' Giphy: '%s'\n", c.Openexchange.ApiKey, c.Giphy.ApiKey)

	var api rategif.API

	api.OE = openexchange.New(c.Openexchange.ApiKey, c.Openexchange.Base)
	api.Giphy = giphy.NewClient()
	api.Giphy.APIKey = c.Giphy.ApiKey
	api.Giphy.Limit = 1

	api.Router = gin.New()

	api.Router.GET("/rate.gif", api.GetRateGif)

	api.Router.Run()
}
