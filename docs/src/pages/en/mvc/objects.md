---
title: Objects
description: Objects, storage for your models
layout: ../../../layouts/MainLayout.astro
---

## Creating Objects

There are two ways we are given to create objects in biplane. First, we can use the Object API `New` method which will create an object for the model's data and generate a new ID. Second, we can use the Object API `Build` method which will create a new Object with a given id that we can later `Insert`.

### Create a New Object

Most of the time we will want to use the `New` method.

```go 
func CreateCounter(
	user auth.User, conf database.Config,
) (*Counter, error) {

	c := Counter{Count: 0}
	_, err := objects.NewClient(conf).
		New(user, "Counter", &c)

	return &c, err
}
```

### Build a Object

Sometimes, when we know the ID of the object before we have created it, we can use the `Build` function with a call to the `Insert` method.

```go
func CreateUserAccount(
	user auth.User, conf database.Config,
	avatar, name string, 
) (*UserAccount, error) {

	a := UserAccount{user.ID, avatar, name}
	err := objects.NewClient(conf).
		Build(a.ID, user, "UserAccount", &a).
		Insert()

	return &a, err
}
```