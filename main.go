package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/line/line-bot-sdk-go/v7/linebot"
)

// 🎀✨ チャットリクエストの構造体だよ！ ユーザーから送られてくるメッセージを受け取るための可愛い箱っこだよ！ 😄💖
type ChatRequest struct {
	Message string `json:"message"` // 💬 ここにメッセージが入るよ！
}

// 💌🌟 チャットレスポンスの構造体だよ！ AIからの素敵な返信を詰め込む魔法の箱だよ！ 🤩🎉
type ChatResponse struct {
	Response string `json:"response"` // 🤖 ここにAIの返信が入るよ！
}

// 🌟🚀 OpenRouter APIに送るリクエストの構造体だよ！ AIと楽しくおしゃべりするための大切な情報がぎゅっと詰まってるよ！ 😃✨
type OpenRouterRequest struct {
	Model    string        `json:"model"`    // 🎯 どのAIモデルを使うか指定するよ！
	Messages []ChatMessage `json:"messages"` // 💬 会話の履歴を保存するよ！
}

// 💭💬 チャットメッセージの構造体だよ！ 誰が話したか、そして何を話したか、心温まる会話の記録を残すよ！ 😊💖
type ChatMessage struct {
	Role    string `json:"role"`    // 👤 話した人の役割（AIかユーザーか）
	Content string `json:"content"` // 📝 実際のメッセージ内容
}

// 🎁✨ OpenRouter APIからのレスポンスの構造体だよ！ AIからのキラキラした返信がここに入るよ！ 🌈🤩
type OpenRouterResponse struct {
	Choices []struct {
		Message struct {
			Content string `json:"content"` // 💬 AIからの返信がここに入るよ！
		} `json:"message"`
	} `json:"choices"`
}

// 🌈🚀 OpenRouter APIを呼び出して、AIと楽しくおしゃべりする関数だよ！
// callOpenRouterAPIは、指定されたユーザーメッセージをOpenRouter APIに送信して、素敵なAIからの返信を取得するよ！ 😃💬
// 環境変数から秘密のAPIキーをゲットし、ユーザーメッセージを含むJSONリクエストを作成してHTTP POSTリクエストを送信するよ！ 📤🔥
// 問題があった場合はエラーメッセージを返してくれるので、安心して使えるよ！ 🤗👌
func callOpenRouterAPI(message string) (string, error) {
	// 🔑 APIキーを環境変数から取得するよ！
	apiKey := os.Getenv("OPENROUTER_API_KEY")
	if apiKey == "" {
		return "", fmt.Errorf("OPENROUTER_API_KEY is not set") // ❌ APIキーがない場合はエラー！
	}

	// 🌐 APIのURLを設定するよ！
	url := "https://openrouter.ai/api/v1/chat/completions"
	
	// 📦 リクエストの内容を作成するよ！
	reqBody := OpenRouterRequest{
		Model: "mistralai/mistral-7b-instruct", // 🤖 使うAIモデルを指定！
		Messages: []ChatMessage{
			{
				Role:    "user",
				Content: message, // 💬 ユーザーからのメッセージを設定！
			},
		},
	}

	// 🔄 JSONに変換するよ！
	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return "", err
	}

	// 📤 HTTPリクエストを作成するよ！
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return "", err
	}

	// 🎯 リクエストヘッダーを設定するよ！
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("HTTP-Referer", os.Getenv("APP_URL"))
	req.Header.Set("X-Title", "Family AI Bot")

	// 🚀 リクエストを送信するよ！
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// 📥 レスポンスの内容を読み取るよ！
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	// ⚠️ エラーチェックするよ！
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("API error: %s", string(body))
	}

	// 🎁 レスポンスを構造体に変換するよ！
	var openRouterResp OpenRouterResponse
	if err := json.Unmarshal(body, &openRouterResp); err != nil {
		return "", err
	}

	// ❌ 返信がない場合はエラー！
	if len(openRouterResp.Choices) == 0 {
		return "", fmt.Errorf("no response from the model")
	}

	// ✨ AIからの返信を返すよ！
	return openRouterResp.Choices[0].Message.Content, nil
}

// 💬🌟 チャットリクエストをスムーズに処理する魔法の関数だよ！
// handleChatは、HTTPリクエストからJSON形式のチャットリクエストを解析し、ユーザーのメッセージを抽出してOpenRouter APIに送信するよ！
// エラーがあった時は、丁寧なエラーメッセージで教えてくれるから、トラブルにもばっちり対応！ 😎👍
func handleChat(w http.ResponseWriter, r *http.Request) {
	// 📥 リクエストの内容を読み取るよ！
	var req ChatRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// 🤖 AIにメッセージを送って返信をもらうよ！
	aiResponse, err := callOpenRouterAPI(req.Message)
	if err != nil {
		http.Error(w, "Error calling AI API: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// 📦 レスポンスを作成するよ！
	response := ChatResponse{
		Response: aiResponse,
	}

	// 📤 レスポンスを送信するよ！
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// 📱💌 LINEのWebhookを受け取って、スマイルいっぱいの返信を返す関数だよ！
// handleLineWebhookは、LINEのWebhookリクエストを解析し、テキストメッセージイベントに応じてAIにお話ししてもらうよ！
// 署名が無効だったり、他のエラーが発生した場合は適切なHTTPステータスを返してトラブルを防ぐよ！ 😇🚦
// エラー情報はログに残して、安心のサポート体制を整えてるよ～！ 💼❤️
func handleLineWebhook(bot *linebot.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		// 🔍 LINEからのリクエストを解析するよ！
		events, err := bot.ParseRequest(req)
		if err != nil {
			if err == linebot.ErrInvalidSignature {
				w.WriteHeader(400) // ❌ 署名が間違ってるよ！
			} else {
				w.WriteHeader(500) // 💥 サーバーエラーだよ！
			}
			return
		}

		// 📨 イベントを処理するよ！
		for _, event := range events {
			if event.Type == linebot.EventTypeMessage {
				switch message := event.Message.(type) {
				case *linebot.TextMessage:
					// 🤖 AIにメッセージを送って返信をもらうよ！
					response, err := callOpenRouterAPI(message.Text)
					if err != nil {
						log.Printf("Error calling OpenRouter API: %v", err)
						continue
					}

					// 📤 LINEに返信を送るよ！
					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(response)).Do(); err != nil {
						log.Printf("Error sending response to LINE: %v", err)
					}
				}
			}
		}
	}
}

// 🚀🌟 メイン関数だよ！ Family AI Botのワクワクがここから始まるよ！
// mainはアプリケーションのエントリーポイントで、環境変数からLINE認証情報を読み込み、LINE Botクライアントを初期化するよ！ 🎉🤖
// Gorilla Muxルーターを使って、"/callback"エンドポイントにWebhookハンドラを設定し、指定またはデフォルトのポートでHTTPサーバーを起動するよ！ 🌐🎀
func main() {
	// 🔑 LINE Botの設定を読み込むよ！
	channelSecret := os.Getenv("LINE_CHANNEL_SECRET")
	channelToken := os.Getenv("LINE_CHANNEL_TOKEN")
	if channelSecret == "" || channelToken == "" {
		log.Fatal("LINE_CHANNEL_SECRET and LINE_CHANNEL_TOKEN must be set") // ❌ 設定がないと動かないよ！
	}

	// 🤖 LINE Botクライアントを作成するよ！
	bot, err := linebot.New(channelSecret, channelToken)
	if err != nil {
		log.Fatal(err)
	}

	// 🛣️ ルーターを設定するよ！
	r := mux.NewRouter()
	
	// 📱 LINE Webhookのエンドポイントを設定するよ！
	r.HandleFunc("/callback", handleLineWebhook(bot)).Methods("POST")

	// 🌐 サーバーを起動するよ！
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // 🎯 ポート番号のデフォルト値だよ！
	}
	
	// 🎉 サーバーを起動するよ！
	log.Printf("Server starting on port %s", port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatal(err)
	}
} 