---
title: Views
description: Views, how we see your application
layout: ../../../layouts/MainLayout.astro
---

Views in biplane can be used to serve either json or html. (HTML under consideration)

## Defining Your Views

Right now views are sent to the user when they use the application's controllers. These have an underlying structure of data that the user should have access to and may or may not represent the underlying model or be different.


```go
type CounterView struct {
	Count     int    `json:"int"`
	Increment string `json:"increment"`
	Decrement string `json:"decrement"`
}

func (c Counter) View() CounterView {
	return CounterView {
		Count: c.Count,
		Increment: fmt.Sprintf("/api/coutner/%d/increment", c.ID),
		Decrement: fmt.Sprintf("/api/coutner/%d/decrement", c.ID),
	}
}
```
