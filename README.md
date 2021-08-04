
![空き家でGo-GIF](https://user-images.githubusercontent.com/43948442/122051689-d81cb480-ce1f-11eb-958d-e9a92fa58cd7.gif)
## サイト概要

### [空き家でGo](http://akiyadego.com)<br>

身の回りにある空き家の写真をシェアする空き家情報サイトです。<br>
自分が見つけた空き家を共有することで、各都道府県の空き家情報が可視化されます。<br>
また、自分なりの活用方法も提案することで空き家の新しい価値が想像されます。<br>

#### テストユーザー
メールアドレス： test@example.com パスワード：test

### 開発した背景

2033年には空き家数2,150万戸、全住宅の3戸に1戸が空き家となる現状があります。<br>
ただ、そんな事実を知っている人は日本に何名いらっしゃるでしょうか。<br>
実際、世間では空き家の問題よりも「持ち家か？賃貸か？」というテーマがよくメディアでも取り上げられています。<br>
しかし、私はリモートワークの加速により、コロナ禍以前よりも空き家の価値が見直され、需要が高騰することで、空き家の新しい活用方法が提案される機会が訪れようとしているのではないか、と考えています。<br>
そこで私は、__空き家の新たな活用方法の提案の場__ 並びに __空き家の見える化__ を促進したいという思いからこちらのサービスを開発しました。

### ターゲットユーザ

- 空き家を所有している持ち主の方
- 持ち主ではないが空き家が身の回りにある方
- 空き家を探している方
- 空き家の活用方法を提案したい方

### 想定されるユースケース
1. 空き家を発見します。
1. 空き家の写真を１枚撮影します。
1. 写真、タイトル、都道府県などを選択し、空き家の情報をシェアします。

## 使用技術
#### バックエンド
- Go 1.16
#### フロントエンド
- Bootstrap 4.5
- HTML/CSS
#### インフラ
- AWS (VPC,EC2,RDS,Route53,IAM)
#### CI/CD
- Github Actions
#### ミドルウェア
- Nginx 1.19
#### データベース
- MySQL 5.7
#### その他
- Docker（docker-compose）
- Git/Github
- Makefile

## 機能一覧
- ユーザー登録
- 登録したユーザでログインすることができる
  - ログイン後、空き家情報を投稿することすることができる
  - ログイン後、投稿した空き家情報を更新することができる
  - ログイン後、投稿した空き家情報を削除することができる

## インフラ構成図
![インフラ構成図8_4](https://user-images.githubusercontent.com/43948442/128181323-8396ba0b-83b8-4226-af65-7d809fd0c37e.jpeg)

## 今後の改良計画
- ユーザーページの作成
- 活用方法提案のためのコメント機能
- 画像を複数枚投稿できるようにする
- 投稿の詳細ページの作成
- Google Maps APIの導入
- Vue.jsによるSPA化

## 作成者情報
白川 生成（Email: kinari.shi21@gmail.com ）