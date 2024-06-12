## README
### 1. Setup environment:

`docker-compose up -d`

### 2. Run server
`go run cmd/server-producer/*.go`

`go run cmd/server-consumer/*.go`

### 3. APIs
- Send notification

    `curl --location --request POST 'http://localhost:8080/notifications/send' \
  --header 'Content-Type: application/json' \
  --data-raw '{
  "from_user_id": 2,
  "to_user_id": 1,
  "message": "Hello from 2"
  }'`


- Get notifications by user_id

    `curl --location --request GET 'http://localhost:8081/notifications/1'`