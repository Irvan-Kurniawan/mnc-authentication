## Requirements
To run this program, first the use must install libraries using the following command :

    go get -u github.com/gin-gonic/gin
    go get gorm.io/gorm
    go get gorm.io/driver/mysql
    go get golang.org/x/crypto/bcrypt
Also, the user must have mysql server running on port 3306 without password and have a database named `mnc_db`
If the user have password for the mysql database, please add the password after `root` using : in the `main.go` file, example :

    ...
    database.Connect("root@tcp(localhost:3306)/mnc_db?parseTime=true")
    ...
Change into :
 

	...
	database.Connect("root:yourpassword@tcp(localhost:3306)/mnc_db?parseTime=true")
	...

## Running The Program
Run the program using command `go run main.go`

## Using The Program
To use the program, there are examples in `json` folder.

1. One of the way to use files in `json` folder is to use `REST Client` extension in VSCode, the extension allows the use to run the HTTP request without using Postman.
 2. First the user must register a new customer using `registercustomer.rest` and `registercustomer2.rest` in order to transfer between the two customer. The username and the password for both customer may be changed as preferred.
 3. After registering customers, the user may login using `login.rest` file, change the username and password to the one that has been inputted earlier in `registercustomer.rest` or `registercustomer2.rest` (whichever logging in at the moment). The program will then give JWT token that is special for one specific customer, the token must be copied and saved.
 4. After login succesfully, the user can topup or transfer using `topupcustomer.rest` and `transfercustomer.rest`.

## Program Flow
When the program started, it will immediately try to connect to mysql server in port 3306 without password and will return error if no mysql server running and stop.
After the program successfully connect to mysql server, it will automatically migrate necessary table and columns so the program would run smoothly, namely `customers` and `histories` table, if the table didn't exist prior.
After that the program will be possible to be accessed through 5 routes :
| Route | Function |
|--|--|
| `/api/login` | To login and allow the user to topup or transfer |
| `/api/logout` | To logout |
| `/api/customer/register` | To register new customer |
| `/api/bank/topup` | To topup current logged in customer |
| `/api/bank/transfer` | To transfer balance to another customer |
Each time a user access any of the aforementioned route, `histories` table will be updated with the current action.

**Login Flow**
By accessing `/api/login`, the program will check if customer exists and will check if the password is correct, only then will the program give token to enable access to topup and transfer feature.

**Logout Flow**
By accessing `/api/logout`, the program will check if customer exists and already logeed in, only then will the program disable the customer `isLogin` to disable access to topup and transfer feature.

**Register Flow**
By accessing `/api/customer/register`, the program will hash the password, so nobody would be able to see the password as part of the security measurement, after that, username and password will be inserted to `customer` table. This program has validation for duplicate usernames, empty usernames or empty password.

**Topup Flow**
By accessing `/api/bank/topup`, the program will use middleware to check for token, invalid token or expired token or a token from logged out customer will not be able to access this route. After making sure that the token is valid, the program will get username from token, then will increase balance of the customer with that username. This program has validation for 0 or below 0 topup nominal.

**TransferFlow**
By accessing `/api/bank/transfer`, the program will use middleware to check for token, invalid token or expired token or a token from logged out customer will not be able to access this route. After making sure that the token is valid, the program will get username from token, then will increase balance of the target customer while also decreasing the sender balance equal to the nominal. This program has validation for 0 or below 0 transfer nominal, validation if the sender and receiver is the same, validation if sender or receiver do not exist, and validation if the nominal to transfer exceed the balance of sender.