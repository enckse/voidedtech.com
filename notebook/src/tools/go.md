Go
===

Some notes on working in go

## Profiling

To do some CPU profiling, the following snippet could be used

```
import "runtime/pprof"

//...

f, err := os.Create("cpu.prof")
if err != nil {
    log.Fatal(err)
}
pprof.StartCPUProfile(f)
defer pprof.StopCPUProfile()
```

Then run
```
go tool pprof cpu.prof
```

To interrogate the resulting output file (e.g. `top10`).

### Testing

alternatively call:

```
go test -cpuprofile cpu.out -memprofile mem.out .
```

on a single package (as that is all it allows)

## gopls/goimports

`gopls` can be used like `goimports` from the command line
```
gopls format <goimports arguments>
```
