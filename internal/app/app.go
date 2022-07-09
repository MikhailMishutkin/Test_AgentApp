package app

import (
	"agentapp/internal/adapter/repository"
	"agentapp/internal/usecase"
)

func Run() {
	//cl := usecase.NewHttp()
	m := repository.NewSliceRepo()
	useCase := usecase.New(m)
	c := usecase.NewCont(useCase)
	//a := c.GetList()
	c.Rand()
}
