package songs

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

const (
	startsWith = "/api/songs"
)

func getSongInfo(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	fmt.Fprintf(w, "get info for song %s", params.ByName("songId"))
}

func ConstructBlueprint(router *httprouter.Router) {
	router.GET(startsWith+"/:songId", getSongInfo)
}
