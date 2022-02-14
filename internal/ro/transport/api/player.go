package api

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type PlayerDispatch struct {
	*Response
}

func NewPlayerDispatch() (*PlayerDispatch, error) {
	return &PlayerDispatch{}, nil
}

func (d *PlayerDispatch) BindGin(g *gin.RouterGroup) {
	g.POST("load", d.load)
}

func (d *PlayerDispatch) load(c *gin.Context) {
	d.responseOkWithData(c, sessions.Default(c).Get("player"))
}
