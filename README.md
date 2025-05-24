<h1 align="center">
Http-Catch-All :mag_right:
</h1>

[![test](https://github.com/TeraNovaLP/Http-Catch-All/workflows/Test/badge.svg)](https://github.com/TeraNovaLP/Http-Catch-All/commits/master)

An HTTP server that catches every ingoing request, logs it and returns an 200 OK response.

[Docker Hub](https://hub.docker.com/r/teranovalp/http-catch-all)

## Usage

### Docker

```sh
docker run -t -i --init --rm -p 8080:8080 teranovalp/http-catch-all:latest
```

### Shell

Download the executable file of the latest Github release and run it like this:

```sh
./http-catch-all -port=8080
```

## Arguments

| Option | Description                         |
| ------ | ----------------------------------- |
| -port  | Port to listen on. Defaults to 8080 |
