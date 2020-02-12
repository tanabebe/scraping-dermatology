package main

import (
	"database/sql"
	"log"

	"github.com/tanabebe/scraping-dermatology/config"
	"github.com/tanabebe/scraping-dermatology/domain"
	"github.com/tanabebe/scraping-dermatology/domain/repository"
	"github.com/tanabebe/scraping-dermatology/internal/service"
	"gopkg.in/ini.v1"
)

// スクレイピング用のstruct
type ScrapingList struct {
	Url      string
	Email    string
	Password string
}

var Scraping domain.ScrapingList
var ConfigDatabase domain.DbConfig
var Db *sql.DB

func init() {

	// config.iniの値を読み込んでおく
	iniCfg, err := ini.Load("config/config.ini")

	if err != nil {
		log.Fatalf("iniファイル読み込みエラー、処理を終了します:%v", err)
	}

	// DB接続用の構造体
	ConfigDatabase = domain.DbConfig{
		DriverName:             iniCfg.Section("db").Key("driverName").MustString(""),
		InstanceConnectionName: iniCfg.Section("db").Key("instanceConnectionName").MustString(""),
		DatabaseUser:           iniCfg.Section("db").Key("databaseUser").MustString(""),
		Password:               iniCfg.Section("db").Key("password").MustString(""),
		DatabaseName:           iniCfg.Section("db").Key("databaseName").MustString(""),
	}

	// スクレイピング対象のパラメーターを保持しておく構造体
	Scraping = domain.ScrapingList{
		Url:      iniCfg.Section("web").Key("url").MustString(""),
		Email:    iniCfg.Section("login").Key("email").MustString(""),
		Password: iniCfg.Section("login").Key("password").MustString(""),
	}

	// DBとの接続開始
	Db, err := config.ConnectDb(ConfigDatabase)

	if err != nil {
		log.Fatalf("データベース接続に失敗しました:%v", err)
	}

	// 予約データが存在するかを確認する
	count := repository.CountByWaitReservation(Db)

	if err != nil {
		log.Fatal(err)
	}

	if count == 0 {
		log.Fatalf("予約待機データが存在しないため、処理を終了します:%v", err)
	}
}

func main() {
	service.RunScraping(Scraping)
}
