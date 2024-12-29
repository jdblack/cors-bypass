# CORS Proxy

### Overview

CORS (Cross-Origin Resource Sharing) is a security feature implemented by web
browsers to prevent unauthorized cross-origin requests. The idea behind CORS is
that a web page can specify which sites are allowed to interact with its
content. This prevents malicious actors from performing man-in-the-middle
attacks by fetching sensitive sites, like a bank's web page, and presenting
them to users within a frame. Overall, this is a beneficial security measure.

However, this security layer can interfere with certain types of web
applications. For example, TypingMind is a web application that provides a
frontend for accessing multiple LLM backends like Claude and OpenAI. Users of
TypingMind utilize plugins to access third-party resources such as websites,
search engines, and APIs. The data returned from these sources can have CORS
restrictions that define the allowed contexts for viewing the data. The user's
web browser checks these restrictions and refuses to interact with the data if
the origin doesn't match.

Some API services set a CORS host restriction to "localhost," allowing usage
from the command line with tools like `curl` or directly in browsers like
Chrome. However, web applications operate under a CORS location that reflects
the website's domain (e.g., typingmind.com) rather than localhost.
Consequently, when the browser attempts to fetch the API request, it recognizes
the localhost restriction and discards the request, deeming it unsafe.

### What CORS Bypass Does

Cors-bypass... well, it bypasses that cors restriction by fetching the
requested page, removing the CORS restriction, then returning it to the user.

This program is designed to handle simple GET requests, enabling developers to
fetch web pages from different origins. It effectively bypasses CORS
restrictions by returning the requested page with a wildcard CORS origin. This
approach works particularly well with APIs that impose CORS origin
restrictions, such as NewsAPI.org. Please note that this application does not
support form submissions or JavaScript execution; its primary purpose is to
retrieve pages and provide a straightforward solution for accessing external
content.

### Features
- Supports HTTP GET requests
- Can communicate with HTTPS endpoints
- Automatically sets CORS headers to allow any origin (`*`)

### Installation
1. Clone the repository or download the source files.
2. Navigate to the directory containing the source files.
3. Run the following command to start the proxy server:

   ```bash
   go run .
   ```
4. On a well-configured system with Go, you can also run:

   ```bash
   go install .
   ```

### Usage & Configuration

The proxy can be started by specifying the port to listen on using the
command-line argument `-p`. If no port is provided, it defaults to 8080. You
can also configure the proxy using environment variables in the format
`GOCORS_<ARGUMENT_NAME>`. For example, to change the port:

```bash
export GCORS_PORT=8000
go run .
```

#### Making a Request
To use the proxy, send an HTTP request to the server with a `url` query
parameter pointing to the target resource. For example:

```bash
curl http://localhost:8080/?url=http://example.com/resource
```

The proxy will forward the request to the specified URL and return the response
from that resource, along with the appropriate CORS headers for localhost.

### Dedication

This app is dedicated to NewsAPI.org, which provides an excellent news search
API engine! Unfortunately, they lack affordable payment plans for hobbyists
such as ourselves. Hopefully, one day they'll offer the same sort of
$fewdollars/mREQ model that has been adopted by the LLM industry! 

