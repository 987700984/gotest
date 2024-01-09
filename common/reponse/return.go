package reponse

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

type Body struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func Reponse(w http.ResponseWriter, Code int, Msg string, resp interface{}) {

	//返回
	httpx.WriteJson(w, http.StatusOK, &Body{Code: Code, Msg: Msg, Data: resp})
}
