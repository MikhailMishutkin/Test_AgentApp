package repository

type Repository struct {
	s []string
}

func NewSliceRepo() *Repository {
	var s []string
	return &Repository{s: s}
}

func (r *Repository) Store() []string {
	s := r.s
	return s
	// !!!!!! нельзя из контроллера минуя юзкейс!!! _, b := r.Gys.GetList()
	// s := []model.MetroList{}
	// jsonErr := json.Unmarshal(b, &s)
	// if jsonErr != nil {
	// 	log.Fatal(jsonErr)
	// }

	// for _, v := range s {
	// 	m = append(m, v.CityName)
	// 	fmt.Print(m)
	// }
	// fmt.Print(m)

	// // for _, v := range m {
	// // 	fmt.Println(v)
	// // }

}
