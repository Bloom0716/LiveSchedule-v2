## 開発手順
dockerデーモンを起動してコンテナを立ち上げる
```
docker compose up -d
```
別のターミナルでコンテナ名を指定してコンテナ内での作業が可能
```
docker compose exec {コンテナ名} {コマンド}
```
作業を終えるとコンテナを停止する
```
docker compose down
```

## 各コンテナ対応ポート番号(仮)
- client: 3000
- main api: 8000
- discord bot api: 5000
- db: 5432