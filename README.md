# 病院予約の動作を自動化するプログラム

## 環境

| Name | version |
| ---- | ------- |
| macOS Mojave | 10.14.6  |
| GoogleChrome | 79.0.3945.79 |
| ChromeDriver | 79.0.3945.36 |

## 動作方法

* watingファイルを同一階層に設置
* cronでgo run main.goを起動する

## cronサンプル

```
MAILTO = "your mail address"
0 8 * * * cd /your path/scraping-dermatology; bash -l -c 'go run /{your path}/scraping-dermatology/main.go'
```