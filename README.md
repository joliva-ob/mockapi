# mockapi

*Mockapi* is a simple **Go** application to create a mock server from a given **API Blueprint** standard document specification.

## How to install and run

First of all, you have to get the source from **github.com** by using the following command:

```bash
go get https://github.com/joliva-ob/mockapi.git
```

And then, just run the program with just two parameters:
+ -f `Path to the apiblueprint file`
+ -p `Server port`

Example:

```bash
go run mockapi -p ./myapi.apib -p 8000
```

or run in background with output redirected to a log file.

```bash
go run mockapi.go -f ./myapi.apib -p 8000 > mockapi.log
```

## Links
+ [API Blueprint](http://www.apiblueprint.org)
+ [Go language](http://www.golang.org)
+ [Go mux: request router and dispatcher](https://github.com/gorilla/mux)

## Contact
+ Joan Oliva [joanoj@gmail.com](joanoj@gmail.com)
