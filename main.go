package main

import (
	"fmt"
	"log"
	"nixconfig-go/pkg/collector"
	"nixconfig-go/pkg/config"
	"nixconfig-go/pkg/requests"
	"nixconfig-go/pkg/utils"
)

func main() {
	c, err := config.GetConfig("config.json")
	if err != nil {
		log.Fatalf("could not load config file. does it exist?\n%v", err)
	}

	baseURL := fmt.Sprintf("https://api.github.com/search/code?q=filename:%s+in:path", c.ConfigType)

	downloadURLs, err := requests.GetNixConfigURLs(baseURL, c.Token)
	if err != nil {
		log.Fatalf("could not get nix download urls.\n%v", err)
	}
	log.Println("got download urls, extracting...")

	for _, u := range downloadURLs {
		log.Println("collecting from", u)
	}

	contents, err := requests.GetNixConfigContentsConcurrently(downloadURLs, c.Token)
	if err != nil {
		log.Fatalf("could not get base64 file contents.\n%v", err)
	}

	log.Println("accessing base64 file contents from json...")

	decodedContents, err := utils.DecodeFromB64(contents)
	if err != nil {
		log.Fatalf("error decoding base64 contents from response json.\n%v", err)
	}
	log.Println("decoded base64 file contents...")

	err = collector.PutFilesToDirectory(decodedContents, c.ConfigType, c.OutputDir)
	if err != nil {
		log.Fatalf("error writing nix files to directory.\n%v", err)
	}
	log.Printf("nix files collected successfully to %s.\n", c.OutputDir)
}
