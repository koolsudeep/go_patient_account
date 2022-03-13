run : build
		./bin/patient_account

build :
		go build -o ./bin/patient_account cmd/api/main.go


clean :
		go clean

test :
		go test ./...
