package web

import (
	"go-infrastructure/web/constant"
	"go-infrastructure/web/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type response model.SystemJsonResponse
type result model.SystemResponseResult

/********************** common **********************/

func NewResponse(c *gin.Context) *response {
	return &response{Ctx: c}
}

/********************** success **********************/

func (s *response) Success(code *constant.MsgCode, data interface{}) {
	if code == nil {
		code = &constant.SUCCESS
	}
	s.Ctx.JSON(http.StatusOK, build(*code, data))
}

/********************** fail **********************/

func (s *response) Fail(code *constant.MsgCode, data interface{}) {
	if code == nil {
		code = &constant.COMMON_FAIL
	}
	s.Ctx.JSON(http.StatusOK, build(*code, data))
}

/********************** private **********************/

func build(code constant.MsgCode, data interface{}) *result {
	r := &result{}
	r.buildResult(code, data)
	return r
}

func (r *result) buildResult(code constant.MsgCode, data interface{}) {
	r.Code = code.Code
	r.Message = code.Msg
	r.Data = data
}
