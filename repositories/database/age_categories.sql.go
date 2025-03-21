// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: age_categories.sql

package database

import (
	"context"
)

const createAgeCategory = `-- name: CreateAgeCategory :one

INSERT INTO age_categories(title)
VALUES (?)
RETURNING id
`

func (q *Queries) CreateAgeCategory(ctx context.Context, title string) (int64, error) {
	row := q.db.QueryRowContext(ctx, createAgeCategory, title)
	var id int64
	err := row.Scan(&id)
	return id, err
}

const deleteAgeCategory = `-- name: DeleteAgeCategory :exec

DELETE FROM age_categories WHERE id = ?
`

func (q *Queries) DeleteAgeCategory(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteAgeCategory, id)
	return err
}

const getAgeCategories = `-- name: GetAgeCategories :many
SELECT id, title FROM age_categories
`

func (q *Queries) GetAgeCategories(ctx context.Context) ([]AgeCategory, error) {
	rows, err := q.db.QueryContext(ctx, getAgeCategories)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []AgeCategory
	for rows.Next() {
		var i AgeCategory
		if err := rows.Scan(&i.ID, &i.Title); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAgeCategoryById = `-- name: GetAgeCategoryById :one

SELECT id, title FROM age_categories WHERE id = ?
`

func (q *Queries) GetAgeCategoryById(ctx context.Context, id int64) (AgeCategory, error) {
	row := q.db.QueryRowContext(ctx, getAgeCategoryById, id)
	var i AgeCategory
	err := row.Scan(&i.ID, &i.Title)
	return i, err
}

const getAllAgeCategoriesOfProject = `-- name: GetAllAgeCategoriesOfProject :many

SELECT ac.id, ac.title FROM age_categories AS ac
JOIN projects_age_categories AS pac
ON ac.id = pac.age_category_id
WHERE pac.project_id = ?
`

func (q *Queries) GetAllAgeCategoriesOfProject(ctx context.Context, projectID int64) ([]AgeCategory, error) {
	rows, err := q.db.QueryContext(ctx, getAllAgeCategoriesOfProject, projectID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []AgeCategory
	for rows.Next() {
		var i AgeCategory
		if err := rows.Scan(&i.ID, &i.Title); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateAgeCategory = `-- name: UpdateAgeCategory :exec

UPDATE age_categories 
SET title = ? 
WHERE id = ?
RETURNING id, title
`

type UpdateAgeCategoryParams struct {
	Title string
	ID    int64
}

func (q *Queries) UpdateAgeCategory(ctx context.Context, arg UpdateAgeCategoryParams) error {
	_, err := q.db.ExecContext(ctx, updateAgeCategory, arg.Title, arg.ID)
	return err
}
