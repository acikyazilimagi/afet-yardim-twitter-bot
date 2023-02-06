# afet-yardim-bot
Web service that can retweet by user authorization with Twitter Oauth v1.

### Build
1. Download the ZIP file using git clone or the [link](https://github.com/acikkaynak/afet-yardim-bot/archive/refs/heads/main.zip).
2. Customize the values in the `main.go` file.
```go
var (
	flags          = flag.NewFlagSet("user-auth", flag.ExitOnError)
	consumerKey    = flags.String("consumer-key", "", "Twitter Consumer Key")
	consumerSecret = flags.String("consumer-secret", "", "Twitter Consumer Secret")
	accessToken    = flags.String("access-token", "", "Twitter Access Token")
	accessSecret   = flags.String("access-secret", "", "Twitter Access Secret")
)
```
3. Build this server.
```bash
go build .
```

### Run & Trigger
```
./afet-yardim-bot
```
Send Request
```url
http://localhost:8080/retweet?id=1622383183646937088
```
