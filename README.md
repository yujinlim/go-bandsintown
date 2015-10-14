# go-bandsintown
> bandsintown golang api package

## Usage
```
go get github.com/yujinlim/go-bandsintown
```

## Example
```go
package main

import "github.com/yujinlim/go-bandsintown"

func main() {
  client := bands.New("apiKey")

  artist, err := client.GetArtist("artistName")
  events, err := client.GetArtistEvents("artistName")
}
```
