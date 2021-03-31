-- +goose Up
INSERT INTO
    pokemon.cards (id, name)
VALUES
    (1, 'Pikachu'),
    (2, 'Bulbasaur'),
    (3, 'Charmander'),
    (4, 'Squirtle');