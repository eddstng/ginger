# Ginger
A lightweight and easy-to-use point of sale system for restaurants and small businesses.

![Logo](https://i.imgur.com/JNGzVKY.png)

## Environment Variables

To run this project, you will need to add the following environment variables to your `server/.env` file.

`DATABASE_URL`=`postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable`

`TEST_DATABASE_URL`=`postgresql://test:test@localhost:6432/test?sslmode=disable`

`DEFAULT_ITEM_CATEGORY_ID`=`1` 
- Will default to 1 if value is not set.


## Running Tests

#### Ginger API
To run tests for the Ginger API, run the following commands

```bash
$ make test
```

## Running System

#### Ginger API
To start the server for the Ginger API, run the following commands

```bash
$ make server
```

