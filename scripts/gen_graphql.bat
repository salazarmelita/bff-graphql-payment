@echo off
echo Generating GraphQL code...

go run github.com/99designs/gqlgen generate

echo Done!