# go-for-dynamodb-playground

こちらのリポジトリは備忘録的な golang で dynamoDB に対して CRUD 操作を行う為のサンプルコードです。備忘録としての意味が強いですが参考になるところがあれば幸いです。

---

## 前提

- AWS アカウントが必要です。
- Docker に関しての知識が多少必要です。(compose の up/down、volumes について)
- Docker をダウンロードし、Docker が起動できる状態
- Docker 環境化で使用する事を想定していますが適宜、必要な箇所を抜粋して利用可能であれば利用してください。

---

## 環境ファイルを準備する

- AWS の key 情報を読み込む際の key 情報は env ファイルに記述する
- 記載する情報は以下の通り

```
go-for-dynamodb-playground/.go_envを作成。以下を記述

region={リージョン情報}
aws_access_key_id={アクセスキー(AWSのIAMから取得)}
aws_secret_access_key={シークレットキー(AWSのIAMから取得)}
```

---

## Docker compose up で立ち上げる
