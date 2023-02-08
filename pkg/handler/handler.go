package handler

import (
	"afet-yardim-twitter-bot/pkg/endpoint"
	apiError "afet-yardim-twitter-bot/pkg/error"
	"afet-yardim-twitter-bot/pkg/structure"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

var ErrQueryParam = fmt.Errorf("query param is not ok")

func NewHTTPHandler(endpoints endpoint.Endpoints) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/retweet", RetweetHandler(endpoints, http.MethodGet))

	return mux
}

func RetweetHandler(endpoints endpoint.Endpoints, method string) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != method {
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte("Method not allowed"))
		}
		queryNameID := r.URL.Query().Get("id")

		if queryNameID == "" {
			RespondWithError(w, http.StatusBadRequest, ErrQueryParam.Error())
			return
		}

		tweetID, err := strconv.Atoi(queryNameID)
		if err != nil {
			RespondWithError(w, http.StatusBadRequest, ErrQueryParam.Error())
			return
		}

		rawResp, err := endpoints.RetweetEndpoint(context.Background(), structure.RetweetRequest{TweetID: tweetID})
		if err != nil {
			apiErr := err.(apiError.ApiError)
			RespondWithError(w, apiErr.StatusCode, apiErr.BaseError.Error())
			return
		}

		response := rawResp.(structure.RetweetResponse)
		RespondWithJSON(w, http.StatusOK, response)
	}
}

func RespondWithError(w http.ResponseWriter, code int, message string) {
	RespondWithJSON(w, code, map[string]string{"error": message})
}

func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
