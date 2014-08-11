# playlist-media/mngo

A Golang MediaNet API client - it has two methods, `Setup` and `GetMP3`:

```go
func Setup(key, secret string)

func GetMP3(id int64, ip string) string
```
