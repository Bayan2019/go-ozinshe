-- name: CreateRefreshToken :exec
INSERT INTO refresh_tokens(token, created_at, updated_at, user_id, expires_at, revoked_at)
VALUES (
    ?, 
    NOW(), NOW(), ?, 
    ?, NULL
);
--

-- name: GetUserFromRefreshToken :one
SELECT users.* FROM users
JOIN refresh_tokens ON users.id = refresh_tokens.user_id
WHERE refresh_tokens.token = ?
    AND revoked_at IS NULL
    AND expires_at > NOW()
ORDER BY created_at DESC;
--

-- name: RevokeToken :exec
UPDATE refresh_tokens
SET updated_at = NOW(), revoked_at = NOW()
WHERE token = ? AND revoked_at IS NULL;
--

-- name: GetRefreshTokenOfUser :one
SELECT token FROM refresh_tokens
WHERE user_id = ?
    AND revoked_at IS NULL
    AND expires_at > NOW()
ORDER BY created_at DESC;
--