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

func (pr *ProjectsRepository) GetAll(ctx context.Context) ([]views.RProject, error) {
	tx, err := pr.Conn.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()
	qtx := pr.DB.WithTx(tx)

	rProjects, err := qtx.GetProjects(ctx)
	if err != nil {
		return nil, err
	}

	projects := []views.RProject{}

	for _, p := range rProjects {
		type1, err := qtx.GetTypeById(ctx, p.TypeID)
		if err != nil {
			return nil, err
		}
		image := database.Image{}
		if p.Cover.Valid {
			imageId := p.Cover.String
			image, err = qtx.GetImage(ctx, imageId)
			if err != nil {
				return nil, err
			}
		}
		projects = append(projects, views.RProject{
			ID:             p.ID,
			CreatedAt:      p.CreatedAt,
			UpdatedAt:      p.UpdatedAt,
			Title:          p.Title,
			Description:    p.Description,
			Type:           type1,
			DurationInMins: p.DurationInMins,
			ReleaseYear:    p.ReleaseYear,
			Director:       p.Director,
			Producer:       p.Producer,
			Cover:          image,
		})
	}

	return projects, tx.Commit()
}

func (pr *ProjectsRepository) GetById(ctx context.Context, id int64) (views.Project, error) {
	tx, err := pr.Conn.Begin()
	if err != nil {
		return views.Project{}, err
	}
	defer tx.Rollback()
	qtx := pr.DB.WithTx(tx)

	rProject, err := qtx.GetProjectById(ctx, id)
	if err != nil {
		return views.Project{}, err
	}

	type1, err := qtx.GetTypeById(ctx, rProject.TypeID)
	if err != nil {
		return views.Project{}, err
	}

	image := database.Image{}
	if rProject.Cover.Valid {
		imageId := rProject.Cover.String
		image, err = qtx.GetImage(ctx, imageId)
		if err != nil {
			return views.Project{}, err
		}
	}

	genres, err := qtx.GetAllGenresOfProject(ctx, id)
	if err != nil {
		return views.Project{}, err
	}

	ageCategories, err := qtx.GetAllAgeCategoriesOfProject(ctx, id)
	if err != nil {
		return views.Project{}, err
	}

	images, err := qtx.GetImagesOfProject(ctx, id)
	if err != nil {
		return views.Project{}, err
	}

	videos, err := qtx.GetVideosOfProject(ctx, id)
	if err != nil {
		return views.Project{}, err
	}

	return views.Project{
		ID:             id,
		CreatedAt:      rProject.CreatedAt,
		UpdatedAt:      rProject.UpdatedAt,
		Title:          rProject.Title,
		Description:    rProject.Description,
		DurationInMins: rProject.DurationInMins,
		ReleaseYear:    rProject.ReleaseYear,
		Director:       rProject.Director,
		Producer:       rProject.Producer,
		Type:           type1,
		Cover:          image,
		Genres:         genres,
		AgeCategories:  ageCategories,
		Images:         images,
		Videos:         videos,
	}, tx.Commit()
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

func (pr *ProjectsRepository) UploadCover(ctx context.Context, project int64, cover string) error {
	tx, err := pr.Conn.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	qtx := pr.DB.WithTx(tx)

	err = qtx.AddImage2Movie(ctx, database.AddImage2MovieParams{
		ID:        cover,
		ProjectID: project,
	})
	if err != nil {
		return err
	}

	err = qtx.SetCover(ctx, database.SetCoverParams{
		ID: project,
		Cover: sql.NullString{
			String: cover,
			Valid:  true,
		},
	})
	if err != nil {
		return err
	}

	return tx.Commit()
}
