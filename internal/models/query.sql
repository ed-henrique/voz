-- name: GetCard :one
select * from cards where id = @id limit 1;

-- name: GetCards :many
select * from cards;

-- name: GetCardsByUser :many
select * from cards where user_id = @user_id;

-- name: InsertCard :one
insert into cards (
    name,
    description,
    user_id
) values (
    @name,
    @description,
    @user_id
) returning id;

-- name: GetUpvotesByCard :many
select count(*) from upvotes where card_id = @card_id;

-- name: GetUpvotesByUser :many
select * from upvotes where user_id = @user_id;

-- name: UpvoteCard :one
insert into upvotes (
    card_id,
    user_id
) values (
    @card_id,
    @user_id
) returning id;

-- name: GetDownvotesByCard :many
select count(*) from downvotes where card_id = @card_id;

-- name: GetDownvotesByUser :many
select * from downvotes where user_id = @user_id;

-- name: DownvoteCard :one
insert into downvotes (
    card_id,
    user_id
) values (
    @card_id,
    @user_id
) returning id;

-- name: GetComment :one
select * from comments where id = @id limit 1;

-- name: GetCommentsByCard :many
select * from comments where card_id = @card_id;

-- name: GetCommentsByUser :many
select * from comments where user_id = @user_id;

-- name: GetCommentsByComment :many
select * from comments where comment_id = @comment_id;

-- name: InsertComment :one
insert into comments (
    content,
    card_id,
    user_id,
    comment_id
) values (
    @content,
    @card_id,
    @user_id,
    CASE
        WHEN @comment_id = 0 THEN NULL
        ELSE @comment_id
    END
) returning id;

-- name: GetUserTypes :many
select * from user_types;

-- name: GetUser :one
select * from users where id = @id limit 1;

-- name: InsertUser :one
insert into users (
    name,
    email,
    username,
    password,
    user_type_id
) values (
    @name,
    @email,
    @username,
    @password,
    @user_type_id
) returning id;
