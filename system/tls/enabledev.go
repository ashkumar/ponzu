package tls

import (
	"log"
	"net/http"
	"path/filepath"
)

// EnableDev generates self-signed SSL certificates to use HTTPS & HTTP/2 while
// working in a development environment. The certs are saved in a different
// directory than the production certs (from Let's Encrypt), so that the
// acme/autocert package doesn't mistake them for it's own.
// Additionally, a TLS server is started using the default http mux.
func EnableDev() {
	setupDev()

	cert := filepath.Join("devcerts", "cert.pem")
	key := filepath.Join("devcerts", "key.pem")
	err := http.ListenAndServeTLS(":10443", cert, key, nil)
	if err != nil {
		log.Fatalln(err)
	}
}
