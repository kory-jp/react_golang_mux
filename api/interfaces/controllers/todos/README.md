# Todo API 詳細

## API 一覧

| No. | API 機能 No. | 種類 | API 名      | 機能概要              |
| --: | -----------: | :--- | :---------- | :-------------------- |
|   0 |     TODO-000 | API  | Create      | Todo 新規作成         |
|   1 |     TODO-001 | API  | Index       | 一覧表示              |
|   2 |     TODO-002 | API  | Show        | 詳細表示              |
|   3 |     TODO-003 | API  | Update      | 更新                  |
|   4 |     TODO-004 | API  | IsFinished  | 完了状態更新          |
|   5 |     TODO-005 | API  | Delete      | 削除                  |
|   6 |     TODO-006 | API  | DeleteIndex | 削除後一覧取得        |
|   7 |     TODO-007 | API  | TagSearch   | タグによる Todos 取得 |

## TODO-000

| API 機能 No. | TODO-000      |
| :----------- | :------------ |
| API 名       | Create        |
| 概要         | Todo 新規作成 |
| URL          | /api/new      |

<br>

### 入力

| アクセス URL | /api/new |
| :----------- | :------- |

### リクエストヘッダー　その他

|  フィルード名   |        内容         |
| :-------------: | :-----------------: |
|     Accept      |  application/json   |
|  Content-Type   | multipart/form-data |
| withCredentials |        true         |

#### POST データ

| JSON Key |       型 | 最大サイズ | 必須 | 暗号化 | 検索条件 |
| :------- | -------: | ---------: | :--: | :----: | :------- |
| title    |   文字列 |    50 文字 |  ○   |        |          |
| content  |   文字列 |  2000 文字 |  ○   |        |          |
| image    | ファイル |            |      |        |          |
| tagIds   |     配列 |            |      |        |          |

<br>

curl コマンド

```
curl -XPOST -b cookie.txt -b 'cookie-name='  -F "title=curlTitle" -F "content=curlContent"  -F "tagIds=[1,2,3]" -H 'Content-Type: multipart/form-data' -H 'Accept: application/json'  http://localhost:8000/api/new
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

## TODO-002

| API 機能 No. | TODO-002   |
| :----------- | :--------- |
| API 名       | Index      |
| 概要         | 一覧表示   |
| URL          | /api/todos |

### 入力

| アクセス URL | /api/todos |
| :----------- | :--------- |

### リクエストヘッダー　その他

|  フィルード名   |        内容         |
| :-------------: | :-----------------: |
|     Accept      |  application/json   |
|  Content-Type   | multipart/form-data |
| withCredentials |        true         |

#### POST データ

無し

<br>

curl コマンド

```
curl -XGET -b cookie.txt -b 'cookie-name='  -H 'Content-Type: multipart/form-data' -H 'Accept: application/json'  "http://localhost:8000/api/todos?page=1"
```

<br>

### 出力

#### 返却データ(JSON)

| JSON Key               |   型   | 最大サイズ | 必須 |           値の説明           |
| :--------------------- | :----: | ---------: | :--: | :--------------------------: |
| status                 |  数値  |            |  ○   |      処理結果ステータス      |
| message                | 文字列 |            |  ○   |          メッセージ          |
| sumPage                |  数値  |            |  ○   | ページネーションの総ページ数 |
| todos                  |  配列  |            |      |                              |
| &emsp; todo_id         |  数値  |            |      |              id              |
| &emsp; todo_title      | 文字列 |    20 文字 |      |           タイトル           |
| &emsp; todo_content    | 文字列 |  2000 文字 |      |             内容             |
| &emsp; todo_imagePath  | 文字列 |            |      |       画像配信先の URL       |
| &emsp; todo_isFinished | 真偽値 |            |      |       Todo の完了状態        |
| &emsp; todo_created_at |  日付  |            |      |           作成日時           |
| &emsp; todo_tags       |  配列  |            |      |           タグ情報           |

<br>
<br>

## TODO-002

| API 機能 No. | TODO-002       |
| :----------- | :------------- |
| API 名       | Show           |
| 概要         | 詳細表示       |
| URL          | /api/todos/:id |

### 入力

| アクセス URL | /api/todos/:id |
| :----------- | :------------- |

### リクエストヘッダー　その他

|  フィルード名   |        内容         |
| :-------------: | :-----------------: |
|     Accept      |  application/json   |
|  Content-Type   | multipart/form-data |
| withCredentials |        true         |

#### POST データ

無し

<br>

curl コマンド

```
curl -XGET -b cookie.txt -b 'cookie-name='  -H 'Content-Type: multipart/form-data' -H 'Accept: application/json'  http://localhost:8000/api/todos/1
```

<br>

### 出力

#### 返却データ(JSON)

| JSON Key               |      型      | 最大サイズ | 必須 |      値の説明      |
| :--------------------- | :----------: | ---------: | :--: | :----------------: |
| status                 |     数値     |            |  ○   | 処理結果ステータス |
| message                |    文字列    |            |  ○   |     メッセージ     |
| todo                   | オブジェクト |            |      |                    |
| &emsp; todo_id         |     数値     |            |      |         id         |
| &emsp; todo_title      |    文字列    |    20 文字 |      |      タイトル      |
| &emsp; todo_content    |    文字列    |  2000 文字 |      |        内容        |
| &emsp; todo_imagePath  |    文字列    |            |      |  画像配信先の URL  |
| &emsp; todo_isFinished |    真偽値    |            |      |  Todo の完了状態   |
| &emsp; todo_created_at |     日付     |            |      |      作成日時      |
| &emsp; todo_tags       |     配列     |            |      |      タグ情報      |

<br>
<br>

## TODO-003

| API 機能 No. | TODO-003              |
| :----------- | :-------------------- |
| API 名       | Update                |
| 概要         | Todo 更新             |
| URL          | /api/todos/update/:id |

<br>

### 入力

| アクセス URL | /api/todos/update/:id |
| :----------- | :-------------------- |

### リクエストヘッダー　その他

|  フィルード名   |        内容         |
| :-------------: | :-----------------: |
|     Accept      |  application/json   |
|  Content-Type   | multipart/form-data |
| withCredentials |        true         |

#### POST データ

| JSON Key |       型 | 最大サイズ | 必須 | 暗号化 | 検索条件 |
| :------- | -------: | ---------: | :--: | :----: | :------- |
| title    |   文字列 |    50 文字 |  ○   |        |          |
| content  |   文字列 |  2000 文字 |      |        |          |
| image    | ファイル |            |      |        |          |
| tagIds   |     配列 |            |      |        |          |

<br>

curl コマンド

```
curl -XPOST -b cookie.txt -b 'cookie-name='  -F "title=curlTitleUp" -F "content=curlContentUp"  -F "tagIds=[3,4]" -H 'Content-Type: multipart/form-data' -H 'Accept: application/json'  http://localhost:8000/api/todos/update/1
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

## TODO-004

| API 機能 No. | TODO-004                  |
| :----------- | :------------------------ |
| API 名       | IsFinished                |
| 概要         | Todo 完了状態更新         |
| URL          | /api/todos/isfinished/:id |

<br>

### 入力

| アクセス URL | /api/todos/isfinished/:id |
| :----------- | :------------------------ |

### リクエストヘッダー　その他

|  フィルード名   |        内容         |
| :-------------: | :-----------------: |
|     Accept      |  application/json   |
|  Content-Type   | multipart/form-data |
| withCredentials |        true         |

#### POST データ

| JSON Key   |     型 | 最大サイズ | 必須 | 暗号化 | 検索条件 |
| :--------- | -----: | ---------: | :--: | :----: | :------- |
| isFinished | 真偽値 |            |  ○   |        |          |

<br>

### 出力

#### 返却データ(JSON)

| JSON Key |   型   | 最大サイズ | 必須 |      値の説明      |
| :------- | :----: | ---------: | :--: | :----------------: |
| status   |  数値  |            |  ○   | 処理結果ステータス |
| message  | 文字列 |            |  ○   |     メッセージ     |

<br>
<br>

## TODO-005

| API 機能 No. | TODO-005              |
| :----------- | :-------------------- |
| API 名       | Delete                |
| 概要         | Todo 削除             |
| URL          | /api/todos/delete/:id |

<br>

### 入力

| アクセス URL | /api/todos/delete/:id |
| :----------- | :-------------------- |

### リクエストヘッダー　その他

|  フィルード名   |        内容         |
| :-------------: | :-----------------: |
|     Accept      |  application/json   |
|  Content-Type   | multipart/form-data |
| withCredentials |        true         |

#### POST データ

無し

<br>

### 出力

#### 返却データ(JSON)

| JSON Key |   型   | 最大サイズ | 必須 |      値の説明      |
| :------- | :----: | ---------: | :--: | :----------------: |
| status   |  数値  |            |  ○   | 処理結果ステータス |
| message  | 文字列 |            |  ○   |     メッセージ     |

<br>
<br>

## TODO-006

| API 機能 No. | TODO-006                     |
| :----------- | :--------------------------- |
| API 名       | DeleteIndex                  |
| 概要         | Todo 削除後一覧取得          |
| URL          | /api/todos/deleteinindex/:id |

<br>

### 入力

| アクセス URL | /api/todos/deleteinindex/:id |
| :----------- | :--------------------------- |

### リクエストヘッダー　その他

|  フィルード名   |        内容         |
| :-------------: | :-----------------: |
|     Accept      |  application/json   |
|  Content-Type   | multipart/form-data |
| withCredentials |        true         |

#### POST データ

無し

<br>

### 出力

#### 返却データ(JSON)

| JSON Key               |   型   | 最大サイズ | 必須 |           値の説明           |
| :--------------------- | :----: | ---------: | :--: | :--------------------------: |
| status                 |  数値  |            |  ○   |      処理結果ステータス      |
| message                | 文字列 |            |  ○   |          メッセージ          |
| sumPage                |  数値  |            |  ○   | ページネーションの総ページ数 |
| todos                  |  配列  |            |      |                              |
| &emsp; todo_id         |  数値  |            |      |              id              |
| &emsp; todo_title      | 文字列 |    20 文字 |      |           タイトル           |
| &emsp; todo_content    | 文字列 |  2000 文字 |      |             内容             |
| &emsp; todo_imagePath  | 文字列 |            |      |       画像配信先の URL       |
| &emsp; todo_isFinished | 真偽値 |            |      |       Todo の完了状態        |
| &emsp; todo_created_at |  日付  |            |      |           作成日時           |
| &emsp; todo_tags       |  配列  |            |      |           タグ情報           |

<br>
<br>

## TODO-007

| API 機能 No. | TODO-007              |
| :----------- | :-------------------- |
| API 名       | TagSearch             |
| 概要         | タグによる Todos 取得 |
| URL          | /api/todos/tag/:id    |

<br>

### 入力

| アクセス URL | /api/todos/tag/:id |
| :----------- | :----------------- |

### リクエストヘッダー　その他

|  フィルード名   |        内容         |
| :-------------: | :-----------------: |
|     Accept      |  application/json   |
|  Content-Type   | multipart/form-data |
| withCredentials |        true         |

#### POST データ

無し

<br>

curl コマンド

```
curl -XGET -b cookie.txt -b 'cookie-name='  -H 'Content-Type: multipart/form-data' -H 'Accept: application/json'  "http://localhost:8000/api/todos/tag/1?page=1"
```

<br>

### 出力

#### 返却データ(JSON)

| JSON Key               |   型   | 最大サイズ | 必須 |           値の説明           |
| :--------------------- | :----: | ---------: | :--: | :--------------------------: |
| status                 |  数値  |            |  ○   |      処理結果ステータス      |
| message                | 文字列 |            |  ○   |          メッセージ          |
| sumPage                |  数値  |            |  ○   | ページネーションの総ページ数 |
| todos                  |  配列  |            |      |                              |
| &emsp; todo_id         |  数値  |            |      |              id              |
| &emsp; todo_title      | 文字列 |    20 文字 |      |           タイトル           |
| &emsp; todo_content    | 文字列 |  2000 文字 |      |             内容             |
| &emsp; todo_imagePath  | 文字列 |            |      |       画像配信先の URL       |
| &emsp; todo_isFinished | 真偽値 |            |      |       Todo の完了状態        |
| &emsp; todo_created_at |  日付  |            |      |           作成日時           |
| &emsp; todo_tags       |  配列  |            |      |           タグ情報           |

<br>
<br>

## 処理結果ステータス

| ステータス | 共通 | メッセージ内容                                 |
| ---------: | :--: | :--------------------------------------------- |
|        200 | 共通 | 通信が成功してデータを取得                     |
|        400 | 共通 | 必須項目が入力されておらずデータを取得できない |
|        401 | 共通 | ログインにかかる認証失敗                       |
