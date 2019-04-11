[![Build Status](https://travis-ci.com/peterfox/waitformysql.svg?token=sykMyWNtTgYgm3dx9GKK&branch=master)](https://travis-ci.com/peterfox/waitformysql)

# WaitForMySQL

A simple command-line tool for blocking scripts from running
until a connection is established or a timeout is reached.

This tool was designed to be used in build scripts for things
like Travis CI where I was often running into problems with the MySQL
docker container command completing before the new database was ready.
This meant further steps failed and I kept having to add more time to
the build by using the sleep command. Eventually I decided to make this.
I now use it to block the build process until the database is ready
or an appropriate amount of time has passed.

## Usage

CLI usage is pretty simple as set out below. You only need specify
the password currently.

```
  -database string
        database name
  -host string
        database host (default "127.0.0.1")
  -password string
        database password
  -port string
        database port (default "3306")
  -timeout duration
        max timeout in seconds (default 25s)
  -username string
        database username (default "root")
```

## Building

A make file is included. You should be able to simply run `make` to
build the binaries for macOS, Linux and Windows.