# Financial Planner BE

## Directory Structure

```
├── cmd -> contain configuration and main file
├── domain -> contain entities object of column as a domain that used by Interactors, Platform, and Transport layer
├── interactors -> Interactors layer that will interact with Repository in Platform layer
│   └── user
├── platform -> Platform layer that will interact with interactors
│   ├── mysql
│   └── yaml
└── transport
    └── http -> contain router also httphandler
```

## List of Package

1. Object Relational Mapping (ORM): [GORM](https://gorm.io/)
2. Web Router: [Gin](https://github.com/gin-gonic/gin)

## How to Run

1. Install Golang first, kindly read [https://golang.org/doc/install](https://golang.org/doc/install)
2. Install all packages required with `go mod download`
3. Create file `config.yaml` on directory `cmd/` based on example format in `cmd/config.example.yaml` and fill the configuration as yours.
4. Launch the file with `go run main.go` 


