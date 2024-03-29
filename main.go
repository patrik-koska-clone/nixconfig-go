package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/patrik-koska-clone/nixconfig-go/pkg/collector"
	"github.com/patrik-koska-clone/nixconfig-go/pkg/config"
	"github.com/patrik-koska-clone/nixconfig-go/pkg/requests"
	"github.com/patrik-koska-clone/nixconfig-go/pkg/utils"
)

var (
	apiPageNumber int
	perPage       int
)

func init() {
	flag.IntVar(&apiPageNumber, "page", 1, "The page number to download")
	flag.IntVar(&perPage, "per-page", 100, "Amount of files to download per page")
	flag.Parse()
}

func main() {
	c, err := config.GetConfig("config.json")
	if err != nil {
		log.Fatalf("could not load config file. does it exist?\n%v", err)
	}

	baseURL := fmt.Sprintf("https://api.github.com/search/code?q=filename:%s+in:path&page=%d&per_page=%d",
		c.ConfigType,
		apiPageNumber,
		perPage)

	downloadURLs, htmlURLs, err := requests.GetNixConfigURLs(baseURL, c.Token)
	if err != nil {
		log.Fatalf("could not get nix download urls.\n%v", err)
	}
	log.Println("got download urls, extracting...")

	for _, u := range downloadURLs {
		log.Println("collecting from", u)
	}

	contents, err := requests.GetNixConfigContents(downloadURLs, c.Token)
	if err != nil {
		log.Fatalf("could not get base64 file contents.\n%v", err)
	}

	log.Println("accessing base64 file contents from json...")

	decodedContents, err := utils.DecodeFromB64(contents)
	if err != nil {
		log.Fatalf("error decoding base64 contents from response json.\n%v", err)
	}
	log.Println("decoded base64 file contents...")

	err = collector.PutFilesToDirectory(decodedContents, htmlURLs, c.ConfigType, c.OutputDir, apiPageNumber)
	if err != nil {
		log.Fatalf("error writing nix files to directory.\n%v", err)
	}
	log.Printf("nix files collected successfully to %s.\n", c.OutputDir)
}
