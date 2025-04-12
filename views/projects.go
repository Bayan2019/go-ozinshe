package views

import (
	"github.com/Bayan2019/go-ozinshe/repositories/database"
)

// type RProject struct {
// 	ID             int64          `json:"id"`
// 	CreatedAt      string         `json:"created_at"`
// 	UpdatedAt      string         `json:"updated_at"`
// 	Title          string         `json:"title"`
// 	Description    string         `json:"description"`
// 	Type           database.Type  `json:"type"`
// 	DurationInMins int64          `json:"duration_in_mins"`
// 	ReleaseYear    int64          `json:"release_year"`
// 	Director       string         `json:"director"`
// 	Producer       string         `json:"producer"`
// 	Keywords       string         `json:"keywords"`
// 	Cover          database.Image `json:"cover"`
// }

type Project struct {
	ID             int64                  `json:"id"`
	CreatedAt      string                 `json:"created_at"`
	UpdatedAt      string                 `json:"updated_at"`
	Title          string                 `json:"title"`
	Description    string                 `json:"description"`
	Type           database.Type          `json:"type"`
	DurationInMins int64                  `json:"duration_in_mins"`
	ReleaseYear    int64                  `json:"release_year"`
	Director       string                 `json:"director"`
	Producer       string                 `json:"producer"`
	Keywords       string                 `json:"keywords"`
	Cover          database.Image         `json:"cover"`
	Genres         []database.Genre       `json:"genres"`
	AgeCategories  []database.AgeCategory `json:"age_categories"`
	Images         []database.Image       `json:"images"`
	Videos         []database.Video       `json:"videos"`
}

type CreateProjectRequest struct {
	Title          string  `json:"title"`
	Description    string  `json:"description"`
	TypeID         int64   `json:"type_id"`
	DurationInMins int64   `json:"duration_in_mins"`
	ReleaseYear    int64   `json:"release_year"`
	Director       string  `json:"director"`
	Producer       string  `json:"producer"`
	Keywords       string  `json:"keywords"`
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
	Keywords       string  `json:"keywords"`
	GenreIds       []int64 `json:"genre_ids"`
	AgeCategoryIds []int64 `json:"age_category_ids"`
}
