package model

import "github.com/gin-gonic/gin"

type SystemJsonResponse struct {
	Ctx *gin.Context
}

type SystemResponseResult struct {
	Code    int         `json:"code"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}
