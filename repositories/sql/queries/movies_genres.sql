-- name: AddGenre2Movie :exec
INSERT INTO movies_genres(movie_id, genre_id)
VALUES (?, ?);
--

-- name: DeleteGenresOfMovie :exec
DELETE FROM movies_genres WHERE movie_id = ?;
--