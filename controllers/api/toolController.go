package api

import (
	"gindemo02/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ToolsController struct {
}

func (t ToolsController) GetTimeStr(c *gin.Context) {

	timestampStr := c.Query("timestamp")
	timestamp, _ := strconv.Atoi(timestampStr)
	c.JSON(200, gin.H{
		"timestamp": "hello:" + timestampStr,
		"timeStr":   models.UnixToTime(timestamp),
	})

}
