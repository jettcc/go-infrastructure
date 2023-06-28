package web

import (
	"go-infrastructure/web/constant"
	"go-infrastructure/web/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type response model.SystemJsonResponse
type result model.SystemResponseResult

func NewResponse(c *gin.Context) *response {
	return &response{Ctx: c}
}

func (s *response) Success(data interface{}) {
	r := &result{}
	r.buildResult(constant.SUCCESS, data)
	s.Ctx.JSON(http.StatusOK, r)
}

func (s *response) Fail() {
	r := &result{}
	r.buildResult(constant.COMMON_FAIL, nil)
	s.Ctx.JSON(http.StatusOK, r)
}

func (r *result) buildResult(code constant.MsgCode, data interface{}) {
	r.Code = code.Code
	r.Message = code.Msg
	r.Data = data
}
