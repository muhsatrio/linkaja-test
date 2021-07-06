# Financial Planner BE

## Directory Structure

```
├── cmd -> contain configuration and main file
├── domain -> contain entities object of column as a domain that used by Interactors, Platform, and Transport layer
├── interactors -> Interactors layer that will interact with Repository in Platform layer
│   └── user -> contain service function and usecase that interaction with platform repository function
├── platform -> Platform layer that will interact with interactors
│   ├── mysql -> contain init function, model and repository function
│   └── yaml -> contain config and init function
└── transport
    └── http -> contain routes, http handler, and error handler 
```

## List of Package

1. Object Relational Mapping (ORM): [GORM](https://gorm.io/)
2. Web Router: [Gin](https://github.com/gin-gonic/gin)

## How to Run

1. Install Golang first, kindly read [https://golang.org/doc/install](https://golang.org/doc/install)
2. Install all packages required with `go mod download`
3. Create file `config.yaml` on directory `cmd/` based on example format in `cmd/config.example.yaml` and fill the configuration as yours.
4. Launch the file with `go run main.go` 

## Todo

- Add login feature
- Add mocks
- Add Dockerfile

## Reference

[Ready for changes with Hexagonal Architecture](https://netflixtechblog.com/ready-for-changes-with-hexagonal-architecture-b315ec967749)
 by [Netflix](https://www.netflix.com/)

