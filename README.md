# WaitForMySQL

A simple command-line tool for blocking scripts from running
until a connection is established or a timeout is reached.

## Usage

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