# President Todo

<p align="center">
  <img src="https://user-images.githubusercontent.com/66899822/172115426-365e3ba9-c3f6-430b-973f-8e846b5bdc0f.jpeg" height="400px"　> 
</p>

## 概要/開発経緯

`Golang` の基本的な知識を習得するために開発している Todo アプリになります。
基礎的なところから理解を図るために極力、フレームワークを利用せずに開発に努めております。
また合わせて、`React+TypeScript+Redux` の組み合わせでアプリを作成することも目標としており、こちらも最低限のパッケージで開発を進めております。

アプリのコンセプトとしては以下のサイトのタスク管理をアプリケーション化したらというコンセプトのもと機能や要件を考えて開発を進めております。
</br>
</br>
[参考サイト](https://president-ac.jp/blog/taskmanagement/)　
</br>
</br>
~~[作成したポートフォリオサイト](https://president-todo.net/)~~

~~平日の 9:30~20:00 の間で起動しており、それ以外の時間帯はサービスを落としております。~~

22/07/15 現在はサービスを終了しております

<br>
<br>

## こだわったところ

- `Client(React)`側において`Redux/TypeScript`において状態管理

`React`を用いた開発における必ず上がる課題として状態管理があるかと思います。現在、便利な`Hooks`なども出てきてはおりますが、現場ではやはりまだ使用率の高く知識を活かせる可能性が高い`Redux`を用いて状態管理をおこなっております。一方、`Redux+TypeScript`を用いる場合によく用いられる`typescript-fsa`や`Redux-Toolkit`は使用せずに型指定した状態(`state`)の受け渡しの流れをより意識できるのではないかと考え素の`TypeScript`で`Redux`の構築を図りました。

</br>

- API(golang)を`Clean Architecture`にて設計

以前に `Ruby on Rails` にて作成したポートフォリオは `Rails`の設計である `MVC`に準じてアプリケーションを作成しました。

[Rails のポートフォリオ](https://github.com/kory-jp/proto2)

今回は別の視点からアプリ開発の知識を得ることを目標に設計に`Clean Architecture`を採用しました。保守性などの観点から`MVC`以上に各担当役割を細かく分けている設計思想なので、コーディングの際は常にファイルの役割を考えつつ多少コードの記載量が増えても役割ごとにファイル分割を意識しました。

</br>

- データ整合性を維持するデータベース処理`Transaction`を、上位レイアのアプリケーション層で実行できるように`interface`を用いたダックタイピングにて実現

`Todo`をタグ分類するために、モデルとして`Todo`,`Tag`中間テーブルとして`Todo_Tag_Relations`を作成。`Todo`と`Todo_Tag_Relations`を保存する際に、どちらか一方が保存処理に失敗した場合、`Rollback`をするように設計しました。ポイントとして、`transaction`にまつわるデータベース操作は`Clean Architecture`において`interfaces/database`層が制御することになりますが、アプリにおけるデータの保存や廃棄の判断はより上位レイアの`usecase/interactor`(アプリケーション)層が関心をもつ必要があります。一方、上位レイア(アプリケーション層)は下位レイア(`interfaces/database`)に依存することは`Clean Architecture`においては設計思想に沿いません。そのため、DIP（依存関係逆転の原則）の考えのもと`interface`を用いて、いわゆるダックタイピングによりアプリケーション層にて`Transaction`操作を実現しました。

</br>

- API(golang)側の`Mock`を利用したテストコード

上述の通り、API 側の設計として、`Clean Architecture`を採用しており各種役割ごとに階層化されており、効率的にテストするために初めて`goMock`を採用しました。
とくにアプリケーションロジックの役割をになっている `api/usecase/interactor` をテストする際に、通信やデータベース処理おいて発生したエラーなどアプリケーションロジックに関係ないエラーを排除できるような設計にできました。

</br>

- インフラ管理を自動化

インフラとして`AWS`を利用しておりますが、構築は`Terraform`を利用しており、インフラ構築やアプリケーションの起動を`Terraform`コマンドで実行しております。
また、`GitHub Action`を利用しており、アプリケーションの変更、修正が`GitHub`にプッシュされたタイミングにテストが自動的に実行され、`main`ブランチにマージされたタイミングで`AWS`へ自動デプロイ。
開発の効率性、開発体験の向上を実現しました。

</br>
</br>

## 使用技術

バックエンド

- Golang(Go Modules)
- Nginx
- MySQL

フロントエンド

- React(create-react-app/Material-ui/Redux)
- TypeScript
- HTML/CSS

Golang 主要パッケージ

- go-playground/validator
- gorilla/mux
- gorilla/sessions
- joho/godotenv
- rs/cors
- ini.v1
- stretchr/testify
- .air(Golang のホットリロード)
- deleve(Golang におけるデバッグ)

インフラ関連

- docker(開発環境+本番環境構築)
- GitHub(バージョン管理)
- GitHub Actions (CI/CD)
- Terraform(IaC)
- AWS
  - VPC
  - Internet Gateway
  - ACM
  - ALB
  - Fargate
  - ECR

<br>
<br>

デザインツール

- Figma

\*事前に Figma によりサイトデザインを構築してから View にまつわるデザイン構築に取り掛かりました。また、ロゴ等に関しても Figma を用いて作成しました。

<img width="360" alt="DC98ADE5-C3F6-4A48-BF0B-D22C0C03A4C1" src="https://user-images.githubusercontent.com/66899822/172099177-48383f0e-1dd4-4327-923b-626d76a2b98c.png">

## 技術選定理由

### バックエンド: Golang

- 前回、`Ruby on Rails`でアプリを開発しており、次は静的型付け言語でフレームワーク無しの開発を考えていた
  [Rails のポートフォリオ](https://github.com/kory-jp/proto2)
- コードがシンプルで静的型付け言語においては学習難易度が低い
- アプリを作成するためのライブラリが豊富
- 近年の言語としての勢い

### フロントエンド: React/TypeScript

- コンポーネント分割に重きをおくライブラリで、ある程度の分割が出来ていれば可読性が上がるのと同時にメンテナンス性が上がるため
- `TypeScript`を併用することで型によるさらなる安全性の向上
- `Vue` との比較で、より深く `JavaScript` の理解が必要な `React` にキャリア初期から触れることで自己成長につながるのではないかと考えたため

### インフラ

#### Docker/Fargate(AWS)

- 開発環境と本番環境での環境差異の解消
- ローカルでも仮想本番環境を実現できてデプロイ作業の負担軽減
- 近年、コンテナデプロイが増加している傾向だったので、知見を得るために挑戦

#### Terraform(IaC)

- インフラをコード化することで、インフラ理解がより深まりまた学習記録として残せる
- 今回の Fargate/RDS と料金が高額なリソースを仕様しており、即時に立ち上げ/破壊が可能なので節約につながる

#### GitHubAction(CI/CD)

- CI において以前作成したポートフォリオにおいては Circle CI を利用していたが、すべて GitHub 上で完結できる点に GitHubAction の魅力を感じた
- コンテナを用いたデプロイはデプロイ作業に時間がかかるために、GitHub Actions の CD 機能を用いて自動化を図り開発体験の向上を目指した

<br>
<br>

## アーキテクチャ設計

_api(golang)/Clean Architecture を採用_

![](https://blog.tai2.net/images/CleanArchitecture.jpg)

以下の理由で採用

- 再利用性の高い設計になり生産性が向上する
- コードの可読性が上がり、メンテナンスが容易になる
- 変化に強い設計になる

[参考サイト 2](https://qiita.com/hirotakan/items/698c1f5773a3cca6193e)

<br>

_client(react)/Atomic Design を採用_

![](https://cdn.codegrid.net/2017-atomic-design/img/hierarchy.png)

以下の理由で採用

- コードの再利用性が高いの効率的に開発が可能
- メンテナンスが容易になる

[参考サイト 3](https://www.codegrid.net/articles/2017-atomic-design-1/)

<br>
<br>

## インフラ構成図

![infra_ver2](https://user-images.githubusercontent.com/66899822/175237186-23c3ddd5-cbc0-4dc7-a603-4a0d24120b5c.png)

<br>
<br>

## 主要機能

User

- 新規登録
- ログイン
- ログアウト

\*session と cookie を利用することでにログイン状態の確認

Todo

- 新規投稿
- 一覧表示
- 詳細表示
- 編集
- 削除
- 完了/未完了切り替え
- タグ分類
- タグ検索
- 重要度で分類
- 緊急度で分類
- 重要度、緊急度で検索

TaskCard

- 新規投稿
- 一覧表示
- 詳細表示
- 編集
- 削除
- 完了/未完了切り替え

その他

画像投稿/保存/配信

Toast による通知機能

Todo が 5 以上投稿されるとページネーションにより、データの一部取得

コマンドを実行することでテストデータを自動作成

一部、ファイルに対してテストコードを作成

<br>
<br>

## モデルデザイン

![ER](https://user-images.githubusercontent.com/66899822/165665661-25880a7c-fa77-4c79-af0a-bf6999ac69f3.png)

<br>
<br>

## ローカルでの開発環境での起動方法

- クローン

```
git clone git@github.com:kory-jp/react_golang_mux.git
```

- ルートディレクトリへ移動

```
cd golang_react_mux
```

- 環境変数設定

React 側

```
touch client/client-app/.env
```

client/client-app/.env へキー(環境変数)を入力

```
REACT_APP_API_URL="http://localhost:8000/api/"
```

golang 側

```
touch api/env/development.env
```

api/env/development.env へキー(環境変数)を入力

```
API_PORT = 8000
LOG_FILE = apiapp.log
DRIVER = mysql
USER_NAME = root
PASSWORD = root
HOST = mysql_container
DB_PORT = 3306
DB_NAME = react_golang_mux
ALLOWED_ORIGINS = http://localhost:8080
ALLOWED_METHODS = "POST GET PUT DELETE"
ALLOWED_HEADERS = "Origin Accept Content-Type X-Requested-With Access-Control-Allow-Origin"
SESSION_KEY = sample-session-key
AWS_BUCKET  = development
AWS_ACCESS_KEY_ID = development
AWS_SECRET_ACCESS_KEY = development
AWS_REGION = sample-region
```

- docker 起動

```
docker-compose build
docker-compose up -d
```

- サイトアクセス

[Todo](http://localhost:8080/)

- コンテナログイン

```
docker exec -it go_container /bin/sh
docker exec -it react_container /bin/sh
docker exec -it mysql_container /bin/sh
```

<br>
<br>

## テストデータ投入

```
docker exec -it go_container /bin/sh
go run infrastructure/seeds/development/seeder.go
```

テストユーザーによるログイン

```
メールアドレス:
sam@exm.com
パスワード
password
```

<br>
<br>

## 各種ログ確認

golang のログを確認

```
docker logs go_container
```

MySQL のクエリを確認

`/mysql/lgos/mysqld.log`

<br>
<br>

## テストコード実行

- 実装したテストコード

```
api/interfaces/controllers/users/user_controller_test.go
api/interfaces/controllers/sessions/session_controller_test.go
api/interfaces/controllers/todos/todo_controller_test.go
api/usecase/user/user_interactor_test.go
api/usecase/session/session_interactor_test.go
api/usecase/todo/todo_interactor_test.go
```

- テストコード実行

```
cd api
sh test.sh
```

- 成功した場合

```
$sh test.sh
TEST START
TEST USECASE USERS OK
TEST USECASE TODOS OK
TEST USECASE SESSIONS OK
TEST USECASE TASK_CARDS OK
TEST CONTROLLER USERS OK
TEST CONTROLLER TODOS OK
TEST CONTROLLER SESSIONS OK
TEST CONTROLLER TASK_CARDS OK
TEST ALL COMPLETED
```

- 失敗した場合

エラーログの詳細が表示

<br>
<br>

## 各種 API 仕様詳細

以下のサイトを参考にデザイン

[［API］ API 仕様書の書き方](https://qiita.com/sunstripe2011/items/9230396febfab2eae2c2)

- [User API](https://github.com/kory-jp/react_golang_mux/tree/main/api/interfaces/controllers/users)
- [Session API](https://github.com/kory-jp/react_golang_mux/tree/main/api/interfaces/controllers/sessions)
- [Todo API](https://github.com/kory-jp/react_golang_mux/tree/main/api/interfaces/controllers/todos)
- [Tag API](https://github.com/kory-jp/react_golang_mux/tree/main/api/interfaces/controllers/tags)

<br>
<br>
