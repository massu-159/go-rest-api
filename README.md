```
#ローカル環境
## DBコンテナ起動
docker compose up -d

## migrateコマンド
GO_ENV=dev go run migrate/migrate.go

## コンテナに接続する
docker exec -u postgres -it [コンテナ名] bash

## コンテナ情報を取得する
docker ps

## Postgresqlに接続する
psql

## テーブル一覧を取得する
\dt

## カラム情報を取得する
\d [テーブル名]

## サーバーの起動
GO_ENV=dev go run main.go
```