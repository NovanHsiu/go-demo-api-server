# Summary
用來展示或教學 Golang Web Server API 基本功能的專案

# 安裝與啟動
## 1. 專案模組安裝
```bash
go install
go mod tidy
```

## 2. Swaggo 安裝方法
1. 安裝 swag
- Go 版本未滿 1.17: `go get -u github.com/swaggo/swag/cmd/swag`
- Go 1.17 以上版本: `go install github.com/swaggo/swag/cmd/swag`

2. 如果找不到 swag 指令，需將 /go/bin 資料夾加入環境變數
    - Linux
    ```bash
    echo "export PATH=$PATH:$HOME/go/bin" >> ~/.bashrc # 參考 go env 指令定義的 GOPATH 確認安裝路徑
    source ~/.bashrc
    ```
    - Windows
    ```bat
    : 參考 go env 指令定義的 GOPATH 確認安裝路徑
    setx PATH %PATH%C:\Go\bin;C:\Go\bin\bin
    ```

## 3. 建置與運行
```bat
: 運行建置腳本
build.bat
: 啟動執行檔，有需要的話可修改 build/configs/config.json 設定檔
build/demo-server-api.exe
```

# Config
## 使用說明
複製 configs/config.json.tmp 為 configs/config.json 作為設定擋使用

## 設定參數說明
### common
一般設定
- port: http api port, ex: `3010`
- ssl_port: [optional] https api port, ex: `3011`
- tls_crt_path: [optional] 憑證檔路徑，ex: `"./ssl/server.key"`
- tls_key_path: [optional] 私密金鑰檔路徑，ex: `"./ssl/server.key"`

### db
資料庫設定
- type: DB類型，有效的類型為 `sqlite`, `mysql`, `postgres`，type為sqlite時不需設定其他參數
- host: DB Server 主機位置，通常為 IP 或 domain name
- port: DB Server 阜口，未設定則使用默認阜口
- dbname: DB Server 資料庫名稱
- user: DB login user 登入資料庫使用者
- sslmode: type postgres 是否起用 SSL 模式，不啟用:`disable`、啟用:`require`
- passwd: 資料庫密碼，數值為空字串則使用預設密碼 `acebiotekUniiForm`
- timezone: [optional] postgresql 使用的時區，ex:`America/Toronto`，預設:`Asia/Taipei`

### file
檔案相關設定
- static_file_dir: 靜態檔案儲存路徑

# gomock
## 安裝 mockgen
1. 確認環境變數 `PATH` 有加入路徑 `$GOPATH/bin`，方可在終端機執行
2. 輸入指令安裝 gomock `go install github.com/golang/mock/mockgen@v1.6.0`

## 產生 mock repository code
1. 執行 internal/app/service interface 的 go:generate 註解產生 mock code，範例: `mockgen -destination automock/user_repository.go -package=automock . UserRepository`

# 模組使用方法說明
- [gin](https://github.com/gin-gonic/gin/blob/master/docs/doc.md): Gin Quick Start 文件，包含: Parameter usage, Upload files, Grouping routes, Model binding and validation 等重點功能使用方法與範例說明
- [gorm](https://gorm.io/docs/index.html): Gorm Quick Start
- [swaggo](https://github.com/swaggo/swag/blob/master/README.md): Go Swag 使用方法說明