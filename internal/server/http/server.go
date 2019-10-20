package http

import (
	"net/http"

	"github.com/greyireland/kratos-demo/internal/model"
	"github.com/greyireland/kratos-demo/internal/service"

	"fmt"

	"github.com/bilibili/kratos/pkg/conf/paladin"
	"github.com/bilibili/kratos/pkg/ecode"
	"github.com/bilibili/kratos/pkg/log"
	bm "github.com/bilibili/kratos/pkg/net/http/blademaster"
	"github.com/bilibili/kratos/pkg/net/http/blademaster/binding"
)

var (
	svc *service.Service
)

// New new a bm server.
func New(s *service.Service) (engine *bm.Engine) {
	var (
		hc struct {
			Server *bm.ServerConfig
		}
	)
	if err := paladin.Get("http.toml").UnmarshalTOML(&hc); err != nil {
		if err != paladin.ErrNotExist {
			panic(err)
		}
	}
	svc = s
	engine = bm.DefaultServer(hc.Server)
	initRouter(engine)
	if err := engine.Start(); err != nil {
		panic(err)
	}
	return
}

func initRouter(e *bm.Engine) {
	e.Ping(ping)
	g := e.Group("/v1/message")
	{
		g.POST("/add", addMessage)
		g.GET("/list", getMessages)
	}

}

func ping(ctx *bm.Context) {
	if err := svc.Ping(ctx); err != nil {
		log.Error("ping error(%v)", err)
		ctx.AbortWithStatus(http.StatusServiceUnavailable)
	}
}

// example for http request handler.
func howToStart(c *bm.Context) {
	k := &model.Kratos{
		Hello: "Golang 大法好 !!!",
	}
	c.JSON(k, nil)
}

func addMessage(c *bm.Context) {
	var msg model.MessageInfo
	err := binding.JSON.Bind(c.Request, &msg)
	fmt.Println(msg, err)
	if err != nil {
		log.Error("param error %s", err)
		c.JSON(nil, ecode.Int(499))
		return
	}
	data, err := svc.AddMessage(c.Context, msg)
	c.JSON(data, err)
}

func getMessages(c *bm.Context) {
	var msg model.MessageInfo
	err := binding.Form.Bind(c.Request, &msg)
	fmt.Println(msg, err)
	if err != nil {
		log.Error("param error %s", err)
		c.JSON(nil, ecode.Int(499))
		return
	}
	data, err := svc.GetMessages(c.Context, msg.UID, msg.PeerUID)
	c.JSON(data, err)
}
