-- name: AddProject2Favourites :exec
INSERT INTO favourites(user_id, project_id)
VALUES (?, ?);
--

-- name: DeleteProjectFromFavourites :exec
DELETE FROM favourites WHERE user_id = ? AND project_id = ?;
--