
# Setup

This project uses direnv, so either install it and run `direnv allow` or do `source .envrc` to
load. Run `createdb pgtypes_development` to create a database to use for code generation.

# Building & Testing

Run:

```
go generate ./...
go test ./...
```
