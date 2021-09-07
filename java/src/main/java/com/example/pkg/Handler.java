package com.example.pkg;

import java.util.HashMap;
import io.opentracing.Span;
import io.opentracing.Scope;
import io.opentracing.Tracer;
import com.example.pkg.Handler;
import datadog.trace.api.Trace;
import io.opentracing.util.GlobalTracer;

public class Handler {

    @Trace(operationName = "first_function")
    public void firstFunction() throws Exception {
        this.secondFunction();
    }

    @Trace(operationName = "second_function")
    public void secondFunction() throws Exception {
        try {
            this.thirdFunction();
        } catch (Exception e) {
            throw new Exception("failed when calling third_function");
        }
    }

    @Trace(operationName = "third_function")
    public void thirdFunction() throws Exception {
        try {
            this.fourthFunction();
        } catch (Exception e) {
            throw new Exception("failed when calling fourth_function");
        }
    }

    @Trace(operationName = "fourth_function")
    public void fourthFunction() throws Exception {
        try {
            HashMap<String, Boolean> map = new HashMap<String, Boolean>();
            System.out.println(map.get(true));
        } catch (Exception e) {
            throw new Exception("failed in fourth_function");
        }
    }

}