package main

import (
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/gin-gonic/gin"
)

const (
	secret = "2bf8fa7ecd6a10f7d467b380040653e2"
	token  = "mP9ehjSoLoFo7PHY9O0qqkCerXfM7tKTVRhziAz0X2LLgxpiM3hOqTeY2Lp7arBDH6Pb/AocexieMVttnXEtlonHGFUsa5ehceyby6JH2rHhmGplA7/CjryhpTi7UNT8Aku89faTS2Eq8LywsQoolwdB04t89/1O/w1cDnyilFU="
)

func main() {
	defer func() {
		fmt.Println("Server shutdown...")
		if err := recover(); err != nil {
			fmt.Println(err, string(debug.Stack()))
		}
	}()

	router := gin.Default()
	router.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "^_^")
	})
	router.Run()

	//log.Println("init line bot")
	//bot, err := linebot.New(secret, token)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//log.Println("set line bot callback func")
	//http.HandleFunc("/callback", func(w http.ResponseWriter, req *http.Request) {
	//	events, err := bot.ParseRequest(req)
	//	if err != nil {
	//		if err == linebot.ErrInvalidSignature {
	//			w.WriteHeader(400)
	//		} else {
	//			w.WriteHeader(500)
	//		}
	//		return
	//	}
	//	for _, event := range events {
	//		if event.Type == linebot.EventTypeMessage {
	//			switch message := event.Message.(type) {
	//			case *linebot.TextMessage:
	//				if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message.Text)).Do(); err != nil {
	//					log.Print(err)
	//				}
	//			case *linebot.StickerMessage:
	//				replyMessage := fmt.Sprintf(
	//					"sticker id is %s, stickerResourceType is %s", message.StickerID, message.StickerResourceType)
	//				if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(replyMessage)).Do(); err != nil {
	//					log.Print(err)
	//				}
	//			}
	//		}
	//	}
	//})
}
