package com.example.ddjavastack;

import com.example.pkg.Handler;
import org.springframework.web.bind.annotation.RestController;
import org.springframework.web.bind.annotation.RequestMapping;

@RestController
public class DDJavaStackController {

    @RequestMapping("/generate-stack")
    public String generateStack() throws Exception {
        try {
            Handler handler = new Handler();
            handler.firstFunction();
        } catch (Exception e) {
            // do nothing
        }
        return "Java stack trace generated!";
    }

}