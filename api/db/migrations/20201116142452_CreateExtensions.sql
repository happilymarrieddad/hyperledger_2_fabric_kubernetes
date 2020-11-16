-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE SCHEMA IF NOT EXISTS tiger;
CREATE SCHEMA IF NOT EXISTS tiger_data;
CREATE SCHEMA IF NOT EXISTS topology;
CREATE EXTENSION IF NOT EXISTS fuzzystrmatch;
CREATE EXTENSION IF NOT EXISTS postgis_tiger_geocoder;
CREATE EXTENSION IF NOT EXISTS address_standardizer;

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
