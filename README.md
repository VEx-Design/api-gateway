# API Gateway

## Introduction
The API Gateway acts as a central entry point for managing and routing API requests to various microservices. It provides authentication.

## Features
- Centralized API management
- Authentication and authorization

## Installation Dependencies
Ensure you have Go installed. After cloning the project, install the required dependencies:
```bash
go mod tidy
```

## Environment Variables
Create a `.env` file in the root directory and configure it based on `.env.example`.

## Running the Project
To run the project, use the following command:

```bash
make restart
```

## Docker
### Building the Docker Image and Running the Container
```bash
make up
```
