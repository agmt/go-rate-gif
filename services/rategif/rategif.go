package rategif

import (
	"fmt"
	"rategif/config"
	"rategif/services/openexchange"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/peterhellberg/giphy"
)

type API struct {
	Router *gin.Engine
	OE     *openexchange.OE
	Giphy  *giphy.Client
	Config *config.Config
}

func (api *API) GetRateGif(c *gin.Context) {
	symbol := c.Query("symbol")
	if len(symbol) == 0 {
		symbol = "RUB"
	}
	symbol = strings.ToUpper(symbol)

	now := time.Now()
	yday := now.AddDate(0, 0, -1)

	rateToday, err := api.OE.AtDate(symbol, now)
	if err != nil {
		c.AbortWithError(400, nil)
		return
	}

	rateYday, err := api.OE.AtDate(symbol, yday)
	if err != nil {
		c.AbortWithError(400, nil)
		return
	}

	giphyQuery := "okay"
	if rateToday.GreaterThan(rateYday) {
		giphyQuery = "rich"
	} else if rateToday.LessThan(rateYday) {
		giphyQuery = "broke"
	}

	s, err := api.Giphy.Search([]string{giphyQuery})
	if err != nil {
		c.AbortWithError(400, nil)
		return
	}

	giphyUrl := s.Data[0].Images.Original.URL
	html := fmt.Sprintf(
		`
<html>
<body>
<img src="%s"/>
</body>
</html>
	`,
		giphyUrl,
	)

	c.Data(200, gin.MIMEHTML, []byte(html))
}
