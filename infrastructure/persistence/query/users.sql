-- name: PersistUser :exec
INSERT INTO "users" (
    id, discord_id, pseudo, tag
) VALUES (
    $1, $2, $3, $4
);

-- name: FindUserByDiscordID :one
SELECT * FROM "users"
WHERE discord_id = $1 LIMIT 1;
