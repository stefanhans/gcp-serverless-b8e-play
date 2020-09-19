```bash
gcloud functions deploy add-vehicle --region europe-west3 \
    --entry-point AddVehicle --runtime go113 --trigger-http \
    --service-account=firestore-play@serverless-devops-play.iam.gserviceaccount.com \
    --allow-unauthenticated 
```
   
    
gcloud functions describe add-vehicle --region europe-west3 --format='value(httpsTrigger.url)'

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