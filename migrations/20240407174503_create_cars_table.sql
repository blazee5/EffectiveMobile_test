-- +goose Up
-- +goose StatementBegin
CREATE TABLE cars(
    id SERIAL PRIMARY KEY,
    reg_num VARCHAR(255) NOT NULL,
    mark VARCHAR(255) NOT NULL,
    model VARCHAR(255) NOT NULL,
    year INT NOT NULL,
    owner_name VARCHAR(255) NOT NULL,
    owner_surname VARCHAR(255) NOT NULL,
    owner_patronymic VARCHAR(255) NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE cars;
-- +goose StatementEnd
