package main

import (
	"fmt"
	"log"
	"net/http"
	"runtime/debug"

	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/linebot"
)

const (
	secret = "2bf8fa7ecd6a10f7d467b380040653e2"
	token  = "mP9ehjSoLoFo7PHY9O0qqkCerXfM7tKTVRhziAz0X2LLgxpiM3hOqTeY2Lp7arBDH6Pb/AocexieMVttnXEtlonHGFUsa5ehceyby6JH2rHhmGplA7/CjryhpTi7UNT8Aku89faTS2Eq8LywsQoolwdB04t89/1O/w1cDnyilFU="
)

var botInstance *linebot.Client

func init() {
	InitLineBot()
}

func main() {
	defer func() {
		fmt.Println("Server shutdown...")
		if err := recover(); err != nil {
			fmt.Println(err, string(debug.Stack()))
		}
	}()

	router := gin.Default()

	router.GET("/", func(ctx *gin.Context) {
		req := ctx.Request
		events, err := botInstance.ParseRequest(req)
		if err != nil {
			if err == linebot.ErrInvalidSignature {
				ctx.JSON(http.StatusBadRequest, err)
			} else {
				ctx.JSON(http.StatusInternalServerError, err)
			}
			return
		}
		for _, event := range events {
			if event.Type == linebot.EventTypeMessage {
				switch message := event.Message.(type) {
				case *linebot.TextMessage:
					if _, err = botInstance.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message.Text)).Do(); err != nil {
						ctx.JSON(http.StatusInternalServerError, err)
						return
					}
				case *linebot.StickerMessage:
					replyMessage := fmt.Sprintf(
						"sticker id is %s, stickerResourceType is %s", message.StickerID, message.StickerResourceType)
					if _, err = botInstance.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(replyMessage)).Do(); err != nil {
						ctx.JSON(http.StatusInternalServerError, err)
						return
					}
				}
			}
		}

		ctx.JSON(http.StatusOK, map[string]interface{}{"msg": "ok"})
	})

	router.Run()
}

func InitLineBot() error {
	log.Println("init line bot")
	var err error
	botInstance, err = linebot.New(secret, token)

	return err
}
