# key-count
A simple golang service that counts keys in redis

## Usage

Set redis address

```bash
export REDIS_ADDRESS="hostname:6379"
```

Set redis password

```bash
export REDIS_PASSWORD="REDIS_PASSWORD_HERE"
```

Run the application

```bash
make run
```

Build the project

```bash
make build
```

Build the docker image and push it

```bash
make build-docker-image
```