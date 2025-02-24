package views

type Role struct {
	ID            int64  `json:"id"`
	Title         string `json:"title"`
	Projects      int64  `json:"projects"`
	Genres        int64  `json:"genres"`
	AgeCategories int64  `json:"age_categories"`
	Types         int64  `json:"types"`
	Users         int64  `json:"users"`
	Roles         int64  `json:"roles"`
}
