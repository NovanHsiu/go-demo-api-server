@echo off
mkdir build
mkdir build\configs
mkdir build\docs
swag init --parseDepth 2
go build -o build/demo-server-api.exe main.go
copy /Y configs\config.json.tmp build\configs\config.json
copy /Y docs\swagger.json build\docs\swagger.json