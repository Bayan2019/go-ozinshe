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

func (pr *ProjectsRepository) GetAll(ctx context.Context) ([]views.Project, error) {
	tx, err := pr.Conn.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()
	qtx := pr.DB.WithTx(tx)

	dProjects, err := qtx.GetProjects(ctx)
	if err != nil {
		return nil, err
	}

	projects, err := pr.DatabaseProjects2viewsProjects(ctx, dProjects)
	if err != nil {
		return nil, err
	}

	// projects := []views.Project{}

	return projects, tx.Commit()
}

func (pr *ProjectsRepository) GetById(ctx context.Context, id int64) (views.Project, error) {
	tx, err := pr.Conn.Begin()
	if err != nil {
		return views.Project{}, err
	}
	defer tx.Rollback()
	qtx := pr.DB.WithTx(tx)

	dProject, err := qtx.GetProjectById(ctx, id)
	if err != nil {
		return views.Project{}, err
	}

	project, err := pr.DatabaseProject2viewsProject(ctx, dProject)
	if err != nil {
		return views.Project{}, err
	}

	return project, tx.Commit()
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
		Keywords:       cpr.Keywords,
	})
	if err != nil {
		return 0, err
	}

	for _, genre_id := range cpr.GenreIds {
		err = qtx.AddGenre2Project(ctx, database.AddGenre2ProjectParams{
			ProjectID: id,
			GenreID:   genre_id,
		})
		if err != nil {
			return 0, err
		}
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
		Keywords:       upr.Keywords,
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

func (pr *ProjectsRepository) DatabaseProject2viewsProject(ctx context.Context, dProject database.Project) (views.Project, error) {
	vProject := views.Project{}
	tx, err := pr.Conn.Begin()
	if err != nil {
		return vProject, err
	}
	defer tx.Rollback()

	qtx := pr.DB.WithTx(tx)

	typ, err := qtx.GetTypeById(ctx, dProject.TypeID)
	if err != nil {
		return vProject, err
	}
	vProject.Type = typ

	// var image database.Image
	if dProject.Cover.Valid {
		image, err := qtx.GetImage(ctx, dProject.Cover.String)
		if err != nil {
			return vProject, err
		}
		vProject.Cover = image
	}

	genres, err := qtx.GetAllGenresOfProject(ctx, dProject.ID)
	if err != nil {
		return vProject, err
	}
	vProject.Genres = genres

	ageCategories, err := pr.DB.GetAllAgeCategoriesOfProject(ctx, dProject.ID)
	if err != nil {
		return vProject, err
	}
	vProject.AgeCategories = ageCategories

	images, err := pr.DB.GetImagesOfProject(ctx, dProject.ID)
	if err != nil {
		return vProject, err
	}
	vProject.Images = images

	videos, err := pr.DB.GetVideosOfProject(ctx, dProject.ID)
	if err != nil {
		return vProject, err
	}
	vProject.Videos = videos

	vProject.ID = dProject.ID
	vProject.CreatedAt = dProject.CreatedAt
	vProject.UpdatedAt = dProject.UpdatedAt
	vProject.Title = dProject.Title
	vProject.Description = dProject.Description
	vProject.DurationInMins = dProject.DurationInMins
	vProject.ReleaseYear = dProject.ReleaseYear
	vProject.Director = dProject.Director
	vProject.Producer = dProject.Producer
	vProject.Keywords = dProject.Keywords

	return vProject, tx.Commit()
}

func (pr *ProjectsRepository) DatabaseProjects2viewsProjects(ctx context.Context, dprojects []database.Project) ([]views.Project, error) {
	tx, err := pr.Conn.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	// qtx := pr.DB.WithTx(tx)
	vProjects := []views.Project{}
	for _, p := range dprojects {
		vProject, err := pr.DatabaseProject2viewsProject(ctx, p)
		if err != nil {
			return vProjects, err
		}
		vProjects = append(vProjects, vProject)
	}
	return vProjects, tx.Commit()
}
