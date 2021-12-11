---
title: Controllers
description: Controllers, what your application does
layout: ../../../layouts/MainLayout.astro
---


Controllers are where we store most of our logic, provide route information, and allow us to parse and authenticate requests from our user.

## Defining Your Controllers

It is strongly recommended that you use this `Controller` mixin provided in the `mixins` package. This will give functionality for user management. If you are making a controller that does not need this you can always use vanilla go structs that implement the `server.Router` interface

### Using the Controller Mixin

The mixin is strongly recommened, by default a `Routes` method is provided so we can use this right away.

```go 
import "github.com/ccutch/biplane"

type MyAppController struct {
	mixins.Controller
}
```

### Using Standard Go

```go
type MyAppController struct {}

func (c MyAppController) Routes(r *mux.Router) {
	r.Path("/").HandlerFunc(c.HelloWorld)
}

func (c MyAppController) HelloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world"))
}
```

## Modules

The Controller mixing comes with functions that help expose objects and auth apis.

### Objects

We can access the Objects api from a controller using the `Objects` method. This will use the database `Driver` we defined in our server application.

```go
import "github.com/ccutch/biplane/mixins"

type MyAppController struct {
	mixins.Controller
}

func (m MyAppController) Handler(w http.ResponseWriter, r *http.Request) {
	var c Counter
	_, err := m.Objects().Get(1, &c)

	if err != nil {
		m.Fail(err)
		return
	}

	m.Display(c.View())
}
```

### Auth Manager

We can use the auth manager to get the user that is calling this route

```go
import "github.com/ccutch/biplane/mixins"

type MyAppController struct {
	mixins.Controller
}

func (c MyAppController) Handler(w http.ResponseWriter, r *http.Request) {
	c.RequireUser(r)

	a := UserAccount{
		c.User(r).ID,
		r.FormValue("avatar"),
		r.FormValue("name"),
	}

	err := c.Objects().Build(a.ID, user, "UserAccount", &a).Insert()
	if err != nil {
		c.Fail(err)
		return
	}

	c.Display(a.View())
}
```