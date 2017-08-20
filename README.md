# Prudle
#### You can mock any HTTP responses to test RESTful API

Prudle - a simple application which allows to generate custom HTTP responses.
If you want to mock any back-end response in your application, prudle can help you.

## Installation

#### Pre-requisite: You must have installed a latest version of [Go](https://golang.org/doc/install).

Run the below command

 ```
    go get -u github.com/kiritym/prudle
 ```

 This will download prudle to `$GOPATH/src/github.com/kiritym/prudle`.
 From this directory run `go build` to create the `prudle` binary.

 ## Usage

Start the server by executing `prudle` binary. By default, server will listen to http://0.0.0.0:8888.

```
prudle -h
Usage of prudle:
  -b string
        Host of server (default "0.0.0.0")
  -p int
        Port of server (default 8888)
```


## Demo

Create your API endpoints with custom HTTP Response:

![](https://github.com/kiritym/prudle/blob/master/screenshots/prudle.gif)

You can Curl your endpoints to mock the required response:

![](https://github.com/kiritym/prudle/blob/master/screenshots/prudle-curl.gif)
