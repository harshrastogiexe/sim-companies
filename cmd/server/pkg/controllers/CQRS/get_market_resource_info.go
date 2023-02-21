package cqrs

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/harshrastogiexe/cmd/server/pkg/models"
	simcompanies "github.com/harshrastogiexe/pkg/sim-companies-proxy"
)

func GetResourceMarketInfo(ctx *gin.Context) {
	id := ctx.Param("id")
	fmt.Println(id)
	d, err := simcompanies.Market.GetMarketData(id)
	if err != nil {
		code := http.StatusInternalServerError
		ctx.JSON(code, models.NewApiError(code, err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, d)
}
