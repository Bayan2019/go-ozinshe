-- +goose Up
CREATE TABLE projects_age_categories(
    project_id INTEGER NOT NULL REFERENCES projects(id) ON DELETE CASCADE,
    age_category_id INTEGER NOT NULL REFERENCES age_categories(id) ON DELETE CASCADE,
    UNIQUE(project_id, age_category_id)
);

-- +goose Down
DROP TABLE projects_age_categories;