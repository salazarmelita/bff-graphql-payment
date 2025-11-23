@echo off
echo Running tests...

echo Running unit tests...
go test ./internal/application/service/... -v

echo Running domain tests...
go test ./internal/domain/... -v

echo All tests completed!