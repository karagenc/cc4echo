# Concurrency-Safe Context for Echo

This is a safe to access from multiple goroutines (concurrency-safe) wrapper for [Echo](https://github.com/labstack/echo)'s `Context`. The idea of a custom `Context` for concurrent access might seem redundant at first, but given a deep thought, there are times when you run goroutines inside an HTTP handler. For projects that heavily make use of goroutines, I wrapped `Context` with a mutex to make it safe for concurrent access.

**Note:** This is experimental.

## Usage

`go get -u github.com/karagenc/cc4echo`

### Replace `Context` in place:
```go
func handler(c echo.Context) error {
	c = cc4echo.New(c)
	// Run your goroutines here...
	return nil
}
```

### Or use middleware:
```go
e := echo.New()

// Make sure that cc4echo's wrapper is registered first.
e.Use(cc4echo.Wrapper())

e.Use(middleware.Logger())
e.Use(middleware.Recover())

e.Start("127.0.0.1:3000")
```

## Why?

`Context` has functions that are dangerous to be executed from multiple goroutines. For example have a look at [Request](https://github.com/labstack/echo/blob/0ce73028d0815e0ecec80964cc2da42d98fafa33/context.go#L231) and [SetRequest](https://github.com/labstack/echo/blob/0ce73028d0815e0ecec80964cc2da42d98fafa33/context.go#L235).

Add the route below to your code and compile with race detector enabled: `go build -race`. Then send an HTTP request to it, and see what happens.

```go
func handler(c echo.Context) error {
	go func() {
		for i := 0; i < 100; i++ {
			r := c.Request()
			r, _ = http.NewRequest("GET", "/", nil)
			c.SetRequest(r)
		}
	}()

	for i := 0; i < 100; i++ {
		r := c.Request()
		r, _ = http.NewRequest("GET", "/", nil)
		c.SetRequest(r)
	}
	return nil
}
```

This package is an opt-in solution to this problem.
