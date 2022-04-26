# User API 詳細

## API 一覧

| No. |  API 機能 No. | 種類 | API 名 | 機能概要             |
| --: | ------------: | :--- | :----- | :------------------- |
|   0 |  TASKCARD-000 | API  | Create | タスクカード新規作成 |
|   1 |  TASKCARD-001 | API  | Index  | タスクカード一覧取得 |
|   2 | TASKCARD-0002 | API  | Show   | タスクカード詳細取得 |
|   3 | TASKCARD-0003 | API  | Update | タスクカード更新     |

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
curl -XPOST -b cookie.txt -b 'cookie-name=' -d '{"todoId":"1","title":"test_title","purpose":"test_purpose","content":"test_content","memo":"test_memo"}' -H 'Content-Type: application/json' -H 'Accept: application/json'  http://localhost:8000/api/taskcard/new
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

## USER-001

| API 機能 No. | TASKCARD-001           |
| :----------- | :--------------------- |
| API 名       | Index                  |
| 概要         | タスクカード一覧取得   |
| URL          | /api/todo/:id/taskcard |

<br>

### 入力

| アクセス URL | /api/todo/:id/taskcard |
| :----------- | :--------------------- |

### リクエストヘッダーその他

|  フィルード名   |       内容       |
| :-------------: | :--------------: |
|     Accept      | application/json |
|  Content-Type   | application/json |
| withCredentials |       true       |

#### POST データ

無し

<br>

curl コマンド

```
curl -XGET -b cookie.txt -b 'cookie-name='  -H 'Content-Type: application/json' -H 'Accept: application/json'  "http://localhost:8000/api/todo/1/taskcard?page=1"
```

<br>

### 出力

#### 返却データ(JSON)

| JSON Key                    |   型   | 最大サイズ | 必須 |           値の説明           |
| :-------------------------- | :----: | ---------: | :--: | :--------------------------: |
| status                      |  数値  |            |  ○   |      処理結果ステータス      |
| message                     | 文字列 |            |  ○   |          メッセージ          |
| sumPage                     |  数値  |            |  ○   | ページネーションの総ページ数 |
| taskCards                   |  配列  |            |  ○   |                              |
| &emsp; taskCards_id         |  数値  |            |      |              id              |
| &emsp; taskCards_userId     |  数値  |            |      |           user_id            |
| &emsp; taskCards_postId     |  数値  |            |      |           post_id            |
| &emsp; taskCards_title      | 文字列 |            |      |           タイトル           |
| &emsp; taskCards_purpose    | 文字列 |            |      |             目的             |
| &emsp; taskCards_content    | 文字列 |            |      |             内容             |
| &emsp; taskCards_memo       | 文字列 |            |      |             メモ             |
| &emsp; taskCards_isFinished | 真偽値 |            |      |        完了未完了状態        |
| &emsp; taskCards_createdAt  |  日付  |            |      |             日付             |

<br>
<br>

## USER-002

| API 機能 No. | TASKCARD-002         |
| :----------- | :------------------- |
| API 名       | Show                 |
| 概要         | タスクカード詳細取得 |
| URL          | /api/taskcard/:id    |

<br>

### 入力

| アクセス URL | /api/taskcard/:id |
| :----------- | :---------------- |

### リクエストヘッダーその他

|  フィルード名   |       内容       |
| :-------------: | :--------------: |
|     Accept      | application/json |
|  Content-Type   | application/json |
| withCredentials |       true       |

#### POST データ

無し

<br>

curl コマンド

```
curl -XGET -b cookie.txt -b 'cookie-name='  -H 'Content-Type: application/json' -H 'Accept: application/json'  "http://localhost:8000/api/taskcard/1"
```

<br>

### 出力

#### 返却データ(JSON)

| JSON Key                   |      型      | 最大サイズ | 必須 |      値の説明      |
| :------------------------- | :----------: | ---------: | :--: | :----------------: |
| status                     |     数値     |            |  ○   | 処理結果ステータス |
| message                    |    文字列    |            |  ○   |     メッセージ     |
| taskCard                   | オブジェクト |            |  ○   |                    |
| &emsp; taskCard_id         |     数値     |            |      |         id         |
| &emsp; taskCard_userId     |     数値     |            |      |      user_id       |
| &emsp; taskCard_postId     |     数値     |            |      |      post_id       |
| &emsp; taskCard_title      |    文字列    |            |      |      タイトル      |
| &emsp; taskCard_purpose    |    文字列    |            |      |        目的        |
| &emsp; taskCard_content    |    文字列    |            |      |        内容        |
| &emsp; taskCard_memo       |    文字列    |            |      |        メモ        |
| &emsp; taskCard_isFinished |    真偽値    |            |      |   完了未完了状態   |
| &emsp; taskCard_createdAt  |     日付     |            |      |        日付        |

<br>
<br>

## USER-003

| API 機能 No. | TASKCARD-003      |
| :----------- | :---------------- |
| API 名       | Update            |
| 概要         | タスクカード更新  |
| URL          | /api/taskcard/:id |

<br>

### 入力

| アクセス URL | /api/taskcard/:id |
| :----------- | :---------------- |

### リクエストヘッダーその他

|  フィルード名   |       内容       |
| :-------------: | :--------------: |
|     Accept      | application/json |
|  Content-Type   | application/json |
| withCredentials |       true       |

#### POST データ

| JSON Key |     型 | 最大サイズ | 必須 | 暗号化 | 検索条件 |
| :------- | -----: | ---------: | :--: | :----: | :------- |
| todo_id  | 数値型 |            |  ○   |        |          |
| title    | 文字列 |    50 文字 |  ○   |        |          |
| purpose  | 文字列 |  1999 文字 |      |        |          |
| content  | 文字列 |  1999 文字 |      |        |          |
| memo     | 文字列 |  1999 文字 |      |        |          |

<br>

curl コマンド

```
curl -XPOST -b cookie.txt -b 'cookie-name=' -d '{"todoId":1,"title":"update_test_title","purpose":"update_test_purpose","content":"update_test_content","memo":"update_test_memo"}' -H 'Content-Type: application/json' -H 'Accept: application/json'  http://localhost:8000/api/taskcard/1
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
