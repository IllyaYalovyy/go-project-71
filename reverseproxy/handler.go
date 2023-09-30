package reverseproxy

import (
	"io"
	"log"
	"net/http"
	"path"
)

func (p *ReverseProxy) handler(w http.ResponseWriter, r *http.Request) {
	proxyRequest, err := p.createProxyRequest(r)
	if err != nil {
		log.Printf("Failed to create a proxy request. Error: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	proxyResponse, err := p.httpClient.Do(proxyRequest)
	if err != nil {
		log.Printf("Failed to call destination. Error: %v", err)
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}
	defer proxyResponse.Body.Close()

	copyHeaders(w.Header(), proxyResponse.Header)
	w.WriteHeader(proxyResponse.StatusCode)

	err = copyBody(w, proxyResponse.Body)
	if err != nil {
		log.Printf("Failed to write response body. Error: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func (p *ReverseProxy) createProxyRequest(r *http.Request) (*http.Request, error) {
	proxyUrl := *r.URL
	proxyUrl.Scheme = p.destination.Scheme // In case of on-host TLS termination
	proxyUrl.Host = p.destination.Host
	proxyUrl.Path = path.Join(p.destination.Path, r.URL.Path)

	proxyRequest, err := http.NewRequestWithContext(r.Context(), r.Method, proxyUrl.String(), r.Body)
	if err != nil {
		return nil, err
	}

	copyHeaders(proxyRequest.Header, r.Header)

	return proxyRequest, nil
}

func copyHeaders(dst, src http.Header) {
	for key, values := range src {
		for _, value := range values {
			dst.Add(key, value)
		}
	}
}

func copyBody(dst io.Writer, src io.Reader) error {
	_, err := io.Copy(dst, src)
	return err
}
