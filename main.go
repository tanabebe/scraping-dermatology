package main

import (
	"log"

	"github.com/sclevine/agouti"
	"gopkg.in/ini.v1"
)

// スクレイピング用のstruct
type ScrapingList struct {
	Url string
	Id string
	Password string
}

var Scraping ScrapingList

// initファイルの読み込み
func init() {
	config, err := ini.Load("config.ini")

	if err != nil {
		log.Fatalf("Could not load ini file %v\n",err)
	}

	Scraping := ScrapingList{
		Url: config.Section("web").Key("url").MustString(""),
		Id: config.Section("login").Key("id").MustString(""),
		Password: config.Section("login").Key("password").MustString(""),
	}
}

func main() {

	driver := agouti.ChromeDriver()

	defer driver.Stop()

	if err := driver.Start(); err != nil {
		log.Fatalf("Error in WebDiver %v\n", err)
	}

	targetWeb, err  := driver.NewPage()
	if err != nil {
		log.Fatalln(err)
	}

	if err := targetWeb.Navigate(Scraping.Url); err != nil {
		log.Fatalln(err)
	}
}
