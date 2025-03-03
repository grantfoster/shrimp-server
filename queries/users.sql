-- name: createUser :one
insert into users (
  email
) values (
  $1
) returning id;
