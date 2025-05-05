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

k3d:
```
$ kompose convert -f docker-compose.yaml
$ k3d cluster create ginger (k3d cluster stop ginger, k3d cluster delete my-cluster)
$ k3d image import ginger-api:latest -c ginger
$ kubectl apply -f .
$ kubectl get pods 
$ kubectl port-forward pod/ginger-api-7fd84dd69d-4f4kt 3000:3000
```


```
kubectl apply -f https://github.com/kubernetes-sigs/metrics-server/releases/latest/download/components.yaml

