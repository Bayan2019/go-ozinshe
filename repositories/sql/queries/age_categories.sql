-- name: GetAgeCategories :many
SELECT * FROM age_categories;
--

-- name: GetAgeCategoryById :one
SELECT * FROM age_categories WHERE id = ?;
--

-- name: CreateAgeCategory :one
INSERT INTO age_categories(title)
VALUES (?)
RETURNING id;
--

-- name: UpdateAgeCategory :exec
UPDATE age_categories 
SET title = ? 
WHERE id = ?
RETURNING *;
--

-- name: DeleteAgeCategory :exec
DELETE FROM age_categories WHERE id = ?;
--

-- name: GetAllAgeCategoriesOfProject :many
SELECT ac.* FROM age_categories AS ac
JOIN projects_age_categories AS pac
ON ac.id = pac.age_category_id
WHERE pac.project_id = ?;
--