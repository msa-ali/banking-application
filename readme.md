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

## Run the Application

```SERVER_ADDRESS=localhost SERVER_PORT=8000 DB_USER=root DB_PASSWD=password DB_ADDR=localhost DB_PORT=3306 DB_NAME=banking go run main.go```

## Write an API to create a New Account for an existing customer

### Acceptance Criteria

- A new account can only be opened with a mininum deposit of 5000.00
- Account can only be of saving or checking type
- In case of an unexpected error, api should return status code 500 (Internal Server Error) along with the error message.
- The API should return the new account id, when the new account is opened with the status code as 201(created)

### Make a transaction in bank account

## AC

- transaction can only be withdrawl or deposit
- amount can't be negative
- withdrawal amount should be available in the account
- successful transaction, should return the updated balance with transaction id response
- error handling should be done for the bad request and unexpected errors from the server side and should return the appropriate http status code with message

## Authorization

### ADMIN USER

- GET ALL CUSTOMERS     GET /customer
- GET CUSTOMER BY ID    GET /customer/{customer_id}
- CREATE NEW ACCOUNT    POST /customer/{customer_id}/account
- MAKE A TRANSACTION    POST /customer/{customer_id}/account/{account_id}

## USER ROLE

- GET CUSTOMER BY ID    GET /customer/{customer_id}
- MAKE A TRANSACTION    POST /customer/{customer_id}/account/{account_id}

SERVER_ADDRESS=localhost SERVER_PORT=8000 DB_USER=root DB_PASSWD=password DB_ADDR=localhost DB_PORT=3306 DB_NAME=banking go run main.go