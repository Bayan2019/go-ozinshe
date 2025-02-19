-- name: GetMovies :many
SELECT * FROM movies;

-- name: GetMoviesByIdsIdAsc :many
SELECT * FROM movies
WHERE id = ANY(@ids::bigint[])
ORDER BY id ASC;

-- name: GetMoviesByIdsTitleAsc :many
SELECT * FROM movies
WHERE id = ANY(@ids::bigint[])
ORDER BY title ASC;

-- name: GetMoviesByIdsReleaseYearAsc :many
SELECT * FROM movies
WHERE id = ANY(@ids::bigint[])
ORDER BY release_year ASC;

-- name: GetMoviesByIdsRatingAsc :many
SELECT * FROM movies
WHERE id = ANY(@ids::bigint[])
ORDER BY title ASC;

-- name: GetMoviesByIdsIdDesc :many
SELECT * FROM movies
WHERE id = ANY(@ids::bigint[])
ORDER BY id DESC;

-- name: GetMoviesByIdsTitleDesc :many
SELECT * FROM movies
WHERE id = ANY(@ids::bigint[])
ORDER BY title DESC;

-- name: GetMoviesByIdsReleaseYearDesc :many
SELECT * FROM movies
WHERE id = ANY(@ids::bigint[])
ORDER BY release_year DESC;

-- name: GetMoviesByIdsRatingDesc :many
SELECT * FROM movies
WHERE id = ANY(@ids::bigint[])
ORDER BY title DESC;

-- name: GetMovieById :one
SELECT * FROM movies WHERE id = ?;

-- name: CreateMovie :one
INSERT INTO movies(title, description, release_year, director, trailer_url, poster_url)
VALUES (?, ?, ?, ?, ?, ?)
RETURNING *;

-- name: UpdateMovie :one
UPDATE movies 
SET title = ?, 
    description = ?, 
    release_year = ?, 
    director = ?, 
    rating = ?, 
    is_watched = ?, 
    trailer_url = ?,
    poster_url = ?
WHERE id = ?
RETURNING *;

-- name: UpdateMovieRating :exec
UPDATE movies
SET rating = ?
WHERE id = ?;

-- name: UpdateMovieIsWatched :exec
UPDATE movies
SET is_watched = ?
WHERE id = ?;

-- name: DeleteMovie :exec
DELETE FROM movies WHERE id = ?;