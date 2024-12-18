# StocksBack

## What is stocks back?

an api for stocks usage

### User functionality

- Creating accounts `/signup`
- Sign in using password or secret key (in the future jwt) 
- Farming (getting solids based on the stock amount) every hour `/farm`
- Getting solids from stocks every day at 21 (server time)
- Buying stocks `/buy`
- Changing name and password `/change/name`, `/change/password`
- Getting user `/get`

### Backend functionality 

- Two database types, that could be easy changed
    1. [using file system (for testing)](/pkg/file_db/main.go)
    2. [PostgreSQL database](/pkg/db/main.go)
- [Graceful shutdown](/pkg/closer/main.go)
- [Custom cron usage](/pkg/cron/main.go)
- [Hasher](/pkg/hash/hash.go)
- [Custom logger](/pkg/logger/main.go)
- [Specific query expressions that can be used in both database types](/pkg/query/query.go)
- [User service for all user activities](/pkg/user_service/main.go)
- [Http handler and server](/http/server/)
- Good structured headers, requests, responses
    1. [Headers](/http/api/input/headers/headers.go)
    2. [Requests](/http/api/input/requests/requests.go)
    3. [Responses](/http/api/responses/responses.go)
- Timeout server and service mode

## Setup program 

### Getting project

```bash
git clone https://github.com/VandiKond/StocksBack.git
```

### Config setup

[Config example](/config/config.yaml)

Replace it with your data

### Running program

```bash
go run cmd/main.go
```

## License 

[LICENSE](LICENSE)


