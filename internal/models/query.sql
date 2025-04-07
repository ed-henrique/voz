-- name: GetCard :one
select * from cards where id = @id limit 1;

-- name: GetCards :many
select * from cards;

-- name: GetCardsByUser :many
select * from cards where user_id = @user_id;

-- name: GetUpvotesByCard :many
select count(*) from upvotes where card_id = @card_id;

-- name: GetUpvotesByUser :many
select * from upvotes where user_id = @user_id;

-- name: GetDownvotesByCard :many
select count(*) from downvotes where card_id = @card_id;

-- name: GetDownvotesByUser :many
select * from downvotes where user_id = @user_id;

-- name: GetComment :one
select * from comments where id = @id limit 1;

-- name: GetCommentsByCard :many
select * from comments where card_id = @card_id;

-- name: GetCommentsByUser :many
select * from comments where user_id = @user_id;

-- name: GetCommentsByComment :many
select * from comments where comment_id = @comment_id;

-- name: GetUser :one
select * from users where id = @id limit 1;
