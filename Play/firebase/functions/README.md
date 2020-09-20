## Imagine

Deploy Cloud Functions to handle HTTP request from the frontend to the Firestore database

1) master data 
    - `Vehicles`
    - `Users`
1) bookings
    - `Bookings`

## Create

##### AddVehicle
```bash
GO111MODULE=on
go mod init && go mod vendor
```

```bash
gcloud functions deploy add-vehicle --region europe-west3 \
    --entry-point AddVehicle --runtime go113 --trigger-http \
    --service-account=firestore-play@serverless-devops-play.iam.gserviceaccount.com \
    --allow-unauthenticated 
```
   
```bash
curl $(gcloud functions describe add-vehicle --region europe-west3 --format='value(httpsTrigger.url)') \
    -d '{
            "DocId": "",
            "Name": "Tesla Deluxe",
            "Type": "eCar",
            "Status": "",
            "ParkingLot": "",
            "GeoPoint": {
              "latitude": 0.1,
              "longitude": 0.1
            },
            "Description": "just a test description"
          }
'
``` 

##### ClearVehicles
```bash
GO111MODULE=on
go mod vendor
```

```bash
gcloud functions deploy clear-vehicles --region europe-west3 \
    --entry-point ClearVehicles --runtime go113 --trigger-http \
    --service-account=firestore-play@serverless-devops-play.iam.gserviceaccount.com \
    --allow-unauthenticated 
```
   
```bash
curl $(gcloud functions describe clear-vehicles --region europe-west3 --format='value(httpsTrigger.url)')
``` 

##### GetVehicles
```bash
GO111MODULE=on
go mod vendor
```

```bash
gcloud functions deploy get-vehicles --region europe-west3 \
    --entry-point GetVehicles --runtime go113 --trigger-http \
    --service-account=firestore-play@serverless-devops-play.iam.gserviceaccount.com \
    --allow-unauthenticated 
```
   
```bash
curl $(gcloud functions describe get-vehicles --region europe-west3 --format='value(httpsTrigger.url)')
``` 

##### AddUser
```bash
GO111MODULE=on
go mod vendor
```

```bash
gcloud functions deploy add-user --region europe-west3 \
    --entry-point AddUser --runtime go113 --trigger-http \
    --service-account=firestore-play@serverless-devops-play.iam.gserviceaccount.com \
    --allow-unauthenticated 
```
   
```bash
curl $(gcloud functions describe add-user --region europe-west3 --format='value(httpsTrigger.url)') \
    -d '{
            "DocId": "",
            "Name": "Alice",
            "Type": "Testuser",
            "Status": "",
            "Description": "just a test description"
          }
'
``` 

##### ClearUsers
```bash
GO111MODULE=on
go mod vendor
```

```bash
gcloud functions deploy clear-users --region europe-west3 \
    --entry-point ClearUsers --runtime go113 --trigger-http \
    --service-account=firestore-play@serverless-devops-play.iam.gserviceaccount.com \
    --allow-unauthenticated 
```
   
```bash
curl $(gcloud functions describe clear-users --region europe-west3 --format='value(httpsTrigger.url)')
``` 

##### GetUsers
```bash
GO111MODULE=on
go mod vendor
```

```bash
gcloud functions deploy get-users --region europe-west3 \
    --entry-point GetVehicles --runtime go113 --trigger-http \
    --service-account=firestore-play@serverless-devops-play.iam.gserviceaccount.com \
    --allow-unauthenticated 
```
   
```bash
curl $(gcloud functions describe get-vehicles --region europe-west3 --format='value(httpsTrigger.url)')
``` 

##### AddBooking
```bash
GO111MODULE=on
go mod vendor
```

```bash
gcloud functions deploy add-booking --region europe-west3 \
    --entry-point AddBooking --runtime go113 --trigger-http \
    --service-account=firestore-play@serverless-devops-play.iam.gserviceaccount.com \
    --allow-unauthenticated 
```
   
```bash
curl $(gcloud functions describe add-booking --region europe-west3 --format='value(httpsTrigger.url)') \
    -d '{
            "DocId": "",
            "User": "Alice",
            "Vehicle": "mycar",
            "VehicleType": "my type",
            "VehicleStatus": "",
            "ParkingLot": "",
            "Status": ""
          }
'
``` 

##### ClearBookings
```bash
GO111MODULE=on
go mod vendor
```

```bash
gcloud functions deploy clear-bookings --region europe-west3 \
    --entry-point ClearBookings --runtime go113 --trigger-http \
    --service-account=firestore-play@serverless-devops-play.iam.gserviceaccount.com \
    --allow-unauthenticated 
```
   
```bash
curl $(gcloud functions describe clear-bookings --region europe-west3 --format='value(httpsTrigger.url)')
``` 

##### GetBookings
```bash
GO111MODULE=on
go mod vendor
```

```bash
gcloud functions deploy get-bookings --region europe-west3 \
    --entry-point GetBookings --runtime go113 --trigger-http \
    --service-account=firestore-play@serverless-devops-play.iam.gserviceaccount.com \
    --allow-unauthenticated 
```
   
```bash
curl $(gcloud functions describe get-bookings --region europe-west3 --format='value(httpsTrigger.url)')
``` 



curl $(gcloud functions describe add-vehicle --region europe-west3 --format='value(httpsTrigger.url)') \
    -d '{
            "Name": "Tesla Deluxe",
            "Type": "eCar",
            "GeoPoint": {
              "latitude": 0.1,
              "longitude": 0.1
            },
            "Description": "just a test description"
          }
'