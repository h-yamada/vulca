# Vulca: Vulnerability Check API

vulca is web api for getting vuls scan result.
It use golang, gin web framework.

# やっていること
- vulsでscanした結果を、vuls.sqlite3から取得し、WebAPIとしてJsonで返却

# 使い方

```
$go get github.com/h-yamada/vulca
$go install github.com/h-yamada/vulca

$vulca --h
Usage of ./vulca:
  -cve-db-path string
    	cve-db-path (default "./cve.sqlite3")
  -vuls-db-path string
    	vuls-db-path (default "./vuls.sqlite3")

$vulca
[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:	export GIN_MODE=release
 - using code:	gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /cve/:cveno               --> github.com/h-yamada/vulca/app/handler.CveDetail (3 handlers)
[GIN-debug] GET    /server/:server           --> github.com/h-yamada/vulca/app/handler.ServerCveList (3 handlers)
[GIN-debug] GET    /serverlist/:cveno        --> github.com/h-yamada/vulca/app/handler.CveServerList (3 handlers)
[GIN-debug] GET    /scanlist                 --> github.com/h-yamada/vulca/app/handler.ScanList (3 handlers)
[GIN-debug] Listening and serving HTTP on :8000
```

# API

- CVE情報取得
 - GET    /cve/:cveno
- 指定したサーバーで検知されたCVEリスト取得
 - GET    /server/:server
- 指定したCVEを検知されているサーバーリスト取得
 - GET    /serverlist/:cveno
- 検知対象のサーバーリスト取得
 - GET    /scanlist




----
- rubotyでchatopsできる -> github.com/h-yamada/ruboty-vulca
