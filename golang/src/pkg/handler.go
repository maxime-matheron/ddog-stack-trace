package pkg

import (
	"context"
	"errors"

	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

func FirstFunction(ctx context.Context) error {
	span, ctx := tracer.StartSpanFromContext(ctx, "first_function")
	defer span.Finish()

	if err := SecondFunction(ctx); err != nil {
		span.SetTag("error", err)
	}

	return nil
}

func SecondFunction(ctx context.Context) error {
	span, ctx := tracer.StartSpanFromContext(ctx, "second_function")
	defer span.Finish()

	if err := ThirdFunction(ctx); err != nil {
		err := errors.New("error when calling third_function in second_function")
		span.SetTag("error", err)
		return err
	}

	return nil
}

func ThirdFunction(ctx context.Context) error {
	span, ctx := tracer.StartSpanFromContext(ctx, "third_function")
	defer span.Finish()

	if err := FourthFunction(ctx); err != nil {
		err := errors.New("error when calling fourth_function in third_function")
		span.SetTag("error", err)
		return err
	}

	return nil
}

func FourthFunction(ctx context.Context) error {
	span, _ := tracer.StartSpanFromContext(ctx, "fourth_function")
	defer span.Finish()

	v := map[string]interface{}{}
	if _, ok := v["a"]; !ok {
		err := errors.New("failed in fourth_function")
		span.SetTag("error", err)
		return err
	}

	return nil
}
