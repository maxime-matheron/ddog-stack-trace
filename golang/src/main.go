package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"ddgolangstack/pkg"

	"github.com/gorilla/mux"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

func GenerateStack(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	span, ctx := tracer.StartSpanFromContext(ctx, "call.first_function")
	defer span.Finish()

	if err := pkg.FirstFunction(ctx); err != nil {
		span.SetTag("error", err)
	}

	fmt.Fprint(w, "<p>Golang stack trace generated!</p>")
	w.WriteHeader(http.StatusOK)
}

func main() {
	// Start Datadog tracer
	tracer.Start(
		tracer.WithAgentAddr("dd-agent"),
		tracer.WithServiceName(os.Getenv("DD_SERVICE")),
		tracer.WithSamplingRules(
			[]tracer.SamplingRule{tracer.RateRule(1)},
		),
	)
	defer tracer.Stop()

	// Register endpoints to router
	r := mux.NewRouter()
	r.HandleFunc("/generate-stack", GenerateStack)
	http.Handle("/", r)

	http.ListenAndServe(fmt.Sprintf(":%v", os.Getenv("SERVER_PORT")), r)
}
