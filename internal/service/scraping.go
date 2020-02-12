package service

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/sclevine/agouti"
	"github.com/tanabebe/scraping-dermatology/domain"
)

func FileDelete(pattern string) error {
	files, err := filepath.Glob(pattern)

	if err != nil {
		return fmt.Errorf("削除対象のファイルパスに異常があります:%v", err)
	}
	for _, file := range files {
		if err := os.Remove(file); err != nil {
			return fmt.Errorf("ファイル削除中にエラーが発生しました:%v", err)
		}
	}
	return nil
}

func RunScraping(scraping domain.ScrapingList) {

	if err := FileDelete("img/*.png"); err != nil {
		log.Fatal(err)
	}

	driver := agouti.ChromeDriver(
	//agouti.ChromeOptions(
	//"args", []string{
	//	"--headless", // browserを非表示で実行
	//}),
	)

	if err := driver.Start(); err != nil {
		log.Printf("WebDriverのstartに失敗しました:%v", err)
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
		log.Printf("WebDriverに対応するPageを返却出来ませんでした:%v", err)
		return
	}

	if err := target.Navigate(scraping.Url); err != nil {
		log.Printf("対象のWeb URLを開く事が出来ませんでした:%v", err)
		return
	}

	if err := target.Screenshot("img/Screen1.png"); err != nil {
		log.Printf("screen shot1の取得に失敗しました:%v", err)
		return
	}

	time.Sleep(time.Second * 1)

	if err := target.FindByLink("今すぐ受付").Click(); err != nil {
		log.Printf("対象リンクテキストのクリックが失敗しました:%v", err)
		return
	}

	if err := target.FindByID("user_email").Fill(scraping.Email); err != nil {
		log.Printf("ログイン時のメールアドレス入力に失敗しました:%v", err)
		return
	}

	if err := target.FindByID("user_password").Fill(scraping.Password); err != nil {
		log.Printf("ログイン時のパスワード入力に失敗しました:%v", err)
		return
	}

	if err := target.Screenshot("img/Screen2.png"); err != nil {
		log.Printf("screen shot2の取得に失敗しました:%v", err)
		return
	}

	if err := target.FindByName("commit").Submit(); err != nil {
		log.Printf("予約受付のログインに失敗しました:%v", err)
		return
	}

	if err := target.Screenshot("img/Screen3.png"); err != nil {
		log.Printf("screen shot3の取得に失敗しました:%v", err)
		return
	}

	time.Sleep(1 * time.Second)

	if err := target.FindByName("commit").Submit(); err != nil {
		log.Printf("予約受付確認に失敗しました:%v", err)
		return
	}

	if err := target.Screenshot("img/Screen4.png"); err != nil {
		log.Printf("screen shot4の取得に失敗しました:%v", err)
		return
	}

	time.Sleep(1 * time.Second)

	if err := target.FindByName("commit").Submit(); err != nil {
		log.Printf("予約受付登録に失敗しました:%v", err)
		return
	}

	if err := target.Screenshot("img/Screen5.png"); err != nil {
		log.Printf("screen shot5の取得に失敗しました:%v", err)
		return
	}

	// close web browser
	if err := target.CloseWindow(); err != nil {
		log.Printf("アクティブなブラウザを閉じる時にエラーが発生しました:%v", err)
	}
}
