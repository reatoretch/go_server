# goでマルチプレイゲーム用サーバ
4人ごとにルームが分かれます
## commandメモ
docker起動
```
docker-compose build

docker-compose up -d
```

## error対処
dockerで他のコンテナとポートが被ってる場合の対処
```
sudo lsfo -i:[bindしてるポート番号]
sudo kill [プロセス]
```
