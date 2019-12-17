package main

import (
	"log"
	"os"
	"path/filepath"
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
var target agouti.WebDriver


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

	if _, err := os.Stat("waiting"); err != nil {
		log.Fatalf("Processing ended because the reserved file does not exist %v\n", err)
	}

	fileDelete("img/*.png")

}

// 削除したいパス内のファイルを削除する
func fileDelete(pattern string) {
	files, err := filepath.Glob(pattern)
	if err != nil {
		log.Fatalln(err)
	}

	for _, file := range files {
		if err := os.Remove(file); err != nil {
			log.Fatalln(err)
		}
	}
}

func main() {

	driver := agouti.ChromeDriver(
		agouti.ChromeOptions(
			"args", []string{
				"--headless", // browserを非表示で実行
			}),
	)

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

	if err := target.Screenshot("img/Screen1.png"); err != nil {
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

	if err := target.Screenshot("img/Screen2.png"); err != nil {
		log.Fatalln(err)
	}

	if err := target.FindByName("commit").Submit(); err != nil {
		log.Fatalln(err)
	}

	if err := target.Screenshot("img/Screen3.png"); err != nil {
		log.Fatalln(err)
	}

	time.Sleep(time.Second * 1)
}
