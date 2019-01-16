# projection Service

This is the projection service

Generated with

```
micro new cinema/projection --namespace=com.cinema --type=srv
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: com.cinema.srv.projection
- Type: srv
- Alias: projection

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
./projection-srv
```

Build a docker image
```
make docker
```