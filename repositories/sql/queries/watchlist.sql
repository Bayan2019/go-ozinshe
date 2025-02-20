-- name: AddProject2Watchlist :exec
INSERT INTO watchlist(user_id, project_id)
VALUES (?, ?);
--

-- name: DeleteProjectFromWatchlist :exec
DELETE FROM watchlist WHERE user_id = ? AND project_id = ?;
--