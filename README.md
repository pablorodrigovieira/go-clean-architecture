# Go Orders Listing Challenge

## 📌 Descrição (Português)

Olá devs!

Agora é a hora de botar a mão na massa. Para este desafio, você precisará criar o **usecase de listagem das orders**.

### ✅ Requisitos

- Esta listagem precisa ser feita com:
  - **Endpoint REST:** `GET /order`
  - **Service gRPC:** `ListOrders`
  - **Query GraphQL:** `ListOrders`
- Não esqueça de criar as **migrações necessárias**.
- Inclua um arquivo `api.http` com as requests para **criar e listar as orders**.
- Para a criação do banco de dados, utilize o **Docker** (`Dockerfile` / `docker-compose.yaml`).  
  Ao rodar `docker compose up`, tudo deverá subir e preparar o banco automaticamente.
- Inclua um **README.md** com os passos a serem executados no desafio.
- Informe a **porta** em que cada serviço irá responder.

## 📌 Description (English)

Hello developers!

Now it’s time to put your hands to work. For this challenge, you need to create a **use case for listing orders**.

### ✅ Requirements

- Implement the listing with:
  - **REST endpoint:** `GET /order`
  - **gRPC service:** `ListOrders`
  - **GraphQL query:** `ListOrders`
- Create the necessary **database migrations**.
- Include an `api.http` file with requests to **create and list orders**.
- Use **Docker** (`Dockerfile` / `docker-compose.yaml`) to create the database.  
  Running `docker compose up` should prepare everything automatically.
- Include a **README.md** explaining the steps to run the challenge.
- Specify the **ports** for each service in your README.

---

# Clean Architecture - Order System (Solution)

This service exposes REST, gRPC, and GraphQL endpoints for order creation, using MySQL and RabbitMQ as dependencies.

## Prerequisites

- Go 1.21+ installed
- Docker and Docker Compose (for MySQL and RabbitMQ)

## Start Dependencies (MySQL + RabbitMQ)

```zsh
docker compose up -d
```

- MySQL: `localhost:3306` (db `orders`, user `root`, pass `root`)
- RabbitMQ: `amqp://guest:guest@localhost:5672/` (Management UI `http://localhost:15672/#/`, user `guest`, password `guest`)

### Accessing MySQL via Docker

Use Docker Compose to open a shell inside the MySQL container and connect with the MySQL CLI:

```zsh
docker compose exec mysql bash
mysql -uroot -p orders
```

- When prompted for the password, enter `root`.
- The last argument `orders` selects the `orders` database after login.

````

## Environment

Create a `.env` file in `classes/clean-architecture/cmd/ordersystem` (same folder as this README):

```env
DB_DRIVER=mysql
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASSWORD=root
DB_NAME=orders
WEB_SERVER_PORT=8080
GRPC_SERVER_PORT=50051
GRAPHQL_SERVER_PORT=8081
RABBITMQ_URL=amqp://guest:guest@localhost:5672/
````

## Run

From the ordersystem command folder:

```zsh
cd classes/clean-architecture/cmd/ordersystem
# Important: run with the generated wire file
go run main.go wire_gen.go
```

## Testing the Web Endpoint

- File: `classes/clean-architecture/api/create_order.http`
- Usage (VS Code):
  - Install the "REST Client" extension (humao.rest-client).
  - Open the file and click "Send Request" above the `POST` line.
- Ensure `WEB_SERVER_PORT` matches the file. Example default is `8080`.

Example contents:

```http
POST http://localhost:8080/order HTTP/1.1
Host: localhost:8080
Content-Type: application/json

{
	"id": "123",
	"price": 100.0,
	"tax": 10.0
}
```

## Testing gRPC with Evans

- Install Evans: `brew install evans` (macOS)
- Open a new terminal and run:

```zsh
cd classes/clean-architecture
evans -r repl
```

- In the Evans REPL:
  - `package pb`
  - `service OrderService`
  - `call CreateOrder` and follow the prompts (provide `id`, `price`, `tax`).
  - If needed, set the host/port: `connect localhost:${GRPC_SERVER_PORT}` before selecting package/service.

## Testing GraphQL

- Open Playground: `http://localhost:${GRAPHQL_SERVER_PORT}/`
- Endpoint used by Playground: `/query`
- Run this mutation and variables:
- Run this mutation and variables (Playground supports both inline input and separate variables):

Mutation:

```graphql
mutation {
  createOrder(input: { id: "123", Price: 100.0, Tax: 10.0 }) {
    id
    Price
    Tax
    FinalPrice
  }
}
```

Expected response shape:

```json
{
  "data": {
    "createOrder": {
      "id": "123",
      "Price": 100,
      "Tax": 10,
      "FinalPrice": 110
    }
  }
}
```

## Testing RabbitMQ Setup (Queue + Binding)

This service publishes order events to RabbitMQ using exchange `amq.direct` with an empty routing key. To receive messages, create a queue and bind it to this exchange with an empty key.

Steps (via Management UI):

- Open `http://localhost:15672/#/` (user: `guest`, password: `guest`).
- Go to Queues → Add a new queue:
  - Name: `orders`
  - Durability: `Durable` (default)
  - Add queue
- Go to Exchanges → search/select `amq.direct`.
- In the Bindings section → "Add binding from this exchange":
  - Destination: `Queue`
  - Queue: `orders`
  - Routing key: leave empty
  - Add binding

Test the flow:

- Start the app and create an order (REST/gRPC/GraphQL examples above).
- In Queues → `orders`, click "Get messages" to see the published JSON payload.
