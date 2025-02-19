-- name: AddMovie2Watchlist :exec
INSERT INTO watchlist(movie_id, added_at)
VALUES (?, NOW());

-- name: GetWatchlistMovies :many
SELECT movie_id
FROM watchlist
ORDER BY added_at;

-- name: DeleteMovieFromWatchlist :exec
DELETE FROM watchlist WHERE movie_id = ?;