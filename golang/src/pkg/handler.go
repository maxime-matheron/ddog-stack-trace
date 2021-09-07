package pkg

import (
	"context"
	"errors"

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
