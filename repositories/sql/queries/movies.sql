-- name: GetMovies :many
SELECT * FROM movies;
--

-- name: GetMovieById :one
SELECT * FROM movies WHERE id = ?;
--

-- name: CreateMovie :one
INSERT INTO movies(title, description, release_year, director, trailer_url, poster_url)
VALUES (?, ?, ?, ?, ?, ?)
RETURNING *;
--

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
--

-- name: UpdateMovieRating :exec
UPDATE movies
SET rating = ?
WHERE id = ?;
--

-- name: UpdateMovieIsWatched :exec
UPDATE movies
SET is_watched = ?
WHERE id = ?;
--

-- name: DeleteMovie :exec
DELETE FROM movies WHERE id = ?;
--