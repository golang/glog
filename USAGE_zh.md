# slog 使用说明 #

[slog] 兼容 [glog] 命令行参数，使用方式同 [glog]

[slog]:https://github.com/hungrybirder/slog
[glog]:https://github.com/golang/glog

## 命令行设置说明 ##

-log_dir string

    进程日志目录，如果不设置或设置为空，则使用临时目录

-alsologtostderr

    配合 -log_dir，写日志文件，同时向 stderr 也输出一份日志

-log_backtrace_at value

    设置文件：行数，打印栈信息，比如 go run main.go -log_backtrace_at=main.go:20

-logtostderr

    设置了 -log_dir， 临时只写 stderr，不写日志文件

-stderrthreshold value

    设置 stderr 日志级别阈值，可选值 info|warning|error|fatal, 默认 stderr 级别是 error

-v value

    设置 Verbose 级别，比如 go run main.go -v=3

-vmodule value

    设置 module verbose 级别， 比如 go run main.go -vmodule=main=3,hello=2
