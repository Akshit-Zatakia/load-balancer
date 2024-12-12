# Load Balancer

To control traffic across servers in a network, load-balancing algorithms are important. By spreading requests evenly, load balancers make sure that no single server is overloaded when several people visit an application.

There are several algorithms for load balancing. For this application below mentioned alorightms are included:

1. Round Robin: <br>
   The Round Robin algorithm is a simple static load balancing approach in which requests are distributed across the servers in a sequential or rotational manner.
2. Least Connections: <br>
   The Least Connections algorithm is a dynamic load balancing approach that assigns new requests to the server with the fewest active connections.
3. Least Response Time: <br>
   The Least Response method is a dynamic load balancing approach that aims to minimize response times by directing new requests to the server with the quickest response time.

## Config File

Based on the strategy defined in the yaml file the load balancer will work.

```yaml
lb_port: 3332
strategy: <least-response or least-connection or round>
backends:
  - "http://localhost:3333"
  - "http://localhost:3334"
```

## How to Run the application

- Define the yaml file based on the description given above.
- Start the mock server using below python commands:
  ```
  python3 -m http.server 3333
  python3 -m http.server 3334
  ```
- Go to the project directory and run the following command:
  ```
  go run main.go
  ```
