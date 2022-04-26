# User API 詳細

## API 一覧

| No. | API 機能 No. | 種類 | API 名 | 機能概要             |
| --: | -----------: | :--- | :----- | :------------------- |
|   0 | TASKCARD-000 | API  | Create | タスクカード新規作成 |

## USER-000

| API 機能 No. | TASKCARD-000         |
| :----------- | :------------------- |
| API 名       | Create               |
| 概要         | タスクカード新規作成 |
| URL          | /api/taskcard/new    |

<br>

### 入力

| アクセス URL | /api/taskcard/new |
| :----------- | :---------------- |

### リクエストヘッダーその他

|  フィルード名   |       内容       |
| :-------------: | :--------------: |
|     Accept      | application/json |
|  Content-Type   | application/json |
| withCredentials |       true       |

#### POST データ

| JSON Key   |       型 | 最大サイズ | 必須 | 暗号化 | 検索条件 |
| :--------- | -------: | ---------: | :--: | :----: | :------- |
| todo_id    |   数値型 |            |  ○   |        |          |
| title      |   文字列 |    50 文字 |  ○   |        |          |
| purpose    |   文字列 |  1999 文字 |      |        |          |
| content    |   文字列 |  1999 文字 |      |        |          |
| memo       |   文字列 |  1999 文字 |      |        |          |
| isFinished | 真偽値型 |            |      |        |          |

<br>

curl コマンド

```
curl -XPOST -b cookie.txt -b 'cookie-name=' -d '{"todo_id":"1","title":"test_title","purpose":"test_purpose","content":"test_content","memo":"test_memo"}' -H 'Content-Type: application/json' -H 'Accept: application/json'  http://localhost:8000/api/taskcard/new
```

<br>

### 出力

#### 返却データ(JSON)

| JSON Key |   型   | 最大サイズ | 必須 |      値の説明      |
| :------- | :----: | ---------: | :--: | :----------------: |
| status   |  数値  |            |  ○   | 処理結果ステータス |
| message  | 文字列 |            |  ○   |     メッセージ     |

<br>
<br>

## 処理結果ステータス

| ステータス | 共通 | メッセージ内容                                 |
| ---------: | :--: | :--------------------------------------------- |
|        200 | 共通 | 通信が成功してデータを取得                     |
|        400 | 共通 | 必須項目が入力されておらずデータを取得できない |
|        401 | 共通 | ログインにかかる認証失敗                       |
