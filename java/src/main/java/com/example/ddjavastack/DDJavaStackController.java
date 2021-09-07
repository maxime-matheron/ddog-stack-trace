package com.example.ddjavastack;

import io.opentracing.Span;
import io.opentracing.Scope;
import io.opentracing.Tracer;
import com.example.pkg.Handler;
import io.opentracing.util.GlobalTracer;
import org.springframework.web.bind.annotation.RestController;
import org.springframework.web.bind.annotation.RequestMapping;

@RestController
public class DDJavaStackController {

    @RequestMapping("/generate-stack")
    public String generateStack() {
        Tracer tracer = GlobalTracer.get();
        Span span = tracer.buildSpan("call.first_function").start();
        try (Scope scope = tracer.activateSpan(span)) {
            Handler handler = new Handler();
            handler.firstFunction();
        } catch (Exception e) {
            // Set error on span
        } finally {
            // Close span in a finally block
            span.finish();
        }

        return "Java stack trace generated!";
    }

}