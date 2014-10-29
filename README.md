# pgweb

Web-based Mysql database browser written in Go. 
This is the fork of [pgweb](https://github.com/sosedoff/pgweb) project, 
adapted for Mysql database.

Many thanks to Dan for the original Postgres version.

## Overview

This is a web-based browser for Mysql database server. Its written in Go
and works on Mac OSX, Linux and Windows machines. Main idea behind using Go for the backend
is to utilize language's ability for cross-compile source code for multiple platforms. 
This project is an attempt to create a very simple and portable application to work with 
PostgreSQL databases.

<img src="screenshots/browse.png" width="345px" />
<img src="screenshots/query.png" width="345px" />

Features:

- Connect to local or remote server
- Browse tables and table rows
- Get table details: structure, size, indices, row count
- Execute SQL query and run analyze on it
- Export query results to CSV
- View query history

## Installation

Go 1.3+ is required. You can install Go with `homebrew`:

```
brew install go
```

To compile source code run the following command:

```
make setup
make dev
```

This will produce `pgweb` binary in the current directory.

There's also a task to compile binaries for other operating systems:

```
make build
```

Under the hood it uses [gox](https://github.com/mitchellh/gox). Compiled binaries
will be stored into `./bin` directory.

## Usage

Start server:

```
pgweb --host localhost --user myuser --db mydb
```

You can also specify a connection URI instead of individual flags:

```
pgweb --url user:password@tcp(host:port)/database
```

Full CLI options:

```
Usage:
  pgweb [OPTIONS]

Application Options:
  -v, --version  Print version
  -d, --debug    Enable debugging mode (false)
      --url=     Database connection string
      --host=    Server hostname or IP (localhost)
      --port=    Server port (5432)
      --user=    Database user (mysql)
      --pass=    Password for user
      --db=      Database name (mysql)
      --listen=  HTTP server listen port (8080)
```

## Contributors

- Dan Sosedoff - https://twitter.com/sosedoff
- Masha Safina - https://twitter.com/mashasafina
- Gennadiy Kovalev

## License

The MIT License (MIT)

Copyright (c) 2014 Dan Sosedoff, <dan.sosedoff@gmail.com>
Changed by Gennadiy Kovalev, <graarh@weird.company>
