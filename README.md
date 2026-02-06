# BCC Canteen

## API Installation

Here is how to run the service in your local environment<br>

### 1. Using Docker (recommended)
In the project's root directory run:
```
$ docker-compose up --build
```

### 2. Manual setup
   Prerequisite :
   - Go v1.25
   - MariaDB
   
Run `db/init.sql` and `db/dml.sql` with your favourite tool for the table and dummy data (The init will make a new database automatically, so be careful for the naming conflict)<br>

Copy variables in `.env.example` and create an `.env` file in the project's root directory, then match the variables with your environment. Make sure that the variables' name left unchanged<br>

In `internal/storage/mariadb/bcc_canteen.go`, untag these lines:
```Go
if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

log.Printf("Environment variables loaded")
```

In the project's root directory run:
```bash
$ go build -o bcc-canteen ./cmd/api
$ ./bcc-canteen
```
## API Testing
To test or having a look on the endpoints, you can either access the online Postman workspace: <br>
[https://www.postman.com/kresnawan1/bcc-canteen](https://www.postman.com/kresnawan1/bcc-canteen) <br>
or import the collection JSONs in `docs/`. <br>
Here are some decent information about the API
> [!CAUTION]
> All requests require `url` environment variable and some requires `access_token` environment variables for the authentication. Also, some endpoints have query params and path variable, so pay attention to those.

### Response object
Every endpoint across the service has the same response format
```json
{
   "success": 1,  // Either 1 or 0
   "message": "", // Response message
   "data": {}     // Resources if any (e.g. All menu, token, etc)
}
```

### Access and refresh token
As you managed to log in, server will set a `refreshToken` cookie, and return the access token in the response. Put the access token you received in Postman environment variable to proceed with auth-protected endpoints.<br><br>
Access token will valid for 10 minutes by default, when your access token was expired, all you need to do is just make a request to `/api/v1/token` route to generate new access token without relogging again

### Order
You must put items on your cart before you place your order. Once an order placed, the menu stock will decreased by order quantity

### Payment
This app comes with assumption that every payment will automatically verified through a payment gateway or stuff so canteen owner isn't needed to verify it manually. Just like some popular food delivery service nowadays.

<!-- ## **📞** Contact

Have any questions? You can contact [Atha](https://www.instagram.com/mhqif/).
## **🎁** Submission

Please follow the instructions on the [Contributing guide](CONTRIBUTING.md).

![cheers](https:
> This is not the only way to join us.
>
> **But, this is the _one and only way_ to instantly pass.** -->

