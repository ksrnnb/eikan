## eikan (仮)
[![Codemagic build status](https://api.codemagic.io/apps/60108cf2cb2f385f48eaa129/my-workflow/status_badge.svg)](https://codemagic.io/apps/60108cf2cb2f385f48eaa129)

栄養管理できるスマホアプリ。
毎日食べたものを記録して、100点満点で数値化する。

## 環境
- mobile: Flutter
- API: Golang
- DB: MongoDB

## ライセンス
[MIT](LICENSE.md)

## 開発環境の立ち上げ
- Go
```bash
docker-compose up -d
```
- Flutter
```bash
cd flutter_eikan
flutter pub get
```
- 環境変数ファイル
  - `golang_eikan/configs/config.example.ini` を`config.ini`にリネームする。
  - `flutter_eikan/.env.example` を`.env`にリネームする。

## GOのビルド
```bash
docker-compose exec go ash
go run main.go
```