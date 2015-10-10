# mockapi

*Mockapi* is a simple **Go** application to create a mock server from a given **API Blueprint** standard document specification.

## How to install

First of all, you have to get the source from **github.com** by using the following command:

```bash
got get https://github.com/joliva-ob/mockapi.git
```

And then, just run the program with just two parameters:
+ -f `Path to the apiblueprint file`
+ -p `Server port`

Example:

```bash
go run mockapi -p ./myapi.apib -p 8000
```

## Links
+ [API Blueprint](http://www.apiblueprint.org)
+ [Go language](http://www.golang.org)

## Contact
+ Joan Oliva [joanoj@gmail.com](joanoj@gmail.com)
