export GOPATH=${PWD}
export GOBIN=${PWD}/bin
export GOTOOLING=${PWD}/tooling
export PATH=${PATH}:${GOBIN}:${GOTOOLING}/bin
# export GOROOT =
export GOOS=darwin
export GOARCH=amd64

# SET RUNTIME ENV VARS
export ZIPKIN_SERVICE_HOST=localhost
export ZIPKIN_SERVICE_PORT=9411
export SERVICE_HOST_OVERRIDE=localhost:8000
#export GOOGLE_PUBSUB_PROJECT_ID
