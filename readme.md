# Simple Banking App with GO

## To initialize a module

RUN `go mod init github.com/Altamashattari/banking-application`

RUN `go mod tidy`

## To run the app

go build
./banking-application

## Run MYSQL on docker
```docker pull --platform linux/x86_64 mysql```
```docker run --name some-mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=password -d mysql```
