package api

import (
	"encoding/json"
	"fmt"
	"github.com/dbstarll/game/internal/ro/transport/model"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"time"
)

type PlayerDispatch struct {
	*Response
}

func NewPlayerDispatch() (*PlayerDispatch, error) {
	return &PlayerDispatch{}, nil
}

func (d *PlayerDispatch) BindGin(g *gin.RouterGroup) {
	g.GET("load", d.load)
	g.POST("save", d.save)
	g.GET("download", d.download)
	g.POST("upload", d.upload)
}

func (d *PlayerDispatch) load(c *gin.Context) {
	if player := sessions.Default(c).Get("player"); player == nil {
		d.responseWithError(c, model.Error404, errors.New("player not found"))
	} else {
		d.responseOkWithData(c, player)
	}
}

func (d *PlayerDispatch) save(c *gin.Context) {
	player := &model.PlayerModel{}
	if err := c.BindJSON(player); err != nil {
		d.responseWithError(c, model.Error500, err)
	} else {
		session := sessions.Default(c)
		session.Set("player", player)
		session.Save()
		d.responseOkWithData(c, nil)
	}
}

func (d *PlayerDispatch) download(c *gin.Context) {
	if playerObj := sessions.Default(c).Get("player"); playerObj == nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else if player, ok := playerObj.(model.PlayerModel); !ok {
		c.AbortWithStatus(http.StatusNotFound)
	} else if data, err := json.MarshalIndent(player, "", "  "); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	} else {
		filename := fmt.Sprintf("%s-%s.json", player.CharacterName, time.Now().Format("20060102150405"))
		c.Header("Content-Disposition", "attachment; filename="+filename)
		c.Header("Content-Transfer-Encoding", "binary")
		c.Data(http.StatusOK, "application/octet-stream", data)
	}
}

func (d *PlayerDispatch) upload(c *gin.Context) {
	player := &model.PlayerModel{}
	if file, _, err := c.Request.FormFile("player"); err != nil {
		d.responseWithError(c, model.Error500, err)
	} else if data, err := ioutil.ReadAll(file); err != nil {
		d.responseWithError(c, model.Error500, err)
	} else if err := json.Unmarshal(data, player); err != nil {
		d.responseWithError(c, model.Error500, err)
	} else {
		session := sessions.Default(c)
		session.Set("player", player)
		session.Save()
		d.responseOkWithData(c, player)
	}
}
