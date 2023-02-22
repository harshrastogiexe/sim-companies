package cqrs

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/harshrastogiexe/cmd/server/pkg/core"
	"github.com/harshrastogiexe/cmd/server/pkg/global"
	proxycore "github.com/harshrastogiexe/pkg/sim-companies-proxy/core"
)

func GetAllResource(ctx *gin.Context) {
	app, found := global.Get[*core.Application](core.APPLICATION_TOKEN)
	if !found {
		panic(fmt.Errorf("instance not found: TOKEN = %s", core.APPLICATION_TOKEN))
	}
	r := []proxycore.ResourceBase{}
	app.DB.Table("resource_base").Find(&r)

	ctx.JSON(http.StatusOK, r)
}
