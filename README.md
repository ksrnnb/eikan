## eikan (仮)
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

## GOのビルド
```bash
docker-compose exec go ash
go run main.go
```