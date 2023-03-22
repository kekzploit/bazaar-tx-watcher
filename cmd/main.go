package main

import (
	"fmt"
	"github.com/kekzploit/bazaar-tx-watcher/pkg/tx"
	"github.com/spf13/viper"
)

func main() {
	// initiate environment variable config file
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./../configs/")
	err := viper.ReadInConfig()
	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	tx.Txs(viper.Get("WALLET.URL").(string), viper.Get("MONGO.URI").(string)) // fetch transactions from zano daemon
}
