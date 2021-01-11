# TL;DR

We add a security configuration to the API config. We deploy, create a service account and test it. Then we create our 
own certificate and upload it to the service account. Finally, we create an appropriate JSON key file and test it.



# Prepare API Gateway

### Add Security Configuration

Edit API Config File "openapi2-functions.yaml"

```yaml
  # Security Object
  security:
    - user-keys_auth_id: []

  # Security Definition
  user-keys_auth_id:
    authorizationUrl: ""
    flow: "implicit"
    type: "oauth2"
    x-google-issuer: "sa-user-keys@serverless-devops-play.iam.gserviceaccount.com"
    x-google-jwks_uri: "https://www.googleapis.com/service_accounts/v1/jwk/sa-user-keys@serverless-devops-play.iam.gserviceaccount.com"
    x-google-audiences: "serverless-devops-play"
```

### Redeploy The Gateway
```bash
gcloud -q beta api-gateway api-configs delete b8e-api-config-jwt \
  --api=b8e-api && \
gcloud beta api-gateway api-configs create b8e-api-config-jwt \
  --api=b8e-api --openapi-spec="../openapi2-functions.yaml" \
  --backend-auth-service-account=sa-b8e-api@serverless-devops-play.iam.gserviceaccount.com && \
gcloud beta api-gateway gateways update b8e-api-gateway \
  --api=b8e-api --api-config=b8e-api-config-jwt \
  --location=europe-west1
```

### Test The Gateway

Create service account and key file for testing
```bash
gcloud iam service-accounts create sa-user-keys && \
gcloud iam service-accounts keys create /Users/stefan/.secret/sa-user-keys@serverless-devops-play.json \
  --iam-account=sa-user-keys@serverless-devops-play.iam.gserviceaccount.com
```

Test key file
```bash
go run main.go --host "https://b8e-api-gateway-5zpz279s.ew.gateway.dev/hello" \
  --audience "serverless-devops-play" \
  --service-account-file "/Users/stefan/.secret/sa-user-keys@serverless-devops-play.json" \
  --service-account-email "sa-user-keys@serverless-devops-play.iam.gserviceaccount.com"
  
# ... Response: Hello World!
```

### Clean-Up Test

Remove the key file
```bash
rm /Users/stefan/.secret/sa-user-keys@serverless-devops-play.json
```

Get service account key id
```bash
gcloud iam service-accounts keys list \
  --iam-account=sa-user-keys@serverless-devops-play.iam.gserviceaccount.com \
  --managed-by=user 
```

Delete the service account key
```bash
gcloud -q iam service-accounts keys delete 21d4a5e694df9e4928fa4cdfd8398ab53d713d5d \
  --iam-account=sa-user-keys@serverless-devops-play.iam.gserviceaccount.com
```

# Prepare Service Account

### Create Private Key And Certificate As New User "alice"

```bash
cd alice && \
openssl genpkey -algorithm RSA -out ca_private.pem -pkeyopt rsa_keygen_bits:2048 && \
openssl req -x509 -new -nodes -key ca_private.pem -sha256 -out ca_cert.pem -subj "/CN=unused"
openssl x509 -noout -text -in ca_cert.pem

# Limit the validity to 6 days
openssl req -x509 -new -nodes -key ca_private.pem -sha256 -out ca_cert_6d.pem -subj "/CN=unused" -days 6
openssl x509 -noout -text -in ca_cert_6d.pem
```

### Upload Certificate
```bash
gcloud iam service-accounts keys upload ca_cert.pem \
  --iam-account=sa-user-keys@serverless-devops-play.iam.gserviceaccount.com

gcloud iam service-accounts keys list \
  --iam-account=sa-user-keys@serverless-devops-play.iam.gserviceaccount.com  
```

### Create New JSON Key File For New User "alice"
```bash
../sa-user-keys_template.bash > alice.json
```

### Test
```bash
cd ../..

go run main.go --host "https://b8e-api-gateway-5zpz279s.ew.gateway.dev/hello" \
  --audience "serverless-devops-play" \
  --service-account-file "keys/alice/alice.json" \
  --service-account-email "sa-user-keys@serverless-devops-play.iam.gserviceaccount.com"

# ... Response: Hello World!
```