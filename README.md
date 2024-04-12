# go-clock

1. Tidy build
   `go mod tidy`
1. Run locally
   `go .`
1. Compile for ARM
   `env GOOS=linux GOARCH=arm go build -o dist/go-clock -v .`
