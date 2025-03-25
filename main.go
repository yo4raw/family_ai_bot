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

type ChatRequest struct {
	Message string `json:"message"`
}

type ChatResponse struct {
	Response string `json:"response"`
}

type OpenRouterRequest struct {
	Model    string        `json:"model"`
	Messages []ChatMessage `json:"messages"`
}

type ChatMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type OpenRouterResponse struct {
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
}

func callOpenRouterAPI(message string) (string, error) {
	apiKey := os.Getenv("OPENROUTER_API_KEY")
	if apiKey == "" {
		return "", fmt.Errorf("OPENROUTER_API_KEY is not set")
	}

	url := "https://openrouter.ai/api/v1/chat/completions"
	
	reqBody := OpenRouterRequest{
		Model: "mistralai/mistral-7b-instruct",
		Messages: []ChatMessage{
			{
				Role:    "user",
				Content: message,
			},
		},
	}

	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("HTTP-Referer", os.Getenv("APP_URL"))
	req.Header.Set("X-Title", "Family AI Bot")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("API error: %s", string(body))
	}

	var openRouterResp OpenRouterResponse
	if err := json.Unmarshal(body, &openRouterResp); err != nil {
		return "", err
	}

	if len(openRouterResp.Choices) == 0 {
		return "", fmt.Errorf("no response from the model")
	}

	return openRouterResp.Choices[0].Message.Content, nil
}

func handleChat(w http.ResponseWriter, r *http.Request) {
	var req ChatRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	aiResponse, err := callOpenRouterAPI(req.Message)
	if err != nil {
		http.Error(w, "Error calling AI API: "+err.Error(), http.StatusInternalServerError)
		return
	}

	response := ChatResponse{
		Response: aiResponse,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func handleLineWebhook(bot *linebot.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		events, err := bot.ParseRequest(req)
		if err != nil {
			if err == linebot.ErrInvalidSignature {
				w.WriteHeader(400)
			} else {
				w.WriteHeader(500)
			}
			return
		}

		for _, event := range events {
			if event.Type == linebot.EventTypeMessage {
				switch message := event.Message.(type) {
				case *linebot.TextMessage:
					// OpenRouter APIを呼び出して応答を取得
					response, err := callOpenRouterAPI(message.Text)
					if err != nil {
						log.Printf("Error calling OpenRouter API: %v", err)
						continue
					}

					// LINEに応答を送信
					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(response)).Do(); err != nil {
						log.Printf("Error sending response to LINE: %v", err)
					}
				}
			}
		}
	}
}

func main() {
	// LINE Botクライアントの初期化
	channelSecret := os.Getenv("LINE_CHANNEL_SECRET")
	channelToken := os.Getenv("LINE_CHANNEL_TOKEN")
	if channelSecret == "" || channelToken == "" {
		log.Fatal("LINE_CHANNEL_SECRET and LINE_CHANNEL_TOKEN must be set")
	}

	bot, err := linebot.New(channelSecret, channelToken)
	if err != nil {
		log.Fatal(err)
	}

	// ルーターの設定
	r := mux.NewRouter()
	
	// LINE Webhookのエンドポイント
	r.HandleFunc("/callback", handleLineWebhook(bot)).Methods("POST")

	// サーバーの起動
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	
	log.Printf("Server starting on port %s", port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatal(err)
	}
} 