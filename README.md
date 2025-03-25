# ✨ AI チャットボット サーバー 🤖

ヤッホー！👋 このプロジェクトは超クールな AI チャットボットなの！💁‍♀️
LINE で会話できちゃうし、賢い AI が返事してくれちゃうの！✨

## 🌟 できること

- LINE でおしゃべりできる～！💬
- OpenRouter API で賢くなってる！🧠
- Render でかんたんデプロイ！🚀

## 🎯 必要なもの

- Go 言語 v1.21 以上 💪
- LINE Developers アカウント 📱
- LINE Messaging API チャネル 💌
- OpenRouter API キー 🔑
- Render アカウント ☁️

## 🎨 セットアップの仕方

### 1️⃣ まずはクローンしちゃおう！
```bash
git clone [repository-url]
cd family_ai_bot
```

### 2️⃣ 依存関係インストールしなきゃ！
```bash
go mod download
```

## 💫 環境変数の設定方法

### 👩‍💻 ローカル開発するなら
以下の環境変数を設定してね！

```bash
# Windows PowerShell 使ってる子向け 🪟
$env:PORT = "8080"
$env:OPENROUTER_API_KEY = "your_openrouter_api_key_here"
$env:LINE_CHANNEL_SECRET = "your_line_channel_secret_here"
$env:LINE_CHANNEL_TOKEN = "your_line_channel_token_here"
$env:APP_URL = "http://localhost:8080"

# Linux/Mac 使ってる子向け 🍎
export PORT=8080
export OPENROUTER_API_KEY=your_openrouter_api_key_here
export LINE_CHANNEL_SECRET=your_line_channel_secret_here
export LINE_CHANNEL_TOKEN=your_line_channel_token_here
export APP_URL=http://localhost:8080
```

### ✨ Render にデプロイする場合
Render のダッシュボードで以下の環境変数を設定してね！

| 環境変数 | 説明 | 設定値 |
|----------|------|--------|
| `OPENROUTER_API_KEY` | OpenRouter の API キー 🔑 | OpenRouter でゲットしてね！ |
| `LINE_CHANNEL_SECRET` | LINE のシークレット 🤫 | LINE Developers でゲット！ |
| `LINE_CHANNEL_TOKEN` | LINE のトークン 🎟️ | LINE Developers でゲット！ |
| `APP_URL` | アプリの URL 🌐 | Render が自動で設定してくれる！ |
| `PORT` | サーバーのポート番号 🔌 | Render が自動で設定してくれる！ |

## 📱 LINE Messaging API の設定方法

1. [LINE Developers Console](https://developers.line.biz/console/) でアカウント作っちゃお！ 📝
2. 新しい Messaging API チャネルを作るよ！ ✨
3. シークレットとトークンをゲットしよう！ 🎁
4. Webhook URL を設定！（例: https://your-app.onrender.com/callback）🔗
5. Webhook 送信を ON にしちゃお！ ✅

## 🚀 Render へのデプロイ方法

1. Render でアカウント作って、新しい Web サービスを作るよ！ 💫
2. GitHub リポジトリと連携！ 🤝
3. 上の環境変数を設定しちゃお！ ⚙️
4. デプロイボタンをポチッと押す！ 🎯
5. デプロイ完了したら、LINE Developers で Webhook URL を更新！ 🔄
   - Webhook URL: `https://[あなたのアプリ名].onrender.com/callback`

## 👩‍💻 開発者向け情報

- `main.go`: メインのプログラムだよ！ 📝
- `render.yaml`: Render の設定ファイル！ ⚙️
- デフォルト AI モデル: mistralai/mistral-7b-instruct 🤖

## ⚠️ 注意事項

- 環境変数は Render でちゃんと設定してね！ 🔐
- LINE Messaging API の制限に気をつけて！ ⚡
- OpenRouter API の制限にも注意だよ！ 🌈
- Webhook は HTTPS 必須だけど Render が自動でやってくれる！ 🔒

困ったことがあったら、いつでも聞いてね！ 💕 