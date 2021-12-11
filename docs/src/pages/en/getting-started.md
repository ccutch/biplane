---
title: Getting Started
description:  Getting started with Biplane
layout: ../../layouts/MainLayout.astro
---

Welcome to Biplane's docs, an MVC framework for go. 

## Give it a spin

We are built ontop of go so installation is easy first make sure you have go installed locally. Then you can run the following command to install biplane.

```bash
go get github.com/ccutch/biplane
```


### Prefabs

Prefabs are a quick way to get started with Biplane. These are prebuilt component that solve common  problems like auth and getting started. Copy the following code into a `main.go` file.

```go
package main

import (
	"github.com/ccutch/biplane"
	"github.com/ccutch/biplane/database"
	"github.com/ccutch/biplane/prefab"
)

func main() {
	biplane.NewServer("", 8080).

		// Using postgres database
		Database(database.Postgres{
			Host:     "localhost",
			Port:     5432,
			User:     "admin",
			Password: "password",
			DBName:   "defaultdb",
		}).

		// Prefab Controllers
		Controller(new(prefab.AuthController)).
		Controller(new(prefab.WelcomeController)).

		// Start Server
		TakeOff()
}
```

### Running Server

Run this new server with the following command.

```bash
go run main.go
```