# scraping-dermatology

## 概要

病院予約受付の操作を以下の通り, 自動化する

1. 病院のWeb予約受付サイトを開く
2. 自身のアカウントでログインを行う
3. 予約受付確認を行う
4. 予約登録を行う

## 環境

| Name | version |
| ---- | ------- |
| macOS Mojave | 10.14.6  |
| GoogleChrome | 79.0.3945.79 |
| ChromeDriver | 79.0.3945.36 |

## 動作方法

1. `touch wating`
2. `cron`で`go run main.go`を起動 or 手動で起動

## cronサンプル

```
MAILTO = "your mail address"
0 8 * * * cd /your path/scraping-dermatology; bash -l -c 'go run /{your path}/scraping-dermatology/main.go'
```