package main

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

var port = ":8588"

type request struct {
	Jsonrpc string
	Id      int
	Method  string
	Params  []string
}

type clientVersion struct {
	Jsonrpc string
	Id      int
	Result  string
}

func requestHandler(res http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var r request
	err := decoder.Decode(&r)
	if err != nil {
		panic(err)
	}

	if r.Method == "web3_clientVersion" {
		var payload clientVersion
		payload.Jsonrpc = r.Jsonrpc
		payload.Id = r.Id
		payload.Result = "ReorgRpc/v1.0.0/linux/go1.12.1" // TODO - iquidus
		respondWithJson(res, http.StatusOK, payload)
	}

	log.Print(r.Method)
	log.Print(r.Params)
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func init() {
	v, _ := strconv.ParseBool(os.Getenv("DEBUG"))
	if v {
		log.SetFormatter(&log.TextFormatter{FullTimestamp: true, TimestampFormat: time.StampNano})
		log.SetLevel(log.DebugLevel)
		log.SetReportCaller(true)
	} else {
		log.SetFormatter(&log.TextFormatter{FullTimestamp: true, TimestampFormat: time.Stamp})
		log.SetLevel(log.InfoLevel)
	}
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", requestHandler).Methods("POST")
	handler := cors.Default().Handler(r)
	if err := http.ListenAndServe(port, handler); err != nil {
		log.Fatal("reorgrpc error: ", err)
	}
}
