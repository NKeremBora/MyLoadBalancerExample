# Simple TCP Load Balancer

This project implements a simple **TCP Load Balancer** in Go, using the `net` package. It listens for incoming client connections on a specified port (default: `8080`) and forwards requests to a set of backend servers using the **Round Robin** algorithm.

## Features

- **TCP Proxy**: Handles client requests and forwards them to backend servers.
- **Round Robin Algorithm**: Distributes incoming requests across backend servers in a sequential manner.
- **Simple Backend Servers**: Can be started using `npx http-server` or other simple HTTP server utilities.
- **Scalable Design**: Easily extendable to support additional load-balancing algorithms.

---

## How It Works

1. The load balancer listens on port `8080` for incoming TCP connections.
2. Each incoming request is forwarded to a backend server (e.g., `localhost:5001`, `localhost:5002`, etc.).
3. The backend server processes the request and sends a response back to the client through the load balancer.

---

## Getting Started

### Prerequisites

- **Go (Golang)** installed on your system. Download from [golang.org](https://golang.org/dl/).
- **Node.js** and **npm** for starting backend servers using `npx http-server`.

---

### Running the Project

#### 1. Clone the Repository
```bash
git clone https://github.com/NKeremBora/MyLoadBalancerExample.git
cd MyLoadBalancerExample
```
#### 2. Start Backend Servers
```bash
npx http-server -p 5001
npx http-server -p 5002
npx http-server -p 5003
```

#### 3. Start main.go and curl 8080 port
```bash
go run main.go
curl localhost:8080
```
