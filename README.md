# companies

This microservice with a REST API handles information about companies. In my microservice architecture, I typically utilize an API gateway, such as KONG, which takes care of tasks like CORS, rate limiting, and user authentication.

This service uses my [custom golang toolkit](https://github.com/tjarkmeyer/golang-toolkit).

## What you need to develop
- Docker
- Go 1.19

## Development
- Run `make start-dev` to start the development environment (starts local PostgreSQL DB)
- Run `make run-dev` to start the service
- To stop environment run `make stop-dev` (stops local PostgreSQL DB)

### Hints
You need a docker-network called

> dev_network

To create the `dev_network` run:

```bash
docker network create dev_network
```

## Deploy to kubernetes
Use kustomize to deploy:
```bash
kubectl create ns app
kubectl create ns db
kubectl apply -k ./k8s/postgres
kubectl apply -k ./k8s/service
```
