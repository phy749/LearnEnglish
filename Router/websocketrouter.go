package router

import (
	controller "github.com/phy749/LearnEnglish/Controller"
	dataoject "github.com/phy749/LearnEnglish/dataoject"
	"net/http"
)

func InitWebSocketRoute(mux *http.ServeMux, hub *dataoject.Hub) {
	mux.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		controller.HandleWebSocket(hub, w, r)
	})
}
