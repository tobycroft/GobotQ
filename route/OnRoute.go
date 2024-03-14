package route

import (
	"github.com/bytedance/sonic"
	"github.com/gin-gonic/gin"
	Net "github.com/tobycroft/TuuzNet"
	"main.go/app/bot/message/index"
	"main.go/app/bot/model/LogErrorModel"
	"main.go/config/types"
	v1 "main.go/route/v1"
	"main.go/tuuz"
	"main.go/tuuz/Redis"
)

func OnRoute(router *gin.Engine) {
	router.Any("", func(context *gin.Context) {
		data, _ := context.GetRawData()
		//fmt.Println(string(data))
		ok := sonic.Valid(data)
		if ok {
			var es index.EventStruct
			err := sonic.Unmarshal(data, &es)
			if err != nil {
				LogErrorModel.Api_insert(err.Error(), tuuz.FUNCTION_ALL())
				return
			}
			mp := map[string]any{}
			err = sonic.Unmarshal(data, &mp)
			if err != nil {
				LogErrorModel.Api_insert(err.Error(), tuuz.FUNCTION_ALL())
				return
			}
			es.Json = mp
			es.RemoteAddr = context.RemoteIP() + ":80"

			Redis.PubSub{}.Publish_struct(types.MessageEvent, es)
		}
		context.String(200, "ok")
	})

	router.Any("/ws", func(c *gin.Context) {
		ws := Net.WsServer{}
		ws.NewServer(c.Writer, c.Request, nil)
	})

	version1 := router.Group("/v1")
	{
		version1.Use(func(context *gin.Context) {
		}, gin.Recovery())
		version1.Any("/", func(context *gin.Context) {
			context.String(0, version1.BasePath())
		})
		index := version1.Group("index")
		{
			v1.IndexRouter(index)
		}
		user := version1.Group("user")
		{
			v1.UserRouter(user)
		}
		bot := version1.Group("bot")
		{
			v1.BotRouter(bot)
		}
		group := version1.Group("group")
		{
			v1.GroupRouter(group)
		}
	}
}
