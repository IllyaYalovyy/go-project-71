package main

import (
	"log"

	"github.com/IllyaYalovyy/go-project-71/config"
	"github.com/IllyaYalovyy/go-project-71/reverseproxy"
)

func main() {

	cfgPath := "local-config.json"

	cfg, err := config.FromJson(cfgPath)
	if err != nil {
		log.Fatalf("failed to read server configuration. Path: '%s'. Error: %v", cfgPath, err)
	}

	proxy, err := reverseproxy.New(cfg.SourceAddress, cfg.SourcePort, cfg.DestinationUrl)
	if err != nil {
		log.Fatalf("failed to create server. Destination URL: '%s'. Error: %v", cfg.DestinationUrl, err)
	}

	log.Printf("Starting proxy server. Listening on port: %d", cfg.SourcePort)
	err = proxy.Run()
	if err != nil {
		log.Fatalf("proxy server failed to start. Error: %v", err)
	}
}
