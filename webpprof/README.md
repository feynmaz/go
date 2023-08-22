# Web profiling guide

```
import (
    ...
    _ "net/http/pprof" // register pprof handlers
)
```

At `http://localhost:3000/debug/pprof/` see available profiles.

To read downloaded profiles:
```
go tools pprof <profile_file_name>
```
