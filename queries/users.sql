-- name: createUser :one
insert into users (
  email
) values (
  $1
) returning id;

-- name: ListUsers :many
SELECT * FROM donkeys;
