# Summary
A simple reverse proxy implementation to learn how to do networking in Golang.

# References
- https://pkg.go.dev/net/
- https://en.wikipedia.org/wiki/Decorator_pattern
- https://www.iana.org/assignments/http-status-codes/http-status-codes.xhtml
- https://www.nginx.com/resources/glossary/reverse-proxy-server/
- https://en.wikipedia.org/wiki/Reverse_proxy
- https://code.visualstudio.com/
- https://go.dev/
- https://www.oreilly.com/library/view/designing-distributed-systems/9781491983638/ch02.html
- https://en.wikipedia.org/wiki/TLS_termination_proxy
- https://en.wikipedia.org/wiki/Weighted_round_robin
- https://medium.com/geekculture/timeout-context-in-go-e88af0abd08d

# Contribution

## Prerequisites
- Basic knowledge of Go programming language
- Git
- IDE

## Getting Started
- Clone the repository:

```sh
git clone https://github.com/IllyaYalovyy/go-project-71
```

- Go into the project directory:

```sh
cd go-project-71
```

- Build the project:

```sh
go build
```

- Test the build:

```sh
go test
```

## Coding Best Practices

Before submitting your PR, please consider following:
- The change doesn't break thread safety
- The change reuse existing code where it makes sense
- The change introduces reusable functions and structures
- The change makes project more maintainable 
- The change doesn't reduce test coverage of the project
- The change is easy to understand
- The change doesn't include any unrelated changes (e.g. style update, random fixes in unrelated files)

## Code duplication
 
  “Duplication is the primary enemy of a well-designed system.” 
― Robert C. Martin, [The Robert C. Martin Clean Code Collection (Collection)](https://www.goodreads.com/work/quotes/18312943)
 
  “It can be better to copy a little code than to pull in a big library for one function. Dependency hygiene trumps code reuse.”
― Rob Pike, [Go Proverbs](https://go-proverbs.github.io/)

We should avoid code duplication in this project. But be careful to not introduce a significant dependency for a very little value.
Try to avoid premature generalization, and start refactoring the code, only when you have enough evidence that the use case can be implemented in a generic way.

## Code Style
Go comes with code style guide and checker. It's no ones favorite code style, and that is a beauty of it. The default formatter [gofmt](https://golang.org/cmd/gofmt/) is pretty loose, so for this project the recommendation is to use [gofumpt](https://github.com/mvdan/gofumpt).

But the bottom-line is that any violation of the accepted code style should break the build. We don't want to spend our time on discussion style in PRs. If you feel strong about anything related to the code style, just send us PR.

## Unit tests

Use simple rule: "The code doesn't work if it has no unit tests".

## Logging
- More better than less (at least in the beginning)
  - It's not a bad idea to add some extra logging initially. At any moment the amount and contents of log messages can be updated or reduced, but it way worse when we don't have enough information to troubleshoot customer or security issue.
- Searchable
  - How we can find all relevant log records? Can we use prefix search for that? Are we starting log message with unrelated generic words or even worse - some dynamic values?
- structured
  - https://medium.com/@hectorjsmith/structured-logging-in-go-6a5c6cbc0730
- no duplication

...
yes, it's a very deep topic. 

## Comments
  “Every time you write a comment, you should grimace and feel the failure of your ability of expression.” 
― Robert C. Martin, [The Robert C. Martin Clean Code Collection (Collection)](https://www.goodreads.com/work/quotes/18312943)
 
Comments are a manifestation of a bad design and code organization in most cases. The [API documentation](https://pkg.go.dev/golang.org/x/tools/cmd/godoc) is an exception.
 
More reading: [Comments Do Not Make Up for Bad Code](https://w.amazon.com/bin/view/Marketplace/Training/Coding_Practices/Clean_Code_Handbook#Comments_Do_Not_Make_Up_for_Bad_Code)

# Design overview

The current design and implementation are rudimentary, and don't cover many use cases. This implementation potentially can be used for on host TLS termination, when it runs as a side car with the server it protects.

## Limitation
- no support for URL query, and other advanced features
- Only supports 1 destination (ideally we should be able to have a pool of backend servers with some routing algorithm, e.g. weighted round robin)
- It uses default pool of goroutines. It can be adjusted based on metrics and hardware.
- It produces only basic metrics (elapsed time). Additional metrics should be added for healthcheck and monitoring. 
- It doesn't support any caching of static content. 
- It doesn't support "session affinity". Why would it? it just support one destination :D. Also, no one sane is writing stateful services any more. 
- It doesn't support any request filtering (authN, throttling, etc)

## Scalability

- The usual way to scale a system is to implement a stateless service, have a cluster that run instances of this service and have a load balancer in front of it. 
- This approach not only allows for horizontal scale, but also helps with availability, and maintainability (see blue/green deployments). 
- This proxy can reduce load on the primary service by doing TLS termination, request filtering, throttling, authentication of requests (if applicable), caching static content. 

## Security

- In current implementation secure connection is on enforced. It's not acceptable in real production. The option to enforce secure connection should be added (HTTPS).
- We should configure proper timeouts to make service more reliable, and protect from some types of attacks. (e.g. slow HTTP post attack)
- We can implement authentication. (e.g. OAuth2)
- We can implement rate limiting. (protects the backend from a denial of service attacks)
- Use web application firewall in front of the proxy server.


