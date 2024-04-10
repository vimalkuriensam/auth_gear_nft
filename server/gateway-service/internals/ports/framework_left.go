package ports

import "net/http"

type RoutesPort interface {
	Routes() http.Handler
}
