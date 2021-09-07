package com.example.pkg;

import com.example.pkg.Handler;
import datadog.trace.api.Trace;

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
            Exception chained = new Exception("failed when calling second_function");
            chained.initCause(e);
            throw chained;
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
            String s = "ababa";
            String[] a = s.split("a");
            System.out.println(a[5]);
        } catch (Exception e) {
            throw new Exception("failed in fourth_function");
        }
    }

}