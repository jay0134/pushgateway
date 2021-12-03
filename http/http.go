package http

import (
	"eagleeye-pushgateway/g"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Dto struct {
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}


func RenderJson(w http.ResponseWriter, v interface{}) {
	bs, err := json.Marshal(v)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Write(bs)
}

func RenderDataJson(w http.ResponseWriter, data interface{}) {
	RenderJson(w, Dto{Msg: "success", Data: data})
}


func Start() {
	if !g.Config().Debug {
		gin.SetMode(gin.ReleaseMode)
	}

	configApiRoutes()

	addr := g.Config().Http.Listen
	if addr == "" {
		return
	}
	s := &http.Server{
		Addr:           addr,
		MaxHeaderBytes: 1 << 30,
	}
	s.ListenAndServe()

}
