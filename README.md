# playlist-media/mngo [![GoDoc](https://godoc.org/github.com/playlist-media/mngo?status.svg)](https://godoc.org/github.com/playlist-media/mngo)

A Golang MediaNet API client - it has two methods, `Setup` and `GetMP3`:

```go
func Setup(key, secret string)

func GetMP3(id int64, ip string) string
```
