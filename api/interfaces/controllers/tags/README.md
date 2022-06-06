# Todo API 詳細

## API 一覧

| No. | API 機能 No. | 種類 | API 名 | 機能概要      |
| --: | -----------: | :--- | :----- | :------------ |
|   0 |      TAG-000 | API  | Index  | Tags 一覧取得 |

## TAG-000

| API 機能 No. | TAG-000  |
| :----------- | :------- |
| API 名       | Index    |
| 概要         | 一覧表示 |
| URL          | /api/tag |

### 入力

| アクセス URL | /api/tag |
| :----------- | :------- |

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
curl -XGET -b cookie.txt -b 'cookie-name='  -H 'Content-Type: application/json' -H 'Accept: application/json'  http://localhost:8000/api/tag
```

<br>

### 出力

#### 返却データ(JSON)

| JSON Key         |   型   | 最大サイズ | 必須 |      値の説明      |
| :--------------- | :----: | ---------: | :--: | :----------------: |
| status           |  数値  |            |  ○   | 処理結果ステータス |
| message          | 文字列 |            |  ○   |     メッセージ     |
| tags             |  配列  |            |      |                    |
| &emsp; tag_id    |  数値  |            |      |         id         |
| &emsp; tag_value | 文字列 |            |      |      バリュー      |
| &emsp; tag_label | 文字列 |            |      |       ラベル       |

<br>
<br>
