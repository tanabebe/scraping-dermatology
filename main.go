package main

import (
	"log"
	"time"

	"github.com/sclevine/agouti"
	"gopkg.in/ini.v1"
)

// スクレイピング用のstruct
type ScrapingList struct {
	Url      string
	Email    string
	Password string
}

var Scraping ScrapingList

// iniファイルの読み込み
func init() {
	config, err := ini.Load("config.ini")

	if err != nil {
		log.Fatalln(err)
	}

	Scraping = ScrapingList{
		Url:      config.Section("web").Key("url").MustString(""),
		Email:    config.Section("login").Key("email").MustString(""),
		Password: config.Section("login").Key("password").MustString(""),
	}
}

func main() {

	driver := agouti.ChromeDriver()

	if err := driver.Start(); err != nil {
		log.Fatalln(err)
	}

	// close web driver
	defer func() {
		if err := driver.Stop(); err != nil {
			log.Fatalln(err)
		}
	}()

	target, err := driver.NewPage()

	if err != nil {
		log.Fatalln(err)
	}

	// close web browser
	defer func() {
		if err := target.CloseWindow(); err != nil {
			log.Fatalln(err)
		}
	}()

	if err := target.Navigate(Scraping.Url); err != nil {
		log.Fatalln(err)
	}

	time.Sleep(time.Second * 1)

	if err := target.FindByLink("今すぐ受付").Click(); err != nil {
		log.Fatalln(err)
	}

	if err := target.FindByID("user_email").Fill(Scraping.Email); err != nil {
		log.Fatalln(err)
	}

	if err := target.FindByID("user_password").Fill(Scraping.Password); err != nil {
		log.Fatalln(err)
	}

	if err := target.FindByName("commit").Submit(); err != nil {
		log.Fatalln(err)
	}

	time.Sleep(time.Second * 1)
}
