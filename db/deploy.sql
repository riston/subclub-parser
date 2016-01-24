
CREATE SCHEMA IF NOT EXISTS sub;

DROP TABLE IF EXISTS movies;

CREATE TABLE sub.movies (
    movie_id VARCHAR(64) PRIMARY KEY NOT NULL,
    data JSONB NOT NULL
);

CREATE UNIQUE INDEX movies_movie_id_uindex ON sub.movies(movie_id);
