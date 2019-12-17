package main

import (
	"fmt"
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
		log.Fatalf("config.iniファイルの読み込みに失敗しました:%v",err)
	}

	Scraping = ScrapingList{
		Url:      config.Section("web").Key("url").MustString(""),
		Email:    config.Section("login").Key("email").MustString(""),
		Password: config.Section("login").Key("password").MustString(""),
	}

	if _, err := os.Stat("waiting"); err != nil {
		log.Fatalf("waitingファイルが見当たらないため、処理を終了しました:%v", err)
	}

	if err := fileDelete("img/*.png"); err != nil {
		log.Fatal(err)
	}

}

// 削除したいパス内のファイルを削除する
func fileDelete(pattern string) error {
	files, err := filepath.Glob(pattern)

	if err != nil {
		//log.Fatalf("削除対象のファイルパスに異常があります:%v",err)
		return fmt.Errorf("削除対象のファイルパスに異常があります:%v",err)
	}
	for _, file := range files {
		if err := os.Remove(file); err != nil {
			return fmt.Errorf("ファイル削除中にエラーが発生しました:%v",err)
		}
	}
	return nil
}

func main() {

	// 一度実行したらエラーが起きても必ず削除するようにする
	defer func() {
		if err := fileDelete("waiting"); err != nil{
			log.Printf("waitingファイルの削除に失敗しました:%v", err)
			return
		}
	}()

	driver := agouti.ChromeDriver(
		agouti.ChromeOptions(
			"args", []string{
				"--headless", // browserを非表示で実行
			}),
	)

	if err := driver.Start(); err != nil {
		log.Printf("WebDriverのstartに失敗しました:%v",err)
		return
	}

	// close web driver
	defer func() {
		if err := driver.Stop(); err != nil {
			log.Printf("WebDriverのcloseに失敗しました:%v", err)
		}
	}()

	target, err := driver.NewPage()

	if err != nil {
		log.Printf("WebDriverに対応するPageを返却出来ませんでした:%v",err)
		return
	}

	// close web browser
	defer func() {
		if err := target.CloseWindow(); err != nil {
			log.Printf("アクティブなブラウザを閉じる時にエラーが発生しました:%v",err)
		}
	}()

	if err := target.Navigate(Scraping.Url); err != nil {
		log.Printf("対象のWeb URLを開く事が出来ませんでした:%v", err)
		return
	}

	if err := target.Screenshot("img/Screen1.png"); err != nil {
		log.Printf("screen shot1の取得に失敗しました:%v",err)
		return
	}

	time.Sleep(time.Second * 1)

	if err := target.FindByLink("今すぐ受付").Click(); err != nil {
		log.Printf("対象リンクテキストのクリックが失敗しました:%v",err)
		return
	}

	if err := target.FindByID("user_email").Fill(Scraping.Email); err != nil {
		log.Printf("ログイン時のメールアドレス入力に失敗しました:%v",err)
		return
	}

	if err := target.FindByID("user_password").Fill(Scraping.Password); err != nil {
		log.Printf("ログイン時のパスワード入力に失敗しました:%v",err)
		return
	}

	if err := target.Screenshot("img/Screen2.png"); err != nil {
		log.Printf("screen shot2の取得に失敗しました:%v",err)
		return
	}

	if err := target.FindByName("commit").Submit(); err != nil {
		log.Printf("予約受付のログインに失敗しました:%v",err)
		return
	}

	if err := target.Screenshot("img/Screen3.png"); err != nil {
		log.Printf("screen shot3の取得に失敗しました:%v",err)
		return
	}

	time.Sleep(1 * time.Second)

	if err := target.FindByName("commit").Submit(); err != nil {
		log.Printf("予約受付確認に失敗しました:%v",err)
		return
	}

	if err := target.Screenshot("img/Screen4.png"); err != nil {
		log.Printf("screen shot4の取得に失敗しました:%v",err)
		return
	}

	time.Sleep(1 * time.Second)

	if err := target.FindByName("commit").Submit(); err != nil {
		log.Printf("予約受付登録に失敗しました:%v",err)
		return
	}

	if err := target.Screenshot("img/Screen5.png"); err != nil {
		log.Printf("screen shot5の取得に失敗しました:%v",err)
		return
	}

	time.Sleep(1 * time.Second)
}
