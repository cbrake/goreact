app_setup() {
  go get -u github.com/jteeuwen/go-bindata/...
  go get -u github.com/golang/dep/cmd/dep
  dep ensure
}

app_build_frontend() {
  (cd frontend && npm run build) || return 1
  echo "generate bindata ..."
  #go-bindata-assetfs -prefix "frontend/dist" frontend/dist/...
  go-bindata -o frontend/bindata.go -pkg frontend -prefix "frontend/dist" frontend/dist/... || return 1
  #(cd frontend/dist && go-bindata-assetfs -o ../bindata.go -pkg frontend ./...) || return 1
  return 0
}

app_build_backend() {
  go build -o goreact cmd/app/main.go || return 1
  go build -o sample-post cmd/sample-post/main.go || return 1
  return 0
}

app_build() {
  app_build_frontend || return 1
  app_build_backend || return 1
  return 0
}

app_build_and_run() {
  echo "rebuilding backend ..."
  app_build_backend || return 1
  echo "starting app ..."
  ./app --dev || return 1
  return 0
}

app_find_backend_files() {
  find -name "vendor" -prune \
    -o -name "*.go" -print
}

app_watch_backend() {
  app_find_backend_files | entr -r ./build_and_run.sh || return 1
  return 0
}

app_watch_frontend() {
  (cd frontend && npm run watch) || return 1
  return 0
}

# the below would be used for running in development
# is not working yet, so just run app_watch_backend
# and app_watch_frontend in separate terminal windows
app_watch() {
  app_watch_backend &
  app_watch_frontend &
  trap 'kill %1; kill %2' SIGINT
  wait
  trap - SIGINT
}

app_build_other() {
  app_build_frontend || return 1
  GOOS=windows GOARCH=amd64 go build -o goreact.exe cmd/app/main.go || return 1
  GOOS=windows GOARCH=amd64 go build -o goreact.exe cmd/app/main.go || return 1
  GOOS=darwin GOARCH=amd64 go build -o goreact_mac cmd/app/main.go || return 1
  GOOS=linux GOARCH=arm go build -o goreact_arm cmd/app/main.go || return 1
}
