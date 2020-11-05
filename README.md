# goでマルチプレイゲーム用サーバ
4人ごとにルームが分かれます
## commandメモ
docker起動
```
./startup.sh
```

## error対処
dockerで他のコンテナとポートが被ってる場合の対処
```
sudo lsof -i:[bindしてるポート番号]
sudo kill [プロセス]
```

## Tagでのversion指定について
基本的には[Semantic Versioning](https://semver.org/lang/ja/)に従います。  
developブランチでリリース前のテストを行う場合はversion末尾に"-test"を付けましょう  
```
ex) git tag v1.0.0-test
```
