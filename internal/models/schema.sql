-- User Types
create table user_types (
    id integer primary key,
    name text not null,
    created_at datetime not null default(strftime('%Y-%m-%dT%H:%M:%f', 'now')),
    updated_at datetime,
    removed_at datetime
);

insert into user_types (name) values ('Common');

-- Users
create table users (
    id integer primary key,
    name text not null,
    email text unique not null,
    username text unique not null,
    password text not null,
    user_type_id integer not null,
    created_at datetime not null default(strftime('%Y-%m-%dT%H:%M:%f', 'now')),
    updated_at datetime,
    removed_at datetime,
    foreign key (user_type_id) references user_types(id)
);

create unique index uidx_user on users(email, username, user_type_id);

-- Cards
create table cards (
    id integer primary key,
    name text not null,
    description text not null,
    user_id integer not null, -- Author
    created_at datetime not null default(strftime('%Y-%m-%dT%H:%M:%f', 'now')),
    updated_at datetime,
    removed_at datetime,
    foreign key (user_id) references users(id)
);

create index idx_cards_by_user on cards(user_id);

-- Upvotes
create table upvotes (
    id integer primary key,
    card_id integer not null,
    user_id integer not null,
    created_at datetime not null default(strftime('%Y-%m-%dT%H:%M:%f', 'now')),
    updated_at datetime,
    removed_at datetime,
    foreign key (card_id) references cards(id),
    foreign key (user_id) references users(id)
);

create index idx_upvotes_by_card on upvotes(card_id);
create index idx_upvotes_by_user on upvotes(user_id);
create unique index uidx_upvotes on upvotes(card_id, user_id);

-- Downvotes
create table downvotes (
    id integer primary key,
    card_id integer not null,
    user_id integer not null,
    created_at datetime not null default(strftime('%Y-%m-%dT%H:%M:%f', 'now')),
    updated_at datetime,
    removed_at datetime,
    foreign key (card_id) references cards(id),
    foreign key (user_id) references users(id)
);

create index idx_downvotes_by_card on downvotes(card_id);
create index idx_downvotes_by_user on downvotes(user_id);
create unique index uidx_downvotes on downvotes(card_id, user_id);

-- Comments
create table comments (
    id integer primary key,
    content text not null,
    card_id integer not null,
    user_id integer not null,
    comment_id integer,
    created_at datetime not null default(strftime('%Y-%m-%dT%H:%M:%f', 'now')),
    updated_at datetime,
    removed_at datetime,
    foreign key (card_id) references cards(id),
    foreign key (user_id) references users(id),
    foreign key (comment_id) references comments(id)
);

create index idx_comments_by_card on comments(card_id);
create index idx_comments_by_user on comments(user_id);
create index idx_comments_by_comment on comments(comment_id) where comment_id is not null;
