# golang-hello-word
Simple Golang web service that replies with JSON including an env var - used for other demos.

### build

```
# specific to my environment
docker build -t us-central1-docker.pkg.dev/e2m-private-test-01/golang-hello-word/golang-hello-word:latest .
docker push us-central1-docker.pkg.dev/e2m-private-test-01/golang-hello-word/golang-hello-word:latest
```