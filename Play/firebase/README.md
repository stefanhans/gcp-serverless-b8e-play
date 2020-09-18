## Imagine

- Create Firestore database, i.e. Datastore in native mode, to persist the confirmed bookings. 

- Implement CRUD 

## Create

Create Datastore in native mode in console for project `serverless-devops-play` - maybe not necessary?

Create Firebase project associated to `serverless-devops-play`

```bash
gcloud config list

firebase login
firebase init                         # see output in firebase_init.out
firebase projects:list
firebase database:instances:list

firebase database:instances:create bookings-play
firebase database:instances:list
```

Change Firestore rules
```text
rules_version = '2';
service cloud.firestore {
  match /databases/{database}/documents {
    match /{document=**} {
      allow read, write: if request.auth != null;
    }
  }
}
```

Set Service Account to access Firestore database

`export LOCAL_CREDENTIALS_DIR=<local credentials directory>`

`export PROJECT_ID=serverless-devops-play`

```bash
SA_NAME=firestore-play

gcloud iam service-accounts create ${SA_NAME} \
    --description="Service account to access Firestore API" \
    --display-name="${SA_NAME}"
    
gcloud projects add-iam-policy-binding ${PROJECT_ID} \
    --member "serviceAccount:${SA_NAME}@${PROJECT_ID}.iam.gserviceaccount.com" --role "roles/firebasedatabase.admin"    
    
    
gcloud iam service-accounts keys create ${LOCAL_CREDENTIALS_DIR}/${PROJECT_ID}-${SA_NAME}.json \
  --iam-account ${SA_NAME}@${PROJECT_ID}.iam.gserviceaccount.com
  

export GOOGLE_APPLICATION_CREDENTIALS="${LOCAL_CREDENTIALS_DIR}/${PROJECT_ID}-${SA_NAME}.json"
ls -l $GOOGLE_APPLICATION_CREDENTIALS
```


Edit `main.go` and run it
```bash
go mod init

go run main.go
```

## Play

##### Establish a commandline frontend and play:

Copy template directory as `commandline` from `github.com/stefanhans/frontend-liner-play/template`

See details of playing in [commandline/README](./commandline/README.md)