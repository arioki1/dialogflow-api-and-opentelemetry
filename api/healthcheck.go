package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (a *API) Healthcheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"ok": true,
	})
}
