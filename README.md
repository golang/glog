# github.com/hungrybirder/slog #

1. Clone from  because [glog] of [issue45](https://github.com/golang/glog/pull/45).

2. [slog] is compatible with [glog] flags.

3. [slog] and [glog] can be used both in your project.


## Install ##

`go get github.com/hungrybirder/slog`

## Usage ##

```
import "github.com/hungrybirder/slog"

slog.Error("hello slog")
slog.Errorf("%s slog", "Have a nice day")
```


## More features ##

### Add *Depthf Functions ###

* __func InfoDepthf(depth int, format string, args ...interface{})__

* __func WarningDepthf(depth int, format string, args ...interface{})__

* __func ErrorDepthf(depth int, format string, args ...interface{})__

* __func FatalDepthf(depth int, format string, args ...interface{})__

* __func ExitDepthf(depth int, format string, args ...interface{}__

[glog]: https://github.com/golang/glog
[slog]: https://github.com/hungrybirder/slog
