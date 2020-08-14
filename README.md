# github.com/hungrybirder/glog #

Clone from [glog](https://github.com/golang/glog) because of [issue45](https://github.com/golang/glog/pull/45).

## Usage ##

Please __DON'T__ `import "github.com/hungrybirder/glog"` in your code.

__Do `import "github.com/golang/glog"`__ and
replace `github.com/golang/glog` to `github.com/hungrybirder/glog`.

```bash
go mod edit -replace github.com/golang/glog=github.com/hungrybirder/glog@v1.0.0
```

## Add *Depthf Functions ##

* __func InfoDepthf(depth int, format string, args ...interface{})__

* __func WarningDepthf(depth int, format string, args ...interface{})__

* __func ErrorDepthf(depth int, format string, args ...interface{})__

* __func FatalDepthf(depth int, format string, args ...interface{})__

* __func ExitDepthf(depth int, format string, args ...interface{}__
