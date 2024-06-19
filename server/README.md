# ginger/server
The API server for `ginger/client`. Built using Go and Chi.

![Logo](https://i.imgur.com/JNGzVKY.png)

## Setup Instructions

Go to the project directory

```bash
  cd ginger/server
```

Build docker image

```bash
  docker build -t ginger-server .
```

Run the docker container

```bash
  docker run -p 3000:3000 ginger-server
```

Call the server

```bash
  curl http://localhost:3000/
```
