# Ticketing Service

This is the Ticketing service

Generated with

```
micro new cinema/ticketing --namespace=com.cinema --type=srv
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: com.cinema.srv.ticketing
- Type: srv
- Alias: ticketing

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
./ticketing-srv
```

Build a docker image
```
make docker
```