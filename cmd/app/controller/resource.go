package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/harshrastogiexe/cmd/repository"
)

type ResourceController struct {
	ResourceRepo repository.DBResourceRepository
}

func (c *ResourceController) GetAllResource(ctx *gin.Context) {
	r := c.ResourceRepo.ListResource()
	ctx.JSON(http.StatusOK, r)
}

func (c *ResourceController) GetResourceByID(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	r := c.ResourceRepo.GetResourceById(uint(id))
	if r == nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "resource not found"})
		return
	}
	ctx.JSON(http.StatusOK, r)
}
