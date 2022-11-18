# Go API Boilerplate

## Overview

Golang reusable API boilerplate with:

- [x] [Gin](https://github.com/gin-gonic/gin)
- [x] [GORM](https://gorm.io)
- [x] Manual migrations
- [x] Seeders
- [x] JWT Authentication
- [x] Account Registration
- [x] Encrypted passwords
- [x] Docker

## Roadmap

- [ ] Add tests
- [ ] Swagger docs
- [ ] Telemetry

## Dependencies

- Go 1.19
- Docker
- Docker Compose

## Setup

Copy the file `.env.example` to `.env` and fill in the values.

See `config/config.go` for more details.

## Running

```bash
make up
```

## Commands

### Build

```bash
make build
```

### Migrate and seed

```bash
./api db setup
```

**OR**

```bash
./api db migrate
./api db seed
```

**PS:** Migrations are automatically run on `serve` command, but seeding is not.

### Serve

```bash
./api serve
```

## Credits

I, Elias Feijó, had to create an API for a personal project from scratch, and decided to make it reusable.

I was inspired by some of the code in [bancodobrasil](https://github.com/bancodobrasil) repositories, such as [featws-api](https://github.com/bancodobrasil/featws-api) and [goauth](https://github.com/bancodobrasil/goauth), some of those I've contributed and even was the solo maintainer. Changed a lot of things, but I'm still grateful for the inspiration.

Hope you enjoy it!
