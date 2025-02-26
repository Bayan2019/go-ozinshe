package views

type CreateProjectRequest struct {
	Title          string  `json:"title"`
	Description    string  `json:"description"`
	TypeID         int64   `json:"type_id"`
	DurationInMins int64   `json:"duration_in_mins"`
	ReleaseYear    int64   `json:"release_year"`
	Director       string  `json:"director"`
	Producer       string  `json:"producer"`
	GenreIds       []int64 `json:"genre_ids"`
	AgeCategoryIds []int64 `json:"age_category_ids"`
}

type UpdateProjectRequest struct {
	Title          string  `json:"title"`
	Description    string  `json:"description"`
	TypeID         int64   `json:"type_id"`
	DurationInMins int64   `json:"duration_in_mins"`
	ReleaseYear    int64   `json:"release_year"`
	Director       string  `json:"director"`
	Producer       string  `json:"producer"`
	GenreIds       []int64 `json:"genre_ids"`
	AgeCategoryIds []int64 `json:"age_category_ids"`
}
