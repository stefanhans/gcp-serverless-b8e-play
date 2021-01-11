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
            "Name": "Tesla Standard",
            "Type": "eCar",
            "Status": "",
            "ParkingLot": "",
            "GeoPoint": {
              "latitude": 0.0,
              "longitude": 0.1
            },
            "Description": "just a test description"
          }
'
``` 

##### GetVehicleById
```bash
GO111MODULE=on
go mod vendor
```

```bash
gcloud functions deploy get-vehicle-by-id --region europe-west3 \
    --entry-point GetVehicleById --runtime go113 --trigger-http \
    --service-account=firestore-play@serverless-devops-play.iam.gserviceaccount.com \
    --allow-unauthenticated 
```
   
```bash
curl $(gcloud functions describe get-vehicle-by-id --region europe-west3 --format='value(httpsTrigger.url)') \
    -d '{
            "DocId": "QShcfGS53yOoGjR8rYJD"
        }
'
``` 

##### DeleteVehicleById
```bash
GO111MODULE=on
go mod vendor
```

```bash
gcloud functions deploy delete-vehicle-by-id --region europe-west3 \
    --entry-point DeleteVehicleById --runtime go113 --trigger-http \
    --service-account=firestore-play@serverless-devops-play.iam.gserviceaccount.com \
    --allow-unauthenticated 
```
   
```bash
curl $(gcloud functions describe delete-vehicle-by-id --region europe-west3 --format='value(httpsTrigger.url)') \
    -d '{
            "DocId": "V9YvdcjtOPBaDwKZDzBo"
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
            "Name": "Id Test",
            "Type": "Testuser",
            "Status": "",
            "Description": "just a test description"
          }
'
``` 

##### GetUserById
```bash
GO111MODULE=on
go mod vendor
```

```bash
gcloud functions deploy get-user-by-id --region europe-west3 \
    --entry-point GetUserById --runtime go113 --trigger-http \
    --service-account=firestore-play@serverless-devops-play.iam.gserviceaccount.com \
    --allow-unauthenticated 
```
   
```bash
curl $(gcloud functions describe get-user-by-id --region europe-west3 --format='value(httpsTrigger.url)') \
    -d '{
            "DocId": "sSkEJTEl5kswlLV3Goy5"
        }
'
``` 

##### DeleteUserById
```bash
GO111MODULE=on
go mod vendor
```

```bash
gcloud functions deploy delete-user-by-id --region europe-west3 \
    --entry-point DeleteUserById --runtime go113 --trigger-http \
    --service-account=firestore-play@serverless-devops-play.iam.gserviceaccount.com \
    --allow-unauthenticated 
```
   
```bash
curl $(gcloud functions describe delete-user-by-id --region europe-west3 --format='value(httpsTrigger.url)') \
    -d '{
            "DocId": "sSkEJTEl5kswlLV3Goy5"
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
    --entry-point GetUsers --runtime go113 --trigger-http \
    --service-account=firestore-play@serverless-devops-play.iam.gserviceaccount.com \
    --allow-unauthenticated 
```
   
```bash
curl $(gcloud functions describe get-users --region europe-west3 --format='value(httpsTrigger.url)')
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
            "User": "6zspJAh9sdVqxkX6g7RZ",
            "Vehicle": "QShcfGS53yOoGjR8rYJD",
            "VehicleType": "small car",
            "VehicleStatus": "",
            "ParkingLot": "",
            "From": "2020-12-30T00:00:00Z",
            "To": "2020-12-31T00:00:00Z",
            "Status": "requesting"
          }
'
```  

##### GetBookingById
```bash
GO111MODULE=on
go mod vendor
```

```bash
gcloud functions deploy get-booking-by-id --region europe-west3 \
    --entry-point GetBookingById --runtime go113 --trigger-http \
    --service-account=firestore-play@serverless-devops-play.iam.gserviceaccount.com \
    --allow-unauthenticated 
```
   
```bash
curl $(gcloud functions describe get-booking-by-id --region europe-west3 --format='value(httpsTrigger.url)') \
    -d '{
            "DocId": "cKlBkP5X4DfmPZDpInbi"
        }
'
```   

##### DeleteBookingById
```bash
GO111MODULE=on
go mod vendor
```

```bash
gcloud functions deploy delete-booking-by-id --region europe-west3 \
    --entry-point DeleteBookingById --runtime go113 --trigger-http \
    --service-account=firestore-play@serverless-devops-play.iam.gserviceaccount.com \
    --allow-unauthenticated 
```
   
```bash
curl $(gcloud functions describe delete-booking-by-id --region europe-west3 --format='value(httpsTrigger.url)') \
    -d '{
            "DocId": "BlPz9H0jcr2HFWnfptm6"
        }
'
```  

##### GetBookingsByRange
```bash
GO111MODULE=on
go mod vendor
```

```bash
gcloud functions deploy get-bookings-by-range --region europe-west3 \
    --entry-point GetBookingsByRange --runtime go113 --trigger-http \
    --service-account=firestore-play@serverless-devops-play.iam.gserviceaccount.com \
    --allow-unauthenticated 
```
   
```bash
curl $(gcloud functions describe get-bookings-by-range --region europe-west3 --format='value(httpsTrigger.url)') \
    -d '{
            "from": "2020-09-19T15:02:21.111175Z",
            "to": "2021-09-19T17:02:21.111176Z"
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

##### GetMasterData
```bash
GO111MODULE=on
go mod vendor
```

```bash
gcloud functions deploy get-master-data --region europe-west3 \
    --entry-point GetMasterData --runtime go113 --trigger-http \
    --service-account=firestore-play@serverless-devops-play.iam.gserviceaccount.com \
    --allow-unauthenticated 
```
   
```bash
curl $(gcloud functions describe get-master-data --region europe-west3 --format='value(httpsTrigger.url)') -d '{}'
``` 



##### Play
```bash
GO111MODULE=on
go mod init && go mod vendor
```

```bash
gcloud functions deploy play --region europe-west3 \
    --entry-point Play --runtime go113 --trigger-http \
    --service-account=firestore-play@serverless-devops-play.iam.gserviceaccount.com \
    --allow-unauthenticated 
```
   
```bash
curl $(gcloud functions describe add-vehicle --region europe-west3 --format='value(httpsTrigger.url)') \
    -d '{
            "DocId": "",
            "Name": "Tesla Standard",
            "Type": "eCar",
            "Status": "",
            "ParkingLot": "",
            "GeoPoint": {
              "latitude": 0.0,
              "longitude": 0.1
            },
            "Description": "just a test description"
          }
'
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