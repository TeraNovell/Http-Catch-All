<h1 align="center">
Http-Catch-All :mag_right:
</h1>

[![test](https://github.com/TeraNovaLP/Http-Catch-All/workflows/Test/badge.svg)](https://github.com/TeraNovaLP/Http-Catch-All/commits/master)

An HTTP server that catches every ingoing request, logs it and returns an 200 OK response.

[Docker Hub](https://hub.docker.com/r/teranovalp/http-catch-all)

## Usage

### Docker

```sh
docker run -t -i --init --rm -p 5000:5000 teranovalp/http-catch-all:latest
```

### Node

Clone the repository and run the following commands in the cloned directory:

```sh
npm i
```

```sh
npm run start
```

## Environment Variables

| VAR  | Description      |
| ---- | ---------------- |
| PORT | Defaults to 5000 |
