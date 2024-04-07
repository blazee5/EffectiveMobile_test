migrate-up:
	goose -dir ./migrations postgres "host=localhost user=postgres password=password port=5432 dbname=cars sslmode=disable" up