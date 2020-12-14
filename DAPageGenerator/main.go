package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func logRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}

type Exampler interface {
	Examples()
}

func main() {

	examplers := []Exampler{
		BarExamples{},
		// examples.Bar3dExamples{},
		// examples.BoxplotExamples{},
		// examples.EffectscatterExamples{},
		// examples.FunnelExamples{},
		// examples.FunnelExamples{},
		// examples.GaugeExamples{},
		GeoExamples{},
		// examples.GraphExamples{},
		// examples.HeatmapExamples{},
		// examples.KlineExamples{},
		// examples.LineExamples{},
		// examples.Line3dExamples{},
		// examples.LiquidExamples{},
		// examples.MapExamples{},
		// examples.ParallelExamples{},
		PieExamples{},
		// examples.RadarExamples{},
		// examples.CustomizeExamples{},
		// examples.SankeyExamples{},
		// examples.ScatterExamples{},
		// examples.Scatter3dExamples{},
		// examples.Surface3dExamples{},
		// examples.ThemeriverExamples{},
		WordcloudExamples{},
	}

	for _, e := range examplers {
		e.Examples()
	}

	fs := http.FileServer(http.Dir("examples/html"))
	log.Println("running server at http://localhost:8089")
	log.Fatal(http.ListenAndServe("localhost:8089", logRequest(fs)))
}

func DBErrLoggerOutput(w gin.ResponseWriter, err error) {
	// dbLogger.Printf("%v err=[%v]\n", MyNowTimeFormatted(), err)
	fmt.Fprintf(w, "<p>Sorry, DB err: %v</p>\n", err)
}
