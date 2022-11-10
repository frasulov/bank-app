-- name: GetSampleById :one
SELECT * FROM samples
WHERE id = $1 LIMIT 1;

-- name: GetSamples :many
SELECT * FROM samples;
