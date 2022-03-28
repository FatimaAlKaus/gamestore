CREATE TABLE IF NOT EXISTS genres
(
    id   serial primary key,
    name varchar(30) unique check ( name <> '' ) not null
);
CREATE TABLE IF NOT EXISTS rating
(
    id   serial primary key,
    name varchar(30) unique check ( name <> '' )
);
CREATE TABLE IF NOT EXISTS games
(
    id     serial primary key,
    name   varchar(30) check ( name <> '' ) not null,
    price  decimal check ( price > 0 )      not null,
    rating int references rating (id)       not null
);
CREATE TABLE IF NOT EXISTS genres_games
(
    genre_id int references genres (id) on update cascade on delete cascade,
    game_id  int references games (id) on update cascade on delete cascade,
    constraint genre_game_key primary key (genre_id, game_id)
);
