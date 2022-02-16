package main

import (
	"context"
	"encoding/gob"
	"flag"
	"github.com/dbstarll/game/internal/ro/transport/api"
	"github.com/dbstarll/game/internal/ro/transport/model"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/memstore"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
	"net/http"
	"os"
)

var (
	ptrGinMode    = flag.String("ginMode", "debug", "Gin Mode: [debug|release|test]")
	ptrListenAddr = flag.String("listen", ":18003", "Listen Addr")
	ptrHealth     = flag.String("health", "/health", "Health check path")
)

func newZapLogger(lc fx.Lifecycle) (*zap.Logger, error) {
	var l *zap.Logger
	var err error
	if *ptrGinMode == "release" {
		l, err = zap.NewProduction()
	} else {
		l, err = zap.NewDevelopment()
	}

	if err != nil {
		return nil, err
	} else if recoverStd, err := zap.RedirectStdLogAt(l, zap.InfoLevel); err != nil {
		return nil, err
	} else {
		recoverGlobals := zap.ReplaceGlobals(l)
		lc.Append(fx.Hook{
			OnStop: func(ctx context.Context) error {
				recoverGlobals()
				recoverStd()
				return nil
			},
		})
		return l, nil
	}
}

func newFxLogger(l *zap.Logger) fxevent.Logger {
	return &fxevent.ZapLogger{Logger: l}
}

func newGinEngine() *gin.Engine {
	gin.SetMode(*ptrGinMode)
	return gin.New()
}

func newMemStore() memstore.Store {
	gob.Register(model.PlayerModel{})
	return memstore.NewStore([]byte("secret"))
}

func newPlayerDispatch() (*api.PlayerDispatch, error) {
	return api.NewPlayerDispatch()
}

func initGin(g *gin.Engine, store memstore.Store) {
	g.Use(gin.Recovery())
	g.Use(gin.LoggerWithWriter(os.Stdout, *ptrHealth))
	g.Use(sessions.Sessions("ro-sid", store))

	// 404 Handler.
	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect API route.")
	})
	g.GET(*ptrHealth, func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})
	g.Static("/static", "./web/static")
	g.Static("/ro", "./web/ro")
}

func bindPlayerDispatch(c *api.PlayerDispatch, g *gin.Engine) {
	c.BindGin(g.Group("player"))
}

func listenAndServe(g *gin.Engine) {
	go func() {
		if err := g.Run(*ptrListenAddr); err != nil {
			zap.L().Error("gin.Engine.Run failed", zap.String("listenAddr", *ptrListenAddr), zap.Error(err))
		}
	}()
}

func main() {
	flag.Parse()

	fx.New(
		fx.WithLogger(newFxLogger),
		fx.Provide(
			newZapLogger,
			newGinEngine,
			newMemStore,
			newPlayerDispatch,
		),
		fx.Invoke(
			initGin,
			bindPlayerDispatch,
			listenAndServe,
		),
	).Run()
}
