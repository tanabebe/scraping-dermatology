# 病院予約の動作を自動化するプログラム

## 前提

* ローカルでの実行確認済み

##　動作方法

* watingファイルを同一階層に設置
* cronでgo run main.goを起動する


```
MAILTO = "your mail address"
2 0 8 * * * cd /your path/scraping-dermatology; bash -l -c 'go run /{your path}/scraping-dermatology/main.go'
```