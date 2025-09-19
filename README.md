# 1. Hasznos eszközök és csomagok
`go test` – a Go beépített teszt futtatója
`go test -v` – részletesebb kimenet
`go test -cover` – teszt lefedettség
`golangci-lint` – statikus kódelemző
`go test -v -cover -bench=. -benchmem`
Testify – segédkönyvtár, ha bővebb assert-eket szeretnél: https://github.com/stretchr/testify

```bash
hello-go
go get github.com/stretchr/testify  //module download
go get github.com/stretchr/testify/assert@v1.10.0
go install github.com/kisielk/errcheck@latest
export PATH=$PATH:$(go env GOPATH)/bin
errcheck .

go mod init example.com/hello
go mod tidy                         // module cleaner
go run main.go                      // build, run and remove code
go build -o main main.go            // build code
```

# 2.Linting
```bash
go install -v github.com/go-delve/delve/cmd/dlv@latest
go install github.com/golangci/golangci-lint/v2/cmd/golangci-lint@latest
go vet                              // syntax check
go fmt                              // code formatting
golangci-lint run                   // go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
```

# 3. Ajánlott források
https://go.dev/ – hivatalos oldal
https://gobyexample.com/ – interaktív Go példák
https://quii.gitbook.io/learn-go-with-tests/ – TDD alapú Go tanulás – ezt nagyon ajánlom, pont azt tanítja, amit keresel.
https://pkg.go.dev/ - csomagkezelo
