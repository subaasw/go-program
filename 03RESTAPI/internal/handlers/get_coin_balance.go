package handlers

import (
	"encoding/json"
	"goapi/api"
	"goapi/internal/tools"
	"net/http"

	"github.com/gorilla/schema"
	log "github.com/sirupsen/logrus"
)

func GetCoinBalance(w http.ResponseWriter, r *http.Request) {
	var params = api.CoinBalanceParams{}
	var decoder *schema.Decoder = schema.NewDecoder()
	var err error

	err = decoder.Decode(&params, r.URL.Query())

	if err != nil {
		log.Error(err)
		api.InternalErrorhandler(w)
		return
	}

	var database *tools.DatabaseInterface
	database, err = tools.NewDatabase()

	if err != nil {
		api.InternalErrorhandler(w)
		return
	}

	var tokenDetails *tools.CoinDetails = (*database).GetUserCoins(params.Username)
	if (tokenDetails).Coins == 0 {
		log.Error(err)
		api.InternalErrorhandler(w)
	}

	var response = api.CoinBalanceResponse{
		Balance: tokenDetails.Coins,
		Code:    http.StatusOK,
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Error(err)
		api.InternalErrorhandler(w)
		return
	}
}
