# Modak Notification Center

 This service is in charge of managing user notifications; when we receive a request, we assume that the user already exists.

All the notifications that are configured in the database will be validated with the notifications received by the user.

If a notification cfg is not found, we still send it, this way if a configuration is missing, we do not block the alerts.

## Requirements

* Go 
* Docker
* MySql

## Installation

To perform a correct installation, it is recommended to use  

```bash
docker-compose up
```
to start the system

## Common setup

Clone the repo and start the server, dependencies will be downloaded automatically.

```bash
git clone https://github.com/torsello/modak-notification-center.git
cd modak-notification-center
```

```bash
go run .
```


## Usage 
Server starts listening on port: **8080**

When server starts, it automatically migrates two databases:
```bash
- rate_limit_cfgs
- user_notifications
```
then the default notification settings are inserted.
Currently it is done this way to provide a temporary solution, then it could be kept in an external file and read from there, or directly from the database.

To modify the configuration of these notifications you must go to: 
```bash
models/rate-limit-cfg.go:22
```

You can now test sending notifications to users, with the following request
```bash
curl --location 'http://localhost:8080/api/v1/notification' \
--header 'Content-Type: application/json' \
--data-raw '[
    {
        "type": "news",
        "receiver": "test@gmail.com",
        "message": "New products are waiting for you"
    },
    {
        "type": "status",
        "receiver": "test2@gmail.com",
        "message": "You have received money"
    },
    {
        "type": "status",
        "receiver": "test3@gmail.com",
        "message": "You made a payment"
    }
]''
```
After this we will receive the status of each notification.

- All notifications that were successful were saved in the database.

- All notifications that were failed, cannot be sent due to configuration validations.

```bash
Response example:
[
    {
        "type": "news",
        "receiver": "test@gmail.com",
        "message": "New products are waiting for you",
        "status": "successful"
    },
    {
        "type": "status",
        "receiver": "test2@gmail.com",
        "message": "You have received money",
        "status": "successful"
    },
    {
        "type": "status",
        "receiver": "test3@gmail.com",
        "message": "You made a payment",
        "status": "failed"
    }
]
```
## Server cfg

If you want to make some changes about the database connection, you should modify:

```bash
.env
```
```bash
DB_NAME=
DB_USERNAME=
DB_PASSWORD=
DB_HOST=
DB_PORT=
```
These values are exposed here for testing purposes, normally they would go in environment variables of the system where they are deployed.

## Unit Testing

To start all tests you must use:

```bash
go test .
```

## Troubleshooting:
- If you use docker compose, make sure you do not have a mysql service running on your pc listening on port 3306, docker creates a container listening on this port. 
## Swagger documentation:
You can find it at
```bash
docs/openapi.yaml
```
and also [here](https://app.swaggerhub.com/apis/MATIASTORSELLO/modak-notification-center/1.0.0)


## Postman collection:
[API Postman](https://api.postman.com/collections/8791767-dad193f3-0965-40ff-ab03-b84822d82c4d?access_key=PMAT-01HDJ1AJ8XFJA8ST0WNZMERDXS)

