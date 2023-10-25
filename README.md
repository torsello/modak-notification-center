# Modak Notification Center

This service is in charge of managing user notifications; when we receive a request, we assume that the user already exists.

All the notifications that are configured in the database will be validated with the notifications received by the user.

If a notification cfg is not found, we still send it, this way if a configuration is missing, we do not block the alerts.

## Requirements

- Go
- Docker
- MySql

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
models/rate_limit_cfg.go:22
```

You can now test sending notifications to users, with the following request

```bash
curl --location 'http://localhost:8080/api/v1/notification' \
--header 'Content-Type: application/json' \
--data-raw '{
    "data": {
        "notifications": [
            {
                "type": "news",
                "receiver": "test1@gmail.com",
                "message": "New products are waiting for you"
            },
            {
                "type": "status",
                "receiver": "test3@gmail.com",
                "message": "You have received money"
            },
            {
                "type": "update",
                "receiver": "test4@gmail.com",
                "message": "You made a payment"
            }
        ]
    }
}'
```

After this we will receive the status of each notification.

- All notifications that were successful were saved in the database.

- All notifications that were failed, cannot be sent due to configuration validations.

```bash
Response example:
{
    "data": {
        "notifications": [
            {
                "type": "news",
                "receiver": "test1@gmail.com",
                "message": "New products are waiting for you",
                "status": "failed"
            },
            {
                "type": "status",
                "receiver": "test3@gmail.com",
                "message": "You have received money",
                "status": "failed"
            },
            {
                "type": "update",
                "receiver": "test4@gmail.com",
                "message": "You made a payment",
                "status": "successful"
            }
        ]
    }
}
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
go test ./tests
```

**NOTE:** Some unit tests communicate directly to the database. I know that the best way is to mock the db, but I had problems with gorm

## Stress Testing

These tests were performed with Apache Benchmark on macOS on a server with the following characteristics:

- Processor: Intel Core i7 2.6GHz (Quad Core)
- RAM: 16 GB 2133 MHz

You can see the evidence [here](https://github.com/torsello/modak-notification-center/blob/main/docs/Stress%20testing%20-%20modak-notification-center.pdf)

## Troubleshooting:

- If you use docker compose, make sure you do not have a mysql service running on your pc listening on port 3306, docker creates a container listening on this port.
- If you use docker compose, note that the Go service does not start until the database is 100% online.
  This may take about 1 minute.

## Swagger documentation:

You can find it at

```bash
docs/openapi.yaml
```

and also [here](https://app.swaggerhub.com/apis/MATIASTORSELLO/modak-notification-center/1.0.1)

## Postman collection:

[API Postman](https://api.postman.com/collections/8791767-dad193f3-0965-40ff-ab03-b84822d82c4d?access_key=PMAT-01HDJ1AJ8XFJA8ST0WNZMERDXS)
