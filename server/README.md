# ginger/server
The API server for `ginger/client`. Built using Go and Chi.

![Logo](https://i.imgur.com/JNGzVKY.png)


## Setup Instructions

Go to the project directory

```bash
$ cd ginger/server
```

Build docker image

```bash
$ docker build -t ginger-server .
```

Run the docker container

```bash
$ docker run -p 3000:3000 ginger-server
```

Call the server

```bash
$ curl http://localhost:3000/
```

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
$ go test ./... -v -cover
```

