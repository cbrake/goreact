app_setup() {
  go get -u github.com/jteeuwen/go-bindata/...
  go get -u github.com/kardianos/govendor
  npm install -g elm
  govendor sync
}

app_build_frontend() {
  (cd frontend && elm-package install --yes) || return 1
  (cd frontend && elm-make --yes main.elm --output elm.js) || return 1
  go-bindata -o frontend/bindata.go -pkg frontend frontend/index.html frontend/elm.js frontend/bootstrap.min.css || return 1
  return 0
}

app_build() {
  app_build_frontend || return 1
  go build -o app cmd/app/main.go || return 1
  go build -o sample-post cmd/sample-post/main.go || return 1
  return 0
}

app_build_windows() {
  app_build_frontend || return 1
  GOOS=windows GOARCH=amd64 go build -o portal.exe cmd/portal/main.go || return 1
  GOOS=windows GOARCH=amd64 go build -o vc-sim.exe cmd/vc-sim/main.go || return 1
}
