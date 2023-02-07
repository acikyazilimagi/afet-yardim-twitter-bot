# afet-yardim-twitter-bot
Web service that can retweet by user authorization with Twitter Oauth v1.

### Build
1. Download the ZIP file using git clone or the [link](https://github.com/acikkaynak/afet-yardim-bot/archive/refs/heads/main.zip).
2. Create your own .env file using .env.example
3. Run.

### Requirements

go 1.19

### Configuration
    Configs can be changed in .env file. 

### Build:

    go build -o afet-yardim-twitter-bot ./cmd/main.go

### run:
    ./afet-yardim-twitter-bot
    
    or
    
    ./afet-yardim-twitter-bot -env-file=.env

### docker

    build:
    
        docker build -t afet-yardim-twitter-bot .

    run:
    
        docker run -d -p 8000:8000 --name afet-yardim-twitter-bot afet-yardim-twitter-bot
    
    docker-compose:
        
        docker-compose up


## Endpoints

### Retweet

**Endpoint:** `/retweet?id={tweetId}`

**Method:** `GET`

**Params:** `id: tweetId`


**Example URL:** http://localhost:8000/retweet?id=1242566795858345985

**Response Example:**

200 OK
```json
{
  "isSuccess": true
}
```

400 Bad Request
```json
{
  "status": 400,
  "error": "twitter: 327 You have already retweeted this Tweet."
}

