## Encrypt and decrypt data with Cloud KMS (Symmetric)

### Enable Cloud KMS Service

```bash
gcloud services enable cloudkms.googleapis.com
```

### Create KMS Keyring

```bash
gcloud kms keyrings create "my-keyring" \
    --location "global"
    
gcloud kms keyrings list \
    --location "global"

gcloud kms keyrings describe "my-keyring" \
    --location "global"
```

### Create KMS Keys

```bash
gcloud kms keys create "my-symmetric-key" \
    --location "global" \
    --keyring "my-keyring" \
    --purpose "encryption"

gcloud kms keys list \
    --keyring "my-keyring" \
    --location "global"

gcloud kms keys describe "my-symmetric-key" \
    --keyring "my-keyring" \
    --location "global"
```

### Encrypt Data

```bash
echo "my-first-contents" > ./data.txt

gcloud kms encrypt \
    --location "global" \
    --keyring "my-keyring" \
    --key "my-symmetric-key" \
    --plaintext-file ./data.txt \
    --ciphertext-file ./data.txt.enc

cat data.txt.enc | base64 > data.txt.enc.b64

cat data.txt.enc.b64
cat data.txt.enc.b64 | base64 -d
```

### Decrypt Ciphertext

```bash
gcloud kms decrypt \
    --location "global" \
    --keyring "my-keyring" \
    --key "my-symmetric-key" \
    --plaintext-file - \
    --ciphertext-file ./data.txt.enc
    
cat data.txt.enc.b64 | base64 -d | gcloud kms decrypt \
    --location "global" \
    --keyring "my-keyring" \
    --key "my-symmetric-key" \
    --plaintext-file - \
    --ciphertext-file -
```

### Rotate Keys

```bash
gcloud kms keys versions list \
    --location "global" \
    --keyring "my-keyring" \
    --key "my-symmetric-key"
    
gcloud kms keys versions describe 1 \
    --location "global" \
    --keyring "my-keyring" \
    --key "my-symmetric-key"
    
gcloud kms keys versions create \
    --location "global" \
    --keyring "my-keyring" \
    --key "my-symmetric-key" \
    --primary
    
gcloud kms keys versions list \
    --location "global" \
    --keyring "my-keyring" \
    --key "my-symmetric-key"
    
gcloud kms keys versions describe 2 \
    --location "global" \
    --keyring "my-keyring" \
    --key "my-symmetric-key"
```

```bash
gcloud kms keys versions disable "1" \
    --location "global" \
    --keyring "my-keyring" \
    --key "my-symmetric-key"

### not working as expected ???
gcloud kms decrypt \
    --location "global" \
    --keyring "my-keyring" \
    --key "my-symmetric-key" \
    --plaintext-file - \
    --ciphertext-file ./data.txt.enc
```
### Play

```bash
echo "my-second-contents" > ./data2.txt

gcloud kms encrypt \
    --location "global" \
    --keyring "my-keyring" \
    --key "my-symmetric-key" \
    --plaintext-file ./data2.txt \
    --ciphertext-file ./data2.txt.enc2

cat data2.txt.enc2 | base64 > data2.txt.enc2.b64

gcloud kms decrypt \
    --location "global" \
    --keyring "my-keyring" \
    --key "my-symmetric-key" \
    --plaintext-file - \
    --ciphertext-file data2.txt.enc2 # or data.txt.enc2 or data.txt.enc


ls -l data*.txt.enc*.b64
ls -l data*.txt

cat data*.txt.enc*.b64

cat data*.txt.enc*.b64 | base64 -d | gcloud kms decrypt \
    --location "global" \
    --keyring "my-keyring" \
    --key "my-symmetric-key" \
    --plaintext-file - \
    --ciphertext-file -
```

```bash
gcloud kms keys versions describe 2 \
    --location "global" \
    --keyring "my-keyring" \
    --key "my-symmetric-key"
    
gcloud kms keys describe "my-symmetric-key" \
    --location "global" \
    --keyring "my-keyring"
    
gcloud kms keyrings describe "my-keyring" \
    --location "global"
```


```bash
gcloud kms keys delete "my-asymmetric-encryption-key" \
    --location "global" \
    --keyring "my-keyring" \
    --purpose "asymmetric-encryption" \
    --default-algorithm "rsa-decrypt-oaep-2048-sha256"

gcloud kms keys list \
    --keyring "my-keyring" \
    --location "global"

gcloud kms keys describe "my-asymmetric-encryption-key" \
    --keyring "my-keyring" \
    --location "global"

    
gcloud kms keys versions get-public-key 1 \
    --location "global" \
    --keyring "my-keyring" \
    --key "my-asymmetric-encryption-key" \
    --output-file -
    
```

```bash
gcloud kms asymmetric-sign \
    --location "global" \
    --keyring "my-keyring" \
    --key "my-2nd-asymmetric-encryption-key" \
    --version 1 \
    --digest-algorithm "sha256" \
    --input-file "data.txt" \
    --signature-file "signature"
```

## Encrypt and decrypt data with Cloud KMS (Asymmetric)

### Create KMS Key

```bash
gcloud kms keys update "my-asymmetric-encryption-key" \
    --location "global" \
    --keyring "my-keyring" \
    --default-algorithm "rsa-decrypt-oaep-4096-sha512"
    
gcloud kms keys describe "my-asymmetric-encryption-key" \
    --location "global" \
    --keyring "my-keyring"
    
gcloud kms keys versions describe 1 --key "my-asymmetric-encryption-key" \
    --location "global" \
    --keyring "my-keyring"
    
gcloud kms keys versions create \
    --location "global" \
    --keyring "my-keyring" \
    --key "my-asymmetric-encryption-key"
    
gcloud kms keys versions describe 2 --key "my-asymmetric-encryption-key" \
    --location "global" \
    --keyring "my-keyring"
    
```

### Encrypt Data

```bash
gcloud kms keys versions get-public-key "1" \
    --location "global" \
    --keyring "my-keyring" \
    --key "my-asymmetric-encryption-key" \
    --output-file ./key.pub
    
openssl pkeyutl -encrypt -pubin \
    -in ./data.txt \
    -inkey ./key.pub \
    -pkeyopt "rsa_padding_mode:oaep" \
    -pkeyopt "rsa_oaep_md:sha512" \
    -pkeyopt "rsa_mgf1_md:sha512" > ./data.txt.asym
    
openssl pkeyutl -encrypt -pubin \
    -in ./data.txt \
    -inkey ./key.pub \
    -pkeyopt "rsa_padding_mode:oaep" \
    -pkeyopt "rsa_oaep_md:sha512" \
    -pkeyopt "rsa_mgf1_md:sha512" | base64 > ./data.txt.asym.b64
    
```

```bash


gcloud kms keys versions describe 1 --key "my-asymmetric-encryption-key" \
    --location "global" \
    --keyring "my-keyring"
    
gcloud kms asymmetric-decrypt \
    --location "global" \
    --keyring "my-keyring" \
    --key "my-asymmetric-encryption-key" \
    --version "1" \
    --plaintext-file - \
    --ciphertext-file ./data.txt.asym

```
gcloud kms asymmetric-decrypt \
        --location="global" \
        --keyring="my-keyring" \
        --key="my-asymmetric-encryption-key" \
        --version="1" \
        --ciphertext-file=./data.txt.asym \
        --plaintext-file=/tmp/my/secret.file.dec
        
## Sign and verify data with Cloud KMS (Asymmetric)

### Create KMS Keyring

```bash
gcloud kms keyrings create "sign-keyring" \
    --location "global"
    
gcloud kms keyrings list \
    --location "global"
    
gcloud kms keyrings describe "sign-keyring" \
    --location "global"
```

### Create KMS Key

```bash
gcloud kms keys create "my-asymmetric-signing-key" \
    --location "global" \
    --keyring "sign-keyring" \
    --purpose "asymmetric-signing" \
    --default-algorithm "rsa-sign-pkcs1-4096-sha512"

gcloud kms keys list \
    --location "global" \
    --keyring "sign-keyring" 

gcloud kms keys describe "my-asymmetric-signing-key" \
    --location "global" \
    --keyring "sign-keyring"
```

### Sign Data

```bash
echo "my-signed-contents" > ./data-sign.txt

gcloud kms asymmetric-sign \
    --location "global" \
    --keyring "sign-keyring" \
    --key "my-asymmetric-signing-key" \
    --version "1" \
    --digest-algorithm "sha512" \
    --input-file ./data-sign.txt \
    --signature-file ./data-sign.txt.sig
```

## Verify Data

```bash
gcloud kms keys versions get-public-key "1" \
    --location "global" \
    --keyring "sign-keyring" \
    --key "my-asymmetric-signing-key" \
    --output-file ./key-sign.pub
    
openssl dgst -sha256 \
    -verify ./key-sign.pub \
    -signature ./data-sign.txt.sig ./data-sign.txt
```
        
## Cleanup

```bash
gcloud kms keys versions destroy 1 \
    --location="global" \
    --keyring="my-keyring" \
    --key="my-asymmetric-encryption-key" 
    
gcloud kms keys versions destroy 2 \
    --location="global" \
    --keyring="my-keyring" \
    --key="my-asymmetric-encryption-key" 
    
    
gcloud kms keys versions destroy 3 \
    --location="global" \
    --keyring="my-keyring" \
    --key="my-asymmetric-encryption-key" 
    
gcloud kms keys versions list \
    --location="global" \
    --keyring="my-keyring" \
    --key="my-asymmetric-encryption-key"
    
    ###
    
gcloud kms keys versions destroy 1 \
    --location="global" \
    --keyring="my-keyring" \
    --key="my-symmetric-key" 
    
gcloud kms keys versions destroy 2 \
    --location="global" \
    --keyring="my-keyring" \
    --key="my-symmetric-key" 
    
gcloud kms keys versions destroy 3 \
    --location="global" \
    --keyring="my-keyring" \
    --key="my-symmetric-key" 
    
gcloud kms keys versions list \
    --location="global" \
    --keyring="my-keyring" \
    --key="my-symmetric-key"
```


gcloud kms keys versions create \
    --location "global" \
    --keyring "my-keyring" \
    --key "my-asymmetric-encryption-key"
    
gcloud kms keys versions list \
    --location "global" \
    --keyring "my-keyring" \
    --key "my-asymmetric-encryption-key"
    
gcloud kms keys versions get-public-key "4" \
    --location "global" \
    --keyring "my-keyring" \
    --key "my-asymmetric-encryption-key" \
    --output-file ./key.pub
    
gcloud kms asymmetric-decrypt \
    --location "global" \
    --keyring "my-keyring" \
    --key "my-asymmetric-encryption-key" \
    --version "4" \
    --plaintext-file - \
    --ciphertext-file ./data.txt.enc