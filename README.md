# Go Orders Listing Challenge

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

````

## Environment

Create a `.env` file in `cmd/ordersystem` (same folder as this README):

```env
DB_DRIVER=mysql
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASSWORD=root
DB_NAME=orders
WEB_SERVER_PORT=:8000
GRPC_SERVER_PORT=50051
GRAPHQL_SERVER_PORT=8080
RABBITMQ_URL=amqp://guest:guest@localhost:5672/
````

## Run

From the ordersystem command folder:

```zsh
go run main.go wire_gen.go
```

## Testing the Web Endpoint

- File: `api/api.http`
- REST API runs on port **8000** (configured in `.env` as `WEB_SERVER_PORT`).

### Create Order

```http
POST http://localhost:8000/order HTTP/1.1
Host: localhost:8000
Content-Type: application/json

{
	"id": "123",
	"price": 100.0,
	"tax": 10.0
}
```

### List Orders

```http
GET http://localhost:8000/order HTTP/1.1
Host: localhost:8000
Content-Type: application/json
```

## Testing gRPC with Evans

- Open a new terminal and run:

```zsh
evans -r repl
```

- In the Evans REPL:
  - `package pb`
  - `service OrderService`
  - Call `CreateOrder`: type `call CreateOrder` and follow the prompts (provide `id`, `price`, `tax`).
  - Call `ListOrders`: type `call ListOrders` (no parameters required).

## Testing GraphQL

- Open Playground: `http://localhost:${GRAPHQL_SERVER_PORT}/`
- Endpoint used by Playground: `/query`
- Run the following mutations and queries:

### Create Order Mutation

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

Expected response:

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

### List Orders Query

```graphql
query {
  listOrders {
    id
    Price
    Tax
    FinalPrice
  }
}
```

Expected response:

```json
{
  "data": {
    "listOrders": [
      {
        "id": "123",
        "Price": 100,
        "Tax": 10,
        "FinalPrice": 110
      }
    ]
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
- Select the queue create and add binding: `amq.direct`.

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

---

# Clean Architecture - Sistema de Pedidos (Solução)

Este serviço expõe endpoints REST, gRPC e GraphQL para criação e listagem de pedidos, usando MySQL e RabbitMQ como dependências.

## Pré-requisitos

- Go 1.21+ instalado
- Docker e Docker Compose (para MySQL e RabbitMQ)

## Iniciar Dependências (MySQL + RabbitMQ)

```zsh
docker compose up -d
```

- MySQL: `localhost:3306` (db `orders`, usuário `root`, senha `root`)
- RabbitMQ: `amqp://guest:guest@localhost:5672/` (Interface de gerenciamento `http://localhost:15672/#/`, usuário `guest`, senha `guest`)

### Acessando MySQL via Docker

Use Docker Compose para abrir um shell dentro do container MySQL e conectar com o CLI MySQL:

```zsh
docker compose exec mysql bash
mysql -uroot -p orders
```

- Quando solicitado a senha, digite `root`.

## Ambiente

Crie um arquivo `.env` em `cmd/ordersystem`:

```env
DB_DRIVER=mysql
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASSWORD=root
DB_NAME=orders
WEB_SERVER_PORT=:8000
GRPC_SERVER_PORT=50051
GRAPHQL_SERVER_PORT=8080
RABBITMQ_URL=amqp://guest:guest@localhost:5672/
```

## Executar

A partir da pasta cmd/ordersystem:

```zsh
go run main.go wire_gen.go
```

## Testando o Endpoint REST

- Arquivo: `api/api.http`
  - Abra o arquivo e clique em "Send Request" acima de cada requisição.
- A API REST roda na porta **8000** (configurada no `.env` como `WEB_SERVER_PORT`).

### Criar Pedido

```http
POST http://localhost:8000/order HTTP/1.1
Host: localhost:8000
Content-Type: application/json

{
	"id": "123",
	"price": 100.0,
	"tax": 10.0
}
```

### Listar Pedidos

```http
GET http://localhost:8000/order HTTP/1.1
Host: localhost:8000
Content-Type: application/json
```

## Testando gRPC com Evans

- Abra um novo terminal e execute:

```zsh
evans -r repl
```

- No REPL do Evans:
  - `package pb`
  - `service OrderService`
  - Chamar `CreateOrder`: digite `call CreateOrder` e siga os prompts (forneça `id`, `price`, `tax`).
  - Chamar `ListOrders`: digite `call ListOrders` (não requer parâmetros).

## Testando GraphQL

- Abra o Playground: `http://localhost:8080/`
- Endpoint usado pelo Playground: `/query`
- Execute as seguintes mutations e queries:

### Mutation para Criar Pedido

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

Resposta esperada:

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

### Query para Listar Pedidos

```graphql
query {
  listOrders {
    id
    Price
    Tax
    FinalPrice
  }
}
```

Resposta esperada:

```json
{
  "data": {
    "listOrders": [
      {
        "id": "123",
        "Price": 100,
        "Tax": 10,
        "FinalPrice": 110
      }
    ]
  }
}
```

## Configuração do RabbitMQ (Fila + Binding)

Este serviço publica eventos de pedidos no RabbitMQ usando o exchange `amq.direct` com uma chave de roteamento vazia. Para receber mensagens, crie uma fila e vincule-a a este exchange com uma chave vazia.

Passos (via Interface de Gerenciamento):

- Abra `http://localhost:15672/#/` (usuário: `guest`, senha: `guest`).
- Vá em Queues → Adicionar uma nova fila:
  - Name: `orders`
  - Durability: `Durable` (padrão)
  - Add queue
- Selecione a fila criada e adicione binding: `amq.direct`.
