# [AkiyaDeGo](http://akiyadego.com)

![akiyadego com](https://user-images.githubusercontent.com/43948442/121775846-d3a88f80-cbc4-11eb-9002-47183abe762b.png)

## サイト概要

### [空き家でGo](http://akiyadego.com)<br>

身の回りにある空き家の写真をシェアする空き家情報サイトです。<br>
自分が見つけた空き家を共有することで、各都道府県の空き家情報が可視化されます。<br>
また、自分なりの活用方法も提案することで空き家の新しい価値が想像されます。<br>

### テーマを選んだ理由

2033年には空き家数2,150万戸、なんと全住宅の3戸に1戸が空き家となる現状があります。<br>
私はその事実を、大学生の時に初めて知りました。そして、それは自分がいた建築学科でも<br>
よく議論の的になっていました。<br>
しかしいざ大学の外に出てみると、世間では空き家の問題には関心がなく、「持ち家か？賃貸か？」<br>
というテーマがよくメディアで取り上げられていました。<br>
私は、__「目に見えてる範囲の問題に少しでも興味を持ってもらいたい」__ という思いからこのアプリを<br>
開発しました。<br>

### ターゲットユーザ

- 空き家を所有している持ち主の方
- 持ち主ではないが空き家が身の回りにある方
- 空き家を探している方
- 空き家の活用方法を提案したい方

### 使い方
1. 空き家を発見します。
1. 写真を撮影します。
1. 写真、タイトル、都道府県などを選択し空き家の情報をシェアします。

## 使用技術
#### バックエンド
- Go 1.16.2
#### フロントエンド
- Bootstrap 4.5
- HTML/CSS
#### データベース
- MySQL 5.7.34
#### ミドルウェア
- Nginx 1.19.10
#### 本番環境
- AWS (VPC,EC2,RDS,Route53,IAM)

## 機能一覧
1. ユーザー登録、ログイン機能
1. 管理者機能（投稿編集、投稿削除）
1. 投稿機能（画像投稿）

## インフラ構成図
![インフラ構成図](https://user-images.githubusercontent.com/43948442/121861404-bf83a000-cd34-11eb-9399-e268fbf40554.jpeg)