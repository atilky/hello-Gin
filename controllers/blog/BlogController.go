package blog

import (
	"gindemo02/dto"
	"gindemo02/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type BlogController struct{}

func (b BlogController) Edit(ctx *gin.Context) {

	request := dto.UpdateRequest{}

	err := ctx.ShouldBind(&request)
	if err != nil {
		ctx.String(http.StatusBadRequest, "invalid parameter")
		return
	}

	blogModel := models.Blog{}
	_, err = blogModel.GetBlogById(request.Id)
	if err != nil {
		ctx.String(http.StatusNotFound, "博客不存在")
		return
	}

	blogModel.UpdateBlog(&models.Blog{
		Id:      request.Id,
		Title:   request.Title,
		Article: request.Article,
	})

	ctx.JSON(200, gin.H{
		"code": 200,
		"msg":  "edit blog success",
	})
}
