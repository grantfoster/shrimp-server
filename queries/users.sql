-- name: createUser :one
insert into users (
  email
) values (
  $1
) returning id;

-- name: getUserPassword :one
select password from users
where username = $1;
