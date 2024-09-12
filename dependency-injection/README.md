# Dependency Injection Pattern Example in Go

Dependency Injection (DI) is a design pattern used to implement IoC (Inversion of Control), allowing the creation of dependent objects to be separated from their behavior.

## Key Roles

> In real projects any names can be used. Here the roles are given to understand the concept.

1. **Service**: The unit that performs a specific task.
2. **Client**: The entity that requires a service to perform its function.
3. **Injector**: The entity that chooses the right service and injects it into the client.

In Dependency Injection, the injector provides services to clients through an interface. The client is unaware of the concrete service it is interacting with. The main advantage of DI is that it provides flexibility by injecting dependencies through interfaces.

## Example Overview

In our Go example, we have defined:
- An interface `Database` representing the service.
- Two concrete implementations of this interface: `MockPostgeDB` and `MockMongoDB`.
- A `Client` struct acting as the client.
- An injector that assigns a service instance to the client through the `NewClient` constructor.

## Diagram

Below is a simple ASCII diagram illustrating the relationship between client, service, and injector.

+----------+      +-----------------+      +-----------+
|  Client  | <--- |    Interface    | ---> |  Service  |
+----------+      +-----------------+      +-----------+
    ^                                           |
    |                                           |
    |           +--------------------+          |
    +---------- |      Injector      | <--------+
                +--------------------+


## How it Works

1. **Client Needs a Service**: The `Client` struct in our example requires a `Database` to function.

2. **Injector Provides Service**: The `NewClient` function acts as an injector. It accepts a `Database` interface and assigns it to the `Client`.

3. **Decoupled Relationship**: The client (`Client`) does not know whether it is using a `MockPostgeDB` or a `MockMongoDB`, thus achieving flexibility and decoupling through the use of interfaces.

## How to Run the Example

```bash
go run main.go
```
