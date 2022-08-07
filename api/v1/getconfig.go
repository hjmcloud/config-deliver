package v1

import "github.com/gogf/gf/v2/frame/g"

type GetConfigReq struct {
	g.Meta   `path:"/getconfig" tags:"GetConfig" method:"get" summary:"GetConfig"`
	Id       string `p:"id" v:"required#id is required"` // id is required
	Filename string `p:"file"`                           // file is optional
	Dl       string `p:"dl"`                             // dl is optional
}
type GetConfigRes struct {
	g.Meta  `mime:"text/html" example:"string"`
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
