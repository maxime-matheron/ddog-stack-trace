package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

func FourthFunction(ctx context.Context) error {
	span, _ := tracer.StartSpanFromContext(ctx, "run.operation")
	defer span.Finish()

	v := map[string]interface{}{}
	if _, ok := v["a"]; !ok {
		err := errors.New("failed in fourth_function")
		span.SetTag("error", err)
		return err
	}
	return nil
}

func ThirdFunction(ctx context.Context) error {
	span, ctx := tracer.StartSpanFromContext(ctx, "call.fourth_function")
	defer span.Finish()

	if err := FourthFunction(ctx); err != nil {
		span.SetTag("error", err)
		return errors.New("failed to call fourth_function in third_function")
	}
	return nil
}

func SecondFunction(ctx context.Context) error {
	span, ctx := tracer.StartSpanFromContext(ctx, "call.third_function")
	defer span.Finish()

	if err := ThirdFunction(ctx); err != nil {
		span.SetTag("error", err)
		return errors.New("failed to call third_function in second_function")
	}

	return nil
}

func FirstFunction(ctx context.Context) error {
	span, ctx := tracer.StartSpanFromContext(ctx, "call.second_function")
	defer span.Finish()

	if err := SecondFunction(ctx); err != nil {
		span.SetTag("error", err)
	}

	return nil
}

func GenerateStack(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	span, ctx := tracer.StartSpanFromContext(ctx, "call.first_function")
	defer span.Finish()

	if err := FirstFunction(ctx); err != nil {
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
