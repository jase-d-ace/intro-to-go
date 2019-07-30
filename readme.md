# Intro to Golang

This repo holds all of my practice golang code. Each folder holds a `.go` file that treats a different part of the fundamentals of the language. 

`hello/hello.go` is the very first thing I tried, which follows the official golang docs tutorial on how to get started. It holds a `main.go` function that calls two other files that I wrote.

`stringutil/reverse.go` is also part of the official tutorial, which covers one loop and a string manipulation.

`httpreqs/req.go` is my first attempt at hitting a JSON endpoint and manipulating data to be served in JSON format.

`server/server.go` is a continuation of my work in `httpreqs/req.go`, only this time it's meant to be a standalone JSON server, with a couple of example routes.

## Next Steps

As of July 30, 2019, my next steps are to continue learning about the basic data structures and datatypes of golang, and hopefully work them into `server/server.go` for a more robust and fully-functioning server.

I'd also like to learn more about how files are structured in golang server architecture so that I can be more aware of how to structure my own code