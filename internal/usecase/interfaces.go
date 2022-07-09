package usecase

import (
	"agentapp/internal/model"
	"net/http"
)

type ListAllRepo interface {
	Store() []string
	//map[int]model.MetroList
}

type GetYStationer interface {
	List() string
	Reguest(q string) *http.Request
	PushSlice() []string
	GetSlice(m []string)
	CheckAnswer(s string, m []string) string
	ListS(id string) string
}

type GYSContoroller interface {
	GetApiAll() string
	GetApiCit() string
	DoReq(q string) []byte
	JsA(body []byte, s []model.MetroList) []model.MetroList
	JsS(body []byte, s model.MetroList) model.MetroList
	GetList() []string
	GetLS() string
	OnScr()
	Rand()
}
