# What is This?
- プログラミング不要のAPIモック
- リクエスト内容をログファイルに出力します
- レスポンスを設定ファイルに書き込むだけでOK
- 今度ちゃんとかく

# Usage
- route.yamlにルーティングとレスポンスを定義
- 必要に応じてconfig.yamlを書き換える
- 起動してリクエストを送る
```
$ go run main.go
```
- リクエスト内容はlogファイルに書き出される
- レスポンスとしてroute.yamlで定義した内容が返ってくる

# License
MIT