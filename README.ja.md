# Vulca: Vulnerability Check API

vulca is web api for getting vuls scan result.
It use golang, gin web framework.

# やっていること
- vulsでscanした結果を、vuls.sqlite3から取得しJsonで返却

# 使い方

```
go get github.com/h-yamada/vulca
go install github.com/h-yamada/vulca

vulca                                                                                       [~/dev/cve]
[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:	export GIN_MODE=release
 - using code:	gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /cve/:cveno               --> github.com/h-yamada/vulca/app/handler.CveDetail (3 handlers)
[GIN-debug] GET    /server/:server           --> github.com/h-yamada/vulca/app/handler.ServerCveList (3 handlers)
[GIN-debug] GET    /serverlist/:cveno        --> github.com/h-yamada/vulca/app/handler.CveServerList (3 handlers)
[GIN-debug] GET    /scanlist                 --> github.com/h-yamada/vulca/app/handler.ScanList (3 handlers)
[GIN-debug] Listening and serving HTTP on :8000
```

----
- rubotyでchatopsできる -> github.com/h-yamada/ruboty-vulca
