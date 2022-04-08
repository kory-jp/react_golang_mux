# Session API 詳細

## API 一覧

| No. | API 機能 No. | 種類 | API 名 | 機能概要     |
| --: | -----------: | :--- | :----- | :----------- |
|   0 |     SESS-000 | API  | Login  | ログイン     |
|   1 |     SESS-001 | API  | Auth   | ログイン確認 |
|   2 |     SESS-002 | API  | Logout | ログアウト   |

## SESS-000

| API 機能 No. | SESS-000   |
| :----------- | :--------- |
| API 名       | Login      |
| 概要         | ログイン   |
| URL          | /api/login |

<br>

### 入力

| アクセス URL | /api/login |
| :----------- | :--------- |

### リクエストヘッダーその他

|  フィルード名   |       内容       |
| :-------------: | :--------------: |
|     Accept      | application/json |
|  Content-Type   | application/json |
| withCredentials |       true       |

#### POST データ

| JSON Key |     型 | 最大サイズ | 必須 | 暗号化 | 検索条件 |
| :------- | -----: | ---------: | :--: | :----: | :------- |
| email    | 文字列 |            |  ○   |        | 完全一致 |
| password | 文字列 |    20 文字 |  ○   |        | 完全一致 |

<br>

### 出力

#### 返却データ(JSON)

| JSON Key              |   型   | 最大サイズ | 必須 |      値の説明      |
| :-------------------- | :----: | ---------: | :--: | :----------------: |
| status                |  数値  |            |  ○   | 処理結果ステータス |
| message               | 文字列 |            |  ○   |     メッセージ     |
| user                  |  配列  |            |  ○   |                    |
| &emsp; user_id        |  数値  |            |      |         id         |
| &emsp; user_name      | 文字列 |         20 |      |        氏名        |
| &emsp; user_email     | 文字列 |            |      |   メールアドレス   |
| &emsp; user_password  | 文字列 |         20 |      |     パスワード     |
| &emsp; user_createdAt |  日付  |            |      |      作成日時      |

<br>
<br>

## SESS-001

| API 機能 No. | SESS-001          |
| :----------- | :---------------- |
| API 名       | Auth              |
| 概要         | ログイン確認      |
| URL          | /api/authenticate |

### 入力

| アクセス URL | /api/authenticate |
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

### 出力

#### 返却データ(JSON)

| JSON Key              |   型   | 最大サイズ | 必須 |      値の説明      |
| :-------------------- | :----: | ---------: | :--: | :----------------: |
| status                |  数値  |            |  ○   | 処理結果ステータス |
| message               | 文字列 |            |  ○   |     メッセージ     |
| user                  |  配列  |            |  ○   |                    |
| &emsp; user_id        |  数値  |            |      |         id         |
| &emsp; user_name      | 文字列 |         20 |      |        氏名        |
| &emsp; user_email     | 文字列 |            |      |   メールアドレス   |
| &emsp; user_password  | 文字列 |         20 |      |     パスワード     |
| &emsp; user_createdAt |  日付  |            |      |      作成日時      |

<br>
<br>

## SESS-002

| API 機能 No. | SESS-002    |
| :----------- | :---------- |
| API 名       | Logout      |
| 概要         | ログアウト  |
| URL          | /api/logout |

### 入力

| アクセス URL | /api/logout |
| :----------- | :---------- |

### リクエストヘッダーその他

|  フィルード名   |       内容       |
| :-------------: | :--------------: |
|     Accept      | application/json |
|  Content-Type   | application/json |
| withCredentials |       true       |

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

## 処理結果ステータス

| ステータス | 共通 | メッセージ内容                                 |
| ---------: | :--: | :--------------------------------------------- |
|        200 | 共通 | 通信が成功してデータを取得                     |
|        400 | 共通 | 必須項目が入力されておらずデータを取得できない |
|        401 | 共通 | ログインにかかる認証失敗                       |
