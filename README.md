# AI Chatbot API Server

このプロジェクトは、Go言語で実装されたLINE Messaging APIを使用したAIチャットボットです。OpenRouter APIを使用してAI機能を提供し、Renderにデプロイすることを想定しています。

## 機能

- LINE Messaging APIとの連携
- OpenRouter APIを使用したAIチャット機能
- Renderへのデプロイ設定

## 必要条件

- Go 1.21以上
- LINE Developersアカウント
- LINE Messaging APIチャネル
- OpenRouter APIキー
- Renderアカウント（デプロイ用）

## セットアップ

1. リポジトリをクローン
```bash
git clone [repository-url]
cd family_ai_bot
```

2. 依存関係のインストール
```bash
go mod download
```

3. ローカル開発時の環境変数設定
```bash
# Windows PowerShell
$env:PORT = "8080"
$env:OPENROUTER_API_KEY = "your_openrouter_api_key_here"
$env:LINE_CHANNEL_SECRET = "your_line_channel_secret_here"
$env:LINE_CHANNEL_TOKEN = "your_line_channel_token_here"
$env:APP_URL = "your_app_url_here"

# Linux/Mac
export PORT=8080
export OPENROUTER_API_KEY=your_openrouter_api_key_here
export LINE_CHANNEL_SECRET=your_line_channel_secret_here
export LINE_CHANNEL_TOKEN=your_line_channel_token_here
export APP_URL=your_app_url_here
```

4. サーバーの起動
```bash
go run main.go
```

## LINE Messaging APIの設定

1. [LINE Developers Console](https://developers.line.biz/console/)でアカウントを作成
2. 新しいMessaging APIチャネルを作成
3. チャネルシークレットとチャネルアクセストークンを取得
4. Webhook URLを設定（例: https://your-app.onrender.com/callback）
5. Webhook送信を有効化

## Renderへのデプロイ

1. Renderでアカウントを作成し、新しいWebサービスを作成
2. GitHubリポジトリと連携
3. Renderダッシュボードで環境変数を設定:
   - `OPENROUTER_API_KEY`: OpenRouterのAPIキー
   - `LINE_CHANNEL_SECRET`: LINEチャネルシークレット
   - `LINE_CHANNEL_TOKEN`: LINEチャネルアクセストークン
   - `APP_URL`: RenderのアプリケーションURL（自動的に設定されます）
   - `PORT`: Renderが自動的に設定します
4. デプロイを実行
5. デプロイ完了後、LINE DevelopersコンソールでWebhook URLを更新

## 開発者向け情報

- `main.go`: メインアプリケーションファイル
- `render.yaml`: Renderデプロイ設定
- デフォルトAIモデル: mistralai/mistral-7b-instruct

## 注意事項

- 環境変数はRenderのダッシュボードで適切に設定してください
- LINE Messaging APIの利用制限に注意してください
- OpenRouter APIの利用制限とレート制限に注意してください
- Webhook URLはHTTPS必須です 