# k8s Scanner Collector

Script (deployable to k8s) to log requests from vulnerability and exploit scanners.

## Helpful Commands

### Build the Docker Container

To build the container with the name `scanner-collector`, run:

```
docker build -t scanner-collector .
```

### Run the Docker Container

To run the docker container locally, run:

```
docker run -p 80:80 scanner-collector
```

and then the collector will be available locally at http://localhost:80.

### Exec Commands in the Docker Container

You can run commands in a running docker container using:

```
docker exec -it <CONTAINER_NAME> <COMMAND>
```

for example, this command drops you into a shell in the docker container:

```
docker exec -it <CONTAINER_NAME> sh
```

Note that this only works on a *running* docker container.

