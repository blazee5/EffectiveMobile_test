-- +goose Up
-- +goose StatementBegin
CREATE TABLE cars(
    id SERIAL PRIMARY KEY,
    reg_num VARCHAR(255) NOT NULL DEFAULT '',
    mark VARCHAR(255) NOT NULL DEFAULT '',
    model VARCHAR(255) NOT NULL DEFAULT '',
    year INT NOT NULL DEFAULT 0,
    owner_name VARCHAR(255) NOT NULL DEFAULT '',
    owner_surname VARCHAR(255) NOT NULL DEFAULT '',
    owner_patronymic VARCHAR(255) NOT NULL DEFAULT ''
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE cars;
-- +goose StatementEnd
