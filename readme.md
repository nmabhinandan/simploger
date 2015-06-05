## Simploger

Simploger is a simple-to-use, level based logger for go programs.

[![GoDoc](https://godoc.org/github.com/nmabhinandan/simploger?status.svg)](https://godoc.org/github.com/nmabhinandan/simploger)

### tl;dr
```go
  import "github.com/nmabhinandan/simploger"
  func main() {
    sl := &simploger.Simplogger {
      Verbosity: 1,
      Logfile: simploger.Logfile{
        Win: "C:\\MyApp\\logs",
        Nix: "var/log/myapp/logs",
      },
    },

    sl.Err("Crashed!! ")
  }  

```
