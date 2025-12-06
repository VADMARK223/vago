package handler

import (
	"net/http"
	"vago/internal/app"
	"vago/internal/moex"

	"github.com/gin-gonic/gin"
)

func ShowMoex(c *gin.Context) {
	cli := moex.NewClient()

	info, market, err := cli.GetBond("SU26225RMFS1")

	if err != nil {
		c.String(500, err.Error())
		return
	}

	app.Dump("info", info)
	app.Dump("market", market)

	dataOut := tplWithCapture(c, "Мос биржа")
	c.HTML(http.StatusOK, "moex.html", dataOut)
}
