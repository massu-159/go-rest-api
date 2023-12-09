# go-rest-api

GoのAPIを学習するため、Todoアプリを作成。

フロントエンドにはReactを利用。

クリーンアーキテクチャを採用。

### ・ backend

GoのREST APIフレームワーク Echoを利用。

O/RマッパーにはGORMを利用。

CSRFトークンを使ったセキュリティ対策

JWT認証

PostmanでAPIテスト

DBはDockerでPostgreSQL環境を構築。

パスワードのハッシュ化。

urlはこちら
https://github.com/massu-159/go-rest-api

frontendはこちら
https://github.com/massu-159/React-Go-front

## 目次
1. コマンド
2. アプリケーションの仕様

## 1. コマンド
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

## 2. アプリケーションの仕様

### 2-1. 仕様
- 認証
  - サインアップ
  - ログイン
  - ログアウト
- Todo
  - Todo一覧表示
  - Todo新規登録
  - Todo更新処理
  - Todo削除処理

### 2-2. 構成技術
```
	github.com/go-ozzo/ozzo-validation/v4 v4.3.0
	github.com/golang-jwt/jwt/v4 v4.4.3
	github.com/joho/godotenv v1.5.1
	github.com/labstack/echo-jwt/v4 v4.1.0
	github.com/labstack/echo/v4 v4.11.3
	golang.org/x/crypto v0.14.0
	gorm.io/driver/postgres v1.5.4
	gorm.io/gorm v1.25.5

	github.com/asaskevich/govalidator v0.0.0-20200108200545-475eaeb16496
	github.com/golang-jwt/jwt v3.2.2+incompatible
	github.com/jackc/pgpassfile v1.0.0
	github.com/jackc/pgservicefile v0.0.0-20221227161230-091c0ba34f0a
	github.com/jackc/pgx/v5 v5.4.3
	github.com/jinzhu/inflection v1.0.0
	github.com/jinzhu/now v1.1.5
	github.com/labstack/gommon v0.4.0
	github.com/mattn/go-colorable v0.1.13
	github.com/mattn/go-isatty v0.0.19
	github.com/valyala/bytebufferpool v1.0.0
	github.com/valyala/fasttemplate v1.2.2
	golang.org/x/net v0.17.0
	golang.org/x/sys v0.13.0
	golang.org/x/text v0.13.0
	golang.org/x/time v0.3.0
```
