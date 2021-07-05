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


