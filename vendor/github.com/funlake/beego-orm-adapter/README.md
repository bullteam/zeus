Beego ORM Adapter [![Build Status](https://travis-ci.org/casbin/beego-orm-adapter.svg?branch=master)](https://travis-ci.org/casbin/beego-orm-adapter) [![Coverage Status](https://coveralls.io/repos/github/casbin/beego-orm-adapter/badge.svg?branch=master)](https://coveralls.io/github/casbin/beego-orm-adapter?branch=master) [![Godoc](https://godoc.org/github.com/casbin/beego-orm-adapter?status.svg)](https://godoc.org/github.com/casbin/beego-orm-adapter)
====

Beego ORM Adapter is the [Beego ORM](https://beego.me/docs/mvc/model/overview.md) adapter for [Casbin](https://github.com/casbin/casbin). With this library, Casbin can load policy from Beego ORM supported database or save policy to it.

Based on [Beego ORM Support](https://beego.me/docs/mvc/model/overview.md), The current supported databases are:

- MySQL: [github.com/go-sql-driver/mysql](https://github.com/go-sql-driver/mysql)
- PostgreSQL: [github.com/lib/pq](https://github.com/lib/pq)
- Sqlite3: [github.com/mattn/go-sqlite3](https://github.com/mattn/go-sqlite3)

## Installation

    go get github.com/casbin/beego-orm-adapter

## Simple MySQL Example

```go
package main

import (
	"github.com/casbin/beego-orm-adapter"
	"github.com/casbin/casbin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Initialize a Beego ORM adapter and use it in a Casbin enforcer:
	// The adapter will use the MySQL database named "casbin".
	// If it doesn't exist, the adapter will create it automatically.
	a := beegoormadapter.NewAdapter("mysql", "mysql_username:mysql_password@tcp(127.0.0.1:3306)/") // Your driver and data source. 

	// Or you can use an existing DB "abc" like this:
	// The adapter will use the table named "casbin_rule".
	// If it doesn't exist, the adapter will create it automatically.
	// a := beegoormadapter.NewAdapter("mysql", "mysql_username:mysql_password@tcp(127.0.0.1:3306)/abc", true)

	e := casbin.NewEnforcer("examples/rbac_model.conf", a)
	
	// Load the policy from DB.
	e.LoadPolicy()
	
	// Check the permission.
	e.Enforce("alice", "data1", "read")
	
	// Modify the policy.
	// e.AddPolicy(...)
	// e.RemovePolicy(...)
	
	// Save the policy back to DB.
	e.SavePolicy()
}
```

## Simple Postgres Example

```go
package main

import (
	"github.com/casbin/beego-orm-adapter"
	"github.com/casbin/casbin"
	_ "github.com/lib/pq"
)

func main() {
	// Initialize a Beego ORM adapter and use it in a Casbin enforcer:
	// The adapter will use the Postgres database named "casbin".
	// If it doesn't exist, the adapter will create it automatically.
	a := beegoormadapter.NewAdapter("postgres", "user=postgres_username password=postgres_password host=127.0.0.1 port=5432 sslmode=disable") // Your driver and data source.

	// Or you can use an existing DB "abc" like this:
	// The adapter will use the table named "casbin_rule".
	// If it doesn't exist, the adapter will create it automatically.
	// a := beegoormadapter.NewAdapter("postgres", "dbname=abc user=postgres_username password=postgres_password host=127.0.0.1 port=5432 sslmode=disable", true)

	e := casbin.NewEnforcer("../examples/rbac_model.conf", a)

	// Load the policy from DB.
	e.LoadPolicy()

	// Check the permission.
	e.Enforce("alice", "data1", "read")

	// Modify the policy.
	// e.AddPolicy(...)
	// e.RemovePolicy(...)

	// Save the policy back to DB.
	e.SavePolicy()
}
```

## Getting Help

- [Casbin](https://github.com/casbin/casbin)

## License

This project is under Apache 2.0 License. See the [LICENSE](LICENSE) file for the full license text.
