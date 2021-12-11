---
title: Models
description: Models, the basic building blocks of your application
layout: ../../../layouts/MainLayout.astro
---

Models are the basic building blocks of your application. They represent state and serve as datastructures in our business logic. Models are often stored to a database using the [objects](/mvc/objects) api.


## Defining Your Models

You can use the model mixin provided or use vanilla go structures for biplane models.

### Using the Model Mixin

The model mixin provided in biplane will give your some methods help with persistance and data parsing. In the following code we are using our models mixin to define `Counter` which has a `Count` field representing some state we want to persist to our objects storage.

```go
import "github.com/ccutch/biplane/mixins"

type Counter struct {
	mixins.Model

	Count int
}
```


### Using Standard Go

You can also just define your models as just a plain vanilla go struct if you do not want a reference back to the object that it is stored under.

```go
type UserAccount struct {
	// We dont need to store an object because we know the ID
	// as we want the user account id to match the id for the
	// user given to use from the auth manager.
	UserID int

	Avatar string
	Name string
}
```
