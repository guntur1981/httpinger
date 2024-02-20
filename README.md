# HTTP Pinger

A simple demonstration of Server Sent Events (SSE) using Golang.

## Description

Like a ping utility used to test the reachability of a host on an Internet Protocol (IP) network, httpinger is a utility to test the availability of a HTTP(S) server.

The server sends a HEAD request to the target URL every 2 seconds then the server sends the results back to the client browser using Server Sent Events (SSE) technology.

## Pros and Cons SSE over Websocket

### Pros

-   Transported over simple HTTP instead of a custom protocol
-   No trouble with corporate firewalls doing packet inspection
-   Built in support for re-connection and event-id

### Cons

-   No binary support, only UTF-8
-   Maximum open connections limit, six concurrent connections per browser.

## Run It

```
go run main.go
```

Open `http://localhost:3000` in your browser.

![Demo](/demo.png)
