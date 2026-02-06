# BCC Canteen

## API Installation

Here is how to run the service in your local environment.<br>

### 1. Using Docker (recommended)
In the project's root directory run:
```
$ docker-compose up --build
```

### 2. Manual setup
   Prerequisite :
   - Go v1.25
   - MariaDB
   
Create new database and run `db/init.sql` with your favourite tool for the table and dummy data.<br>

Copy variables in `.env.example` and create an `.env` file in the project's root directory, then match the variables with your environment. Make sure that the variables' name left unchanged.<br>

In `internal/storage/mariadb/bcc_canteen.go`, untag these lines:
```Go
if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

log.Printf("Environment variables loaded")
```

In the project's root directory run:
```bash
$ ./run.sh
```
or
```bash
$ go build -o bcc-canteen ./cmd/api
$ ./bcc-canteen
```
## API Testing
To test or having a look on the endpoints, you can either access the online Postman workspace: <br>
[https://www.postman.com/kresnawan1/bcc-canteen](https://www.postman.com/kresnawan1/bcc-canteen) <br>
or import the collections JSON in [docs](https://github.com/kresnawan/freepass-2026/tree/kresnawan/docs). <br>

> [!CAUTION]
> All requests require `url` environment variable and some requires `access_token` environment variables for the authentication. Also, some endpoints have query params and path variable, so pay attention to those.

Here are some decent information about the API <br>

## Response object
Every endpoint across the service has the same response format
```json
{
   "success": 1,  // Either 1 or 0
   "message": "", // Response message
   "data": {}     // Resources if any (e.g. All menu, token, etc)
}
```
## Root user
By default, the root user credentials are:<br><br>
username: root<br>
password: root123

## Canteen ownership
Yes, every canteen could have multiple owners and every owner could have multiple canteens as well, it is by design.

## Access and refresh token
As you managed to log in, server will set a `refreshToken` cookie, and return the access token in the response. Put the access token you received in Postman environment variable to proceed with auth-protected endpoints.<br>
Access token will valid for 10 minutes by default, when your access token was expired, all you need to do is just make a request to `/api/v1/token` route to generate new access token without relogging again

## Order
You must put items on your cart before you place your order. Once an order placed, the menu stock will decreased by order quantity.<br>
In the database, order status is represented by integer, which has meaning as follows: <br><br>
6 : Waiting payment<br>
7 : Cooking <br>
8 : Ready <br>
9 : Completed <br>

## Payment
This app comes with assumption that every payment will automatically verified through a payment gateway or stuff so canteen owner isn't needed to verify it manually. Just like some popular food delivery service nowadays.

# Flow
So, from the very scratch, the flow are as follows:

1. Admin make a canteen
2. Admin make an owner account
3. Admin add the canteen ownership to the owner
4. Owner make a menu for a canteen
5. Owner add stocks to the menu
6. Customer see menus
7. Customer put menu on their cart
8. Customer place order
9. Customer pays order
10. Customer wait for the order
11. Owner update menu status
12. Order status is completed
13. Customer give a feedback
14. Owner can delete the feedback if it inappropriate


<!-- ## **📞** Contact

Have any questions? You can contact [Atha](https://www.instagram.com/mhqif/).
## **🎁** Submission

Please follow the instructions on the [Contributing guide](CONTRIBUTING.md).

![cheers](https:
> This is not the only way to join us.
>
> **But, this is the _one and only way_ to instantly pass.** -->

