package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/k1/go-blog/pkg/errcode"
)

type Reponse struct {
	Ctx *gin.Context
}

type Pager struct {
	Page      int `json:"page"`
	PageSize  int `json:"page_size"`
	TotalRows int `json:"total_rows"`
}

func NewReponse(ctx *gin.Context) *Reponse {
	return &Reponse{Ctx: ctx}
}

func (r *Reponse) ToResponse(data interface{}) {
	if data == nil {
		data = gin.H{}
	}
	r.Ctx.JSON(http.StatusOK, data)
}

func (r *Reponse) ToResponseList(list interface{}, totalRows int) {
	r.Ctx.JSON(
		http.StatusOK,
		gin.H{
			"list": list,
			"pager": Pager{
				Page:      GetPage(r.Ctx),
				PageSize:  GetPageSize(r.Ctx),
				TotalRows: totalRows,
			},
		},
	)
}

func (r *Reponse) ToErrorResponse(err *errcode.Error) {
	response := gin.H{
		"code":    err.Code,
		"msg":     err.Msg,
		"details": err.Details,
	}
	r.Ctx.JSON(err.StatusCode(), response)
}
