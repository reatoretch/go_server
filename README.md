# goでマルチプレイゲーム用サーバ
4人ごとにルームが分かれます
## commandメモ
docker起動
```
docker-compose up -d --build --scale app=2
```

## error対処
dockerで他のコンテナとポートが被ってる場合の対処
```
sudo lsof -i:[bindしてるポート番号]
sudo kill [プロセス]
```
