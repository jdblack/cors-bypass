
# CORS Proxy

This is a simple CORS proxy written in Go that helps bypass Cross-Origin Resource Sharing (CORS) restrictions imposed by browsers. The proxy allows web applications to interact with resources hosted on different domains.

## Features

- Supports HTTP GET, POST, and OPTIONS methods.
- Automatically sets CORS headers to allow any origin (`*`).
- Handles preflight requests (OPTIONS).
- Accepts a URL parameter to fetch the desired resource.

## Installation

1. Clone the repository or download the source files.
2. Navigate to the directory containing the source files.
3. Run the following command to start the proxy server:

   ```bash
   go run .
   ```
4. On a well configured system with Go, one can type "go install ."


### Usage & Configuration

The proxy can be started by specifying the port to listen on using the command-line argument -p. If no port is provided, it defaults to 8080. You
can also configure the proxy with environment variables by using the format GOCORS_<ARGUMENT_NAME>.   For example, to change the port:

```bash
export GCORS_PORT=8000
go run .
```


#### Making a Request
To use the proxy, send an HTTP request to the server with a url query parameter pointing to the target resource. For example:

```bash
curl http://localhost:8080/?url=http://example.com/resource
Expected Response

The proxy will forward the request to the specified URL and return the response
from that resource, along with the appropriate CORS headers for localhost . In
case of errors, it will return an appropriate HTTP status code:

400 Bad Request if the url parameter is missing.
502 Bad Gateway if the target URL is unreachable.
500 Internal Server Error for issues reading the response.
```


