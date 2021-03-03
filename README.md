# Library for testing oauth implementation with JWT tokens

## How to run it in Docker

    docker build -t oauth-test -f docker/Dockerfile .
    docker run -it -p 8000:8000 oauth

## How to generate a JWT token with curl

    curl -XPOST 127.0.0.1:8000/token --header 'content-type: application/x-www-form-urlencoded' --data grant_type=client_credentials -u admin:foobar -d scope=fosite -d audience=admins

## How to get private and public keys

Just watch to the console output when the library is started

## Hot to verify the access token

- Copy the `access_token` value from the curl command mentioned above 
- Go to https://jwt.io/
- Copy public key and see if "Signature Verified" is shown

## To deploy in K8s


    kubectl apply -f deployments/k8s

