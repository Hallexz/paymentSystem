# Payment Microservice

This is a microservice written in Go for handling payments using gRPC and REST. 

Use tools:
- Golang 1.22
- Protobuf
- Docker
- Kubernetes
  

## Features

- **Payment Service**: Handles payment requests and processes them using a simulated bank.
- **Transaction Service**: Manages transactions and stores them in an in-memory data structure.
- **User Service**: Handles user-related operations (not implemented in the provided code).

## Getting Started

1. Clone the repository:

```
https://github.com/Hallexz/paymentSystem.git
```

2. Navigate to the project directory:

```
cd paymentSystem
```

3. Build the Docker image:

```
docker build -t payment-microserv .
```

4. Run the Docker container:

```
docker run -p 50051:50051 payment-microserv
```

The microservice will start running on `localhost:50051`.

## Usage

The microservice exposes the following gRPC services:

- `PaymentService`
  - `Pay(PaymentRequest) returns (PaymentResponse)`: Processes a payment request.
- `TransactionService`
  - `GetTransactions(GetTransactionsRequest) returns (stream Transaction)`: Retrieves a stream of transactions.
- `UserService`
  - `CreateUser(CreateUserRequest) returns (CreateUserResponse)`: Creates a new user (not implemented).
  - `GetUser(GetUserRequest) returns (User)`: Retrieves user information (not implemented).

You can use a gRPC client (e.g., gRPC CLI, Bloom RPC, or a custom client) to interact with the microservice.

## Kubernetes Deployment

The repository includes a Kubernetes deployment manifest (`payment-deployment.yaml`) for deploying the microservice on a Kubernetes cluster. The deployment creates three replicas of the microservice.

To deploy the microservice on a Kubernetes cluster, run:

```
kubectl apply -f payment-deployment.yaml
```

## Contributing

Contributions are welcome! If you find any issues or have suggestions for improvements, please open an issue or submit a pull request.
