package usecase

import (
	"fmt"
	"net/http"
)

const (
	ApiURL = "https://api.hh.ru/metro"
)

type GetYStation struct {
	R ListAllRepo
}

func New(r ListAllRepo) *GetYStation {
	return &GetYStation{R: r}

}

func (gys *GetYStation) List() string {
	return ApiURL
}

func (gys *GetYStation) Reguest(q string) *http.Request {
	req, err := http.NewRequest(http.MethodGet, q, nil)
	if err != nil {
		panic(err)
	}
	return req
}

func (gys *GetYStation) PushSlice() []string {
	s := gys.R.Store()
	return s
}

func (gys *GetYStation) GetSlice(m []string) {
	for _, v := range m {
		fmt.Println(v)
	}
}

func (gys *GetYStation) CheckAnswer(s string, m []string) string {
	var f string
	for i := range m {
		if m[i] != s {
			//fmt.Println("поробуйте ещё раз")
			continue
		} else {
			f = m[i]
			break
		}
	}
	return f
}

func (gys *GetYStation) ListS(id string) string {
	secApi := ApiURL + "/" + id
	return secApi
}
