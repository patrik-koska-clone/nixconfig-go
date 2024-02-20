package main

import (
	"flag"
	"log"

	"github.com/patrik-koska-clone/nixconfig-go/pkg/collector"
	"github.com/patrik-koska-clone/nixconfig-go/pkg/config"
	githubadapter "github.com/patrik-koska-clone/nixconfig-go/pkg/githubadapter"
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
	log.Println("config read success")

	ga := githubadapter.New(c)

	base64Contents, downloadURLs, err := ga.SearchAndDownload(apiPageNumber, perPage, c.ConfigType)
	if err != nil {
		log.Fatalf("could not search api\n%v", err)
	}
	log.Println("got downloaded contents..")

	decodedContents, err := utils.DecodeFromB64(base64Contents)
	if err != nil {
		log.Fatalf("could not decode base64 contents\n%v", err)
	}
	log.Println("decoded base64 contents..")

	err = collector.PutFilesToDirectory(decodedContents, downloadURLs, c.ConfigType, c.OutputDir, apiPageNumber)
	if err != nil {
		log.Fatalf("could not write files to directory\n%v", err)
	}

}
