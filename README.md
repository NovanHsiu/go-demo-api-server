# Summary
用來展示或教學 Golang Web Server API 基本功能的專案

# Config
- configs/config.json

## common
一般設定
- port: http api port, ex: `3010`
- ssl_port: [optional] https api port, ex: `3011`
- tls_crt_path: [optional] 憑證檔路徑，ex: `"./ssl/server.key"`
- tls_key_path: [optional] 私密金鑰檔路徑，ex: `"./ssl/server.key"`

## db
資料庫設定
- type: DB類型，有效的類型為 `sqlite`, `mysql`, `postgres`，type為sqlite時不需設定其他參數
- host: DB Server 主機位置，通常為 IP 或 domain name
- port: DB Server 阜口，未設定則使用默認阜口
- dbname: DB Server 資料庫名稱
- user: DB login user 登入資料庫使用者
- sslmode: type postgres 是否起用 SSL 模式，不啟用:`disable`、啟用:`require`
- passwd: 資料庫密碼，數值為空字串則使用預設密碼 `acebiotekUniiForm`
- timezone: [optional] postgresql 使用的時區，ex:`America/Toronto`，預設:`Asia/Taipei`

## file
檔案相關設定
- static_file_dir: 靜態檔案儲存路徑