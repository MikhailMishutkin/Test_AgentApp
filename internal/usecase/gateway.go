package usecase

import (
	"agentapp/internal/model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"sort"
	"time"
)

type Controller struct {
	Gys GetYStationer
	cl  *http.Client
}

// New
func NewCont(gys GetYStationer) *Controller {
	return &Controller{Gys: gys}
}

// addr api all
func (c *Controller) GetApiAll() string {
	q := c.Gys.List()
	return q
}

// addr api stations in city
func (c *Controller) GetApiCit() string {
	id := c.GetLS()
	q := c.Gys.ListS(id)
	return q
}

//do rqst
func (c *Controller) DoReq(q string) []byte {
	req := c.Gys.Reguest(q)
	c.cl = http.DefaultClient
	resp, err := c.cl.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	return body
}

//json unmarsh
func (c *Controller) JsA(body []byte, s []model.MetroList) []model.MetroList {
	jsonErr := json.Unmarshal(body, &s)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	return s
}

func (c *Controller) JsS(body []byte, s model.MetroList) model.MetroList {
	jsonErr := json.Unmarshal(body, &s)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	return s
}

//list of cities
func (c *Controller) GetList() []string {
	q := c.GetApiAll()
	body := c.DoReq(q)
	var s []model.MetroList
	s = c.JsA(body, s)
	m := c.Gys.PushSlice()

	for _, v := range s {
		m = append(m, v.CityName)
	}

	sort.StringSlice(m).Sort()

	return m
}

// list id cities
func (c *Controller) GetLS() string {
	q := c.GetApiAll()
	body := c.DoReq(q)
	var s []model.MetroList
	s = c.JsA(body, s)

	m := c.GetList()
	check := c.Gys.CheckAnswer(c.OnScr(), m)
	var id string

	for _, v := range s {
		if v.CityName == check {
			id = v.CityId
		}

	}

	return id
}

//list cts on screen
func (c *Controller) OnScr() string {
	fmt.Println("--Выберите город из перечня с сайта hh.ru, в котором хотите работать---")
	c.Gys.GetSlice(c.GetList())
	var p string
	fmt.Scan(&p)
	return p
}

// rand station in response
func (c *Controller) Rand() {
	rand.Seed(time.Now().UnixNano()) // активировали рандомайзер

	q := c.GetApiCit()
	body := c.DoReq(q)
	var s model.MetroList
	s = c.JsS(body, s)

	m := s.Lines[0].Stations
	n := rand.Intn(len(m))
	fmt.Println("Станция, на которой у вас будет прекрасная работа: ", s.Lines[0].Stations[n].StName)
}

// func (c *Controller) Choice() {
// 	cs := c.Gys.GetFromS()
// 	for _, v := range cs {
// 		fmt.Println(v.CityName)
// 	}
// }
