# Profiling guide
Within the video guide ["Benchmark tests and profiling in Go"](https://youtu.be/0UBv8GV8XdY).

### Commands

Run tests
```
go test .
```

Run benchmarks

```
go test -bench .
```

Run benchmarks for a certain time

```
go test -bench . -benchtime 10s
```

Run benchmarks and output memory profile
```
go test -bench . -memprofile mem.out
```
Read mem.out file
```
go tool pprof mem.out
```
Output graph in Gif format
```
(pprof) gif
```