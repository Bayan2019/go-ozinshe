-- name: AddAgeCategory2Project :exec
INSERT INTO projects_age_categories(project_id, age_category_id)
VALUES (?, ?);
--

-- name: DeleteAgeCategoriesOfProject :exec
DELETE FROM projects_age_categories WHERE project_id = ?;
--