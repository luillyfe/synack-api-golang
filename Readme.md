# Concurrency in go 

It describe how to make concurrent http request using channels, how to parse json into Structs and custom its fields.

## Installation

Take a look at [go commands](https://golang.org/cmd/go/) for more details.

```bash
go build -o /yourfoldertodeploy .
go run /yourfoldertodeploy/yourexecfile
```

## Concurrenry using channels

```go
go func(ch chan []googleItems) {
    var bing googleResponse

    err := json.Unmarshal(searchOnBing(query), &bing)
    if err != nil {
	log.Fatal(err)
    }

    ch <- bing.Items
}(ch)

go func(ch chan []googleItems) {
    var google googleResponse

    err := json.Unmarshal(searchOnGoogle(query), &google)
    if err != nil {
	log.Fatal(err)
    }

    ch <- bing.Items
}(ch)

responseJSON, _ := json.Marshal(&response{Bing: <-ch, Google: <-ch})
```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[MIT](https://choosealicense.com/licenses/mit/)