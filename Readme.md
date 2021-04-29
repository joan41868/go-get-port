# go-get-port

### A helper library, which fetches a random port.
### Ranges:
* MIN_BOUNDARY = 1
* MAX_BOUNDARY = 2^16 (65536)

## How do i get to use this lib?

```sh
go get github.com/joan41868/go-get-port@v0.1.0
```
## Usage

```go

package main


import ( 
    gp "go_get_port"
    "log"
    )

func main(){
    port := gp.GetPort()
    fmt.Println(port) // some random port
}

```