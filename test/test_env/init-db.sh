#!/bin/bash
set -e
psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
    CREATE TABLE IF NOT EXISTS TEST_PLAIN_ENTITY_TABLE (
        ID SERIAL PRIMARY KEY,
        FIELD1 INT,
        FIELD2 TEXT
    );
    CREATE TABLE IF NOT EXISTS TEST_PARENT_ENTITY_TABLE (
        ID SERIAL PRIMARY KEY,
        NAME TEXT
    );
    CREATE TABLE IF NOT EXISTS TEST_CHILD1_TABLE (
        ID SERIAL PRIMARY KEY,
        TYPE TEXT,
        PARENT_ID INT,
        FOREIGN KEY (PARENT_ID) REFERENCES TEST_PARENT_ENTITY_TABLE (id)
    );

    CREATE TABLE IF NOT EXISTS TEST_CHILD2_TABLE (
        ID SERIAL PRIMARY KEY,
        SIZE float,
        PARENT_ID INT,
        FOREIGN KEY (PARENT_ID) REFERENCES TEST_PARENT_ENTITY_TABLE (id)
    );

EOSQL
