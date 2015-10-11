# revel-gorm

[GORM](https://github.com/jinzhu/gorm) integration module for [Revel](https://revel.github.io) framework.

# Getting Started

## Install

    % go get github.com/Gr1N/revel-gorm

## Documentation

[![GoDoc](https://godoc.org/github.com/Gr1N/revel-gorm?status.svg)](https://godoc.org/github.com/Gr1N/revel-gorm)

`go doc` format documentation for this project can be viewed online without
installing the package by using the GoDoc page at:
http://godoc.org/github.com/Gr1N/revel-gorm

## Configuration

Settings can be configured via the following directives in app.conf.

### module.gorm

    module.gorm = github.com/Gr1N/revel-gorm

Please review the documentation at [Revel — Modules — Overview](http://revel.github.io/manual/modules.html) for more information.

### gorm.driver

    gorm.driver = postgres

Specifies the name of the database/sql driver.

### gorm.spec

    gorm.spec = "user=username dbname=databasename sslmode=disable"

Specifies the data source name of your database/sql database.

### gorm.max_idle_conns

    gorm.max_idle_conns = 10

### gorm.max_open_conns

    gorm.max_open_conns = 100

### gorm.singular_table

    gorm.singular_table = false

Please review the documentation at [GORM — Conventions](https://github.com/jinzhu/gorm#conventions).

### gorm.log_mode

    gorm.log_mode = false

Please review the documentation at [GORM — Logger](https://github.com/jinzhu/gorm#logger).

## Initialization

Add following code to your `app/init.go`:

    package app

    import (
        ...
        "github.com/revel/revel"
        gorm "github.com/Gr1N/revel-gorm/app"
        ...
    )

    func init() {
        ...
        revel.OnAppStart(func() {
            // Initialize GORM...
            dbm := gorm.InitDB()
            // ...and migrate
            dbm.AutoMigrate(&User{})
        })
        ...
    }

And to your `app/controllers/init.go`:

    package controllers

    import (
        "github.com/revel/revel"
        gorm "github.com/Gr1N/revel-gorm/app/controllers"
    )

    func init() {
        revel.InterceptMethod((*gorm.TransactionalController).Begin, revel.BEFORE)
        revel.InterceptMethod((*gorm.TransactionalController).Commit, revel.AFTER)
        revel.InterceptMethod((*gorm.TransactionalController).Rollback, revel.FINALLY)
    }

And embed the `gorm.TransactionalController` on your custom controller:

    package controllers

    import (
        gorm "github.com/Gr1N/revel-gorm/app/controllers"
    )

    type Application struct {
        gorm.TransactionalController
    }

# TODO

- [x] Documentation
— [ ] Sample application
- [ ] Tests
- [ ] ...

# License

*revel-gorm* is licensed under the MIT license. See the license file for details.
