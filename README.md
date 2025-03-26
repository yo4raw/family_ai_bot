# ✨ AI チャットボット サーバー 🤖
[![Build Status](https://img.shields.io/badge/build-passing-brightgreen)](https://github.com/yo4raw/family_ai_bot/actions)
[![License](https://img.shields.io/badge/license-MIT-blue)](LICENSE)

## 📑 目次

- [はじめに](#はじめに)
- [できること](#できること)
- [必要なもの](#必要なもの)
- [セットアップの仕方](#セットアップの仕方)
- [環境変数の設定方法](#環境変数の設定方法)
- [LINE Messaging API の設定方法](#line-messaging-api-の設定方法)
- [Render へのデプロイ方法](#render-へのデプロイ方法)
- [開発者向け情報](#開発者向け情報)
- [注意事項](#注意事項)

## はじめに

こんにちは、みんな！👋  
この README では、私たちの超クールな AI チャットボット サーバーの魅力をたっぷりと紹介するよ！🤖💬  
LINE でのおしゃべり、賢い AI の応答、そして Render を使ったスムーズなデプロイなど、楽しい機能が盛りだくさん！✨

## 🌟 できること

- LINE 上で気軽に会話が楽しめるよ！💬
- 賢い AI がパパッとあなたの質問にお答え！🧠
- Render で簡単にデプロイできるから、どこでも利用できる！🚀

## 🎯 必要なもの

このプロジェクトを始めるために、以下のアイテムが必要だよ！👍

- Go 言語 (バージョン 1.21 以上) 💪
- LINE Developers のアカウント 📱
- LINE Messaging API のチャネル 💌
- OpenRouter API のキー 🔑
- Render のアカウント ☁️

## 🎨 セットアップの仕方

まずはリポジトリをクローンして、ローカル環境でプロジェクトを起動しちゃおう！👇

### 1️⃣ リポジトリをクローンしよう！
```bash
git clone [repository-url]
cd family_ai_bot
```

### 2️⃣ 依存関係をインストールしよう！
```bash
go mod download
```

## 💫 環境変数の設定方法

プロジェクトで必要な環境変数を設定しよう！ローカル開発用と Render デプロイ用、それぞれの方法をチェックしてね。👇

### 👩‍💻 ローカル開発の場合
以下の環境変数を設定してね！

```bash
# Windows PowerShell を使っている人へ 🪟
$env:PORT = "8080"
$env:OPENROUTER_API_KEY = "your_openrouter_api_key_here"
$env:LINE_CHANNEL_SECRET = "your_line_channel_secret_here"
$env:LINE_CHANNEL_TOKEN = "your_line_channel_token_here"
$env:APP_URL = "http://localhost:8080"

# Linux/Mac を使っている人へ 🍎
export PORT=8080
export OPENROUTER_API_KEY=your_openrouter_api_key_here
export LINE_CHANNEL_SECRET=your_line_channel_secret_here
export LINE_CHANNEL_TOKEN=your_line_channel_token_here
export APP_URL=http://localhost:8080
```

### ✨ Render にデプロイする場合
Render のダッシュボードから、下記の環境変数を設定しよう！🎉

| 環境変数              | 説明                                      | 設定値                                     |
|-----------------------|-------------------------------------------|-------------------------------------------|
| `OPENROUTER_API_KEY`  | OpenRouter の API キー 🔑                | OpenRouter で取得しよう！                 |
| `LINE_CHANNEL_SECRET` | LINE のシークレット 🤫                   | LINE Developers で取得！                   |
| `LINE_CHANNEL_TOKEN`  | LINE のトークン 🎟️                       | LINE Developers で取得！                   |
| `APP_URL`             | アプリの URL 🌐                          | Render が自動で設定してくれるよ！           |
| `PORT`                | サーバーのポート番号 🔌                   | Render が自動で設定してくれるよ！           |

## 📱 LINE Messaging API の設定方法

LINE Developers Console で下記の手順に従って、Messaging API の設定をしよう！🎉

1. [LINE Developers Console](https://developers.line.biz/console/) でアカウントを作成しよう！📝
2. 新しい Messaging API チャネルを作成しよう！✨
3. チャネルシークレットとチャネルトークンを取得しよう！🎁
4. Webhook URL を設定しよう！（例: `https://your-app.onrender.com/callback`）🔗
5. Webhook の送信を ON にしよう！✅

## 🚀 Render へのデプロイ方法

Render を使って、サーバーをどこからでも公開しよう！以下の手順で進めてね。👇

1. Render でアカウントを作成し、新しい Web サービスを作成する！💫
2. GitHub リポジトリと連携しよう！🤝
3. 上記の環境変数を設定しよう！⚙️
4. デプロイボタンを押して、デプロイを実行しよう！🎯
5. デプロイ完了後、LINE Developers で Webhook URL を更新しよう！🔄  
   - Webhook URL: `https://[あなたのアプリ名].onrender.com/callback`

## 👩‍💻 開発者向け情報

開発者のみなさんへ、便利な情報をお届けします！🛠️

- `main.go`: アプリのメインプログラム。自由にカスタマイズしてね！ 📝
- `render.yaml`: Render の設定ファイル。デプロイの参考にどうぞ！ ⚙️
- デフォルト AI モデル: `mistralai/mistral-7b-instruct` 🤖

## ⚠️ 注意事項

最後に、いくつか大事なポイントを確認してね！🔍

- Render 上で環境変数が正しく設定されているか確認しよう！🔐
- LINE Messaging API の利用制限に注意してね！⚡
- OpenRouter API の利用制限もチェックしておこう！🌈
- Webhook は HTTPS が必須だよ！（Render が自動で対応してくれるから安心！）🔒

もし何か困ったことがあったら、いつでも遠慮なく質問してね！💕