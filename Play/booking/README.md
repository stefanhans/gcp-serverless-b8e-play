## Imagine

- Define Go types for bookings 

- Implement client

    1) create a booking 
    2) marshal booking to JSON 
    3) send JSON request to server
    4) unmarshal JSON reply to booking
    
- Implement server

    1) receive JSON request
    2) unmarshal JSON request to booking
    3) confirm booking
    4) marshal booking to JSON reply
    5) send JSON reply to client

## Create

Create file structure
```bash
mkdir types server client
touch main.go types/types.go server/main.go client/main.go
```

Edit `types/types.go` and `main.go`, and run it
```bash
go run main.go
```

Edit `server/main.go` and run it in one terminal
```bash
cd server

go run main.go
```

Edit `client/main.go` and run it in another terminal
```bash
cd client

go run main.go
```

## Play

- More generalized Go types with interfaces?

- Use `func Book(w http.ResponseWriter, r *http.Request)` as Cloud Function?

##### Cloud Function


```bash
mkdir functions
cd functions

touch book.go

GO111MODULE=on
go mod init && go mod vendor
```

```bash
gcloud config configurations activate serverless-devops-play
gcloud auth login
gcloud services enable cloudbuild.googleapis.com

gcloud functions deploy book --region europe-west1 \
    --entry-point Book --runtime go113 --trigger-http \
    --allow-unauthenticated
```

```bash   
go run main.go
 
gcloud functions describe book --region europe-west1 --format='value(httpsTrigger.url)'

curl $(gcloud functions describe book --region europe-west1 --format='value(httpsTrigger.url)') \
    -d '{
            "id": 1,
            "user": {
                "id": 1,
                "name": "Alice"
            },
            "share": {
                "id": 1,
                "name": "Tesla Deluxe",
                "type": 0
            },
            "from": "2020-09-17T16:48:43.627845+02:00",
            "to": "2020-09-17T17:48:43.627848+02:00",
            "status": 0,
            "status-time": "2020-09-17T16:48:43.627848+02:00"
        }
'
```


## Share

## Reflect

