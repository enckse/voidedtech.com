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
