package repositories

import (
	"context"
	"database/sql"

	"github.com/Bayan2019/go-ozinshe/repositories/database"
	"github.com/Bayan2019/go-ozinshe/views"
)

type ProjectsRepository struct {
	Conn *sql.DB
	DB   *database.Queries
}

func NewProjectsRepository(db *sql.DB) *ProjectsRepository {
	return &ProjectsRepository{
		Conn: db,
		DB:   database.New(db),
	}
}

func (pr *ProjectsRepository) Create(ctx context.Context, cpr views.CreateProjectRequest) (int64, error) {
	tx, err := pr.Conn.Begin()
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	qtx := pr.DB.WithTx(tx)

	id, err := qtx.CreateProject(ctx, database.CreateProjectParams{
		Title:          cpr.Title,
		Description:    cpr.Description,
		TypeID:         cpr.TypeID,
		DurationInMins: cpr.DurationInMins,
		ReleaseYear:    cpr.ReleaseYear,
		Director:       cpr.Director,
		Producer:       cpr.Producer,
	})
	if err != nil {
		return 0, err
	}

	for _, genre_id := range cpr.GenreIds {
		err = qtx.AddGenre2Project(ctx, database.AddGenre2ProjectParams{
			ProjectID: id,
			GenreID:   genre_id,
		})
	}

	for _, age_category_id := range cpr.AgeCategoryIds {
		err = qtx.AddAgeCategory2Project(ctx, database.AddAgeCategory2ProjectParams{
			ProjectID:     id,
			AgeCategoryID: age_category_id,
		})
		if err != nil {
			return 0, err
		}
	}

	return id, tx.Commit()
}

func (pr *ProjectsRepository) Update(ctx context.Context, id int64, upr views.UpdateProjectRequest) error {
	tx, err := pr.Conn.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	qtx := pr.DB.WithTx(tx)

	err = qtx.UpdateProject(ctx, database.UpdateProjectParams{
		ID:             id,
		Title:          upr.Title,
		Description:    upr.Description,
		TypeID:         upr.TypeID,
		DurationInMins: upr.DurationInMins,
		ReleaseYear:    upr.ReleaseYear,
		Director:       upr.Director,
		Producer:       upr.Producer,
	})
	if err != nil {
		return err
	}

	err = qtx.DeleteAgeCategoriesOfProject(ctx, id)
	if err != nil {
		return err
	}

	err = qtx.DeleteGenresOfProject(ctx, id)
	if err != nil {
		return err
	}

	for _, genre_id := range upr.GenreIds {
		err = qtx.AddGenre2Project(ctx, database.AddGenre2ProjectParams{
			ProjectID: id,
			GenreID:   genre_id,
		})
		if err != nil {
			return err
		}
	}

	for _, genre_id := range upr.GenreIds {
		err = qtx.AddGenre2Project(ctx, database.AddGenre2ProjectParams{
			ProjectID: id,
			GenreID:   genre_id,
		})
		if err != nil {
			return err
		}
	}

	return tx.Commit()
}
