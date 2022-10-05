# API Server Template
<hr>

This template is inspired by https://ldej.nl/post/enterprise-go-architecture/

## Language, Libraries, and Dependencies
|       Purpose        |                         Name                         | Version | License |
|:--------------------:|:----------------------------------------------------:|:-------:|:-------:|
|       Language       |          [Go](https://github.com/golang/go)          | 1.18.6  |  BSD-3  |
|      Framework       |       [Gin](https://github.com/gin-gonic/gin)        |  1.8.1  |   MIT   |
| Configuration Access |       [Viper](https://github.com/spf13/viper)        | 1.13.0  |   MIT   |
|         Mock         | [go-sqlmock](https://github.com/DATA-DOG/go-sqlmock) |  1.5.0  |  BSD-3  |
|       Logging        |     [logrus](https://github.com/sirupsen/logrus)     |  1.9.0  |   MIT   |
|         ORM          |       [GORM](https://github.com/go-gorm/gorm)        | 1.23.10 |   MIT   |
|  API Documentation   |        [Swag](https://github.com/swaggo/swag)        |  1.8.6  |   MIT   |
|    Message Queue     |      [Go NSQ](https://github.com/nsqio/go-nsq/)      |  1.1.0  |   MIT   |

## Folder Structure
```markdown
├── server
    ├── clients
        ├── microservice-2
    ├── constants
    ├── consumers
    ├── controller
    ├── db
    ├── error
    ├── mocks
    ├── models
        ├── request
        ├── response
    ├── repository
    ├── router
    ├── service
    ├── utils
├── build-scripts
├── config
├── database        // only migrations
├── Dockerfile
├── Dockerfile.debug
├── docs            // swagger
├── infrastructure  // docker
├── go.mod
├── go.sum
├── main.go
├── Makefile
├── README.md
├── sqlite.db
└── template.toml
```

## Steps
1. Copy `template.toml` to `server.toml`
2. Update `server.toml` values
3. [OPTIONAL] Replace `web-server` in all files to server name
