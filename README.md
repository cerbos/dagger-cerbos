# Dagger Cerbos Module

A [Dagger](https://dagger.io) module for using Cerbos in CI environments. This module contains a `Compile` Dagger function for [compiling and testing Cerbos policies](https://docs.cerbos.dev/cerbos/latest/policies/compile) and a `Server` service for starting a Cerbos server.

```sh
# View usage information
dagger -m github.com/cerbos/dagger-cerbos call compile --help
dagger -m github.com/cerbos/dagger-cerbos call server --help

# Compile and run tests on a Cerbos policy repository
dagger -m github.com/cerbos/dagger-cerbos call compile --policy-dir=./cerbos

# Start a Cerbos server with the default disk driver
dagger -m github.com/cerbos/dagger-cerbos call server --policy-dir=./cerbos up

# Start a Cerbos server instance configured to use an in-memory SQLite policy repository
dagger -m github.com/cerbos/dagger-cerbos call server --config=storage.driver=sqlite3,storage.sqlite3.dsn=:memory:,server.adminAPI.enabled=true up
```
