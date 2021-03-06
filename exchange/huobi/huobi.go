package huobi

import (
	"github.com/nntaoli-project/GoEx/huobi"
	"net/http"
	"github.com/nntaoli-project/GoEx"
	"github.com/locxiang/quantitative-trading/app/models"
	"github.com/locxiang/quantitative-trading/exchange/util"
	"github.com/lexkong/log"
)

type Okex struct {
	ApiKey    string
	SecretKey string
}

func (o *Okex) Init(tradeChan chan<- *models.Trade, pairs []goex.CurrencyPair) {

	var huoBiPro = huobi.NewHuoBiProSpot(http.DefaultClient, "", "")

	for _, pair := range pairs {
		huoBiPro.GetTradeWithWs(pair, func(trade *goex.Trade) {
			tradeChan <- util.ConvertTrade(huoBiPro.GetExchangeName(), trade)
		})
	}

	log.Info("huoBiPro 建立成功")

}
