## Go言語テスト
Go言語のテストコードサンプル集です

## コマンド

```
$ go test -run keyword

$ go test -v //各関数の結果と実行時間を表示

$ go test -v

$ go test -v -run Keyword

$ go test -bench .

$ go test -c //バイナリ生成　-o {NAME} //出力名を定義、指定しないとpackage名となる

$ go test -i //依存パッケージのインストール、テストは実行されない

$ go test -json //JSON形式で吐き出される

$ go test -exec "go tool test2json" // 詳細ログがいらない場合はこっち

$ go test -o {NAME}//バイナリを吐くテストは実行される

$ go test -failfast //1つでも失敗するとそれ移行実行しない

$ go test -timeout 5 //タイムアウト、デフォルトは10分、0を指定でタイムアウト無効化

$ cd test7/
$ go test -args -num 3
```

## 参考
- [testing チートシート](https://qiita.com/nirasan/items/b357f0ad9172ab9fa19b)