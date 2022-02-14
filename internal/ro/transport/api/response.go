package api

import (
	"github.com/dbstarll/game/internal/ro/transport/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
}

func (r *Response) responseWithError(c *gin.Context, errorCode *model.ErrorCode, err error) {
	c.JSON(http.StatusOK, &model.Response{
		Code: errorCode.Code(),
		Ok:   errorCode.IsOk(),
		Data: map[string]string{"error": err.Error()},
		Msg:  errorCode.Desc(),
	})
}

func (r *Response) responseOkWithData(c *gin.Context, data interface{}) {
	errorCode := model.Error200
	c.JSON(http.StatusOK, &model.Response{
		Code: errorCode.Code(),
		Ok:   errorCode.IsOk(),
		Data: data,
		Msg:  errorCode.Desc(),
	})
}
