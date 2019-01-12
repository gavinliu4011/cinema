# Movie Service

This is the Movie service

Generated with

```
micro new cinema/movie --namespace=com.cinema --type=srv
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: com.cinema.srv.movie
- Type: srv
- Alias: movie

## Dependencies

Micro services depend on service discovery. The default is consul.

```
# install consul
brew install consul

# run consul
consul agent -dev
```

## Usage

A Makefile is included for convenience

Build the binary

```
make build
```

Run the service
```
./movie-srv
```

Build a docker image
```
make docker
```