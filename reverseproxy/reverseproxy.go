package reverseproxy

import (
	"fmt"
	"net/http"
	"net/url"
	"time"
)

const defaultTimeout = time.Second * 15

type ReverseProxy struct {
	source        string
	destination   *url.URL
	httpClient    *http.Client // httpClient is designed to be reused across multiple goroutines
	handleRequest func(http.ResponseWriter, *http.Request)
}

func New(sourceAddress string, sourcePort int, destinationUrl string) (*ReverseProxy, error) {
	parsedUrl, err := url.Parse(destinationUrl)
	if err != nil {
		return nil, err
	}

	proxy := &ReverseProxy{
		httpClient:  &http.Client{Timeout: defaultTimeout},
		source:      fmt.Sprintf("%s:%d", sourceAddress, sourcePort),
		destination: parsedUrl,
	}

	proxy.handleRequest = withLogging(withElapsedTime(proxy.handler))

	return proxy, nil
}

func (p *ReverseProxy) Run() error {
	return http.ListenAndServe(p.source, p)
}

func (p *ReverseProxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	p.handleRequest(w, r)
}
