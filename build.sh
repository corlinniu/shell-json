rm -rf bin
mkdir bin
go mod tidy
go build -o bin/json src/json.go
go build -o bin/csv src/csv.go
go build -o bin/excel src/excel.go
go build -o bin/html src/html.go