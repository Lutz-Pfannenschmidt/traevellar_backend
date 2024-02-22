package api

import (
	"net/http"

	"github.com/Lutz-Pfannenschmidt/yagll"
)

type Api struct {
	Url string
}

func NewApi(url string) *Api {
	return &Api{Url: url}
}

func (a *Api) get(url string) (*http.Response, error) {
	yagll.Debugf("GET %s", a.Url+url)
	return http.Get(a.Url + url)
}

func (a *Api) buildQuery(url string, params ...map[string]string) string {
	res := "?"
	for _, param := range params {
		for key, value := range param {
			res += key + "=" + value + "&"
		}
	}
	res = res[:len(res)-1]
	return url + res
}
