package utils

import (
	"log"
	"os"
)

var (
	// EtherscanURI : etherscan api url
	EtherscanURI string
	// EtherscanAPIKey : n
	EtherscanAPIKey string
	// EthereumNetwork : n
	EthereumNetwork string
)

// GetEnv : the .env variable by key or its default
func GetEnv(key, fallback string) string {
	log.Printf("looking for env key: %s", key)
	value, exists := os.LookupEnv(key)
	if exists {
		return value
	}
	return fallback
}

// SetEtherscanEnvironment : sets etherscan.io environement
func SetEtherscanEnvironment() {
	EthereumNetwork = GetEnv("ETHEREUM_NETWORK", "mainnet")
	EtherscanAPIKey = GetEnv("ETHERSCAN_API_KEY", "")
	EtherscanURI = GetEnv("ETHERSCAN_URI", "")
}
