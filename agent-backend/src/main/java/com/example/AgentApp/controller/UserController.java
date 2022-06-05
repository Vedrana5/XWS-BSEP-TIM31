package com.example.AgentApp.controller;

import com.example.AgentApp.model.User;

import com.example.AgentApp.security.TokenUtilss;
import com.example.AgentApp.service.UserService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.CrossOrigin;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import javax.servlet.http.HttpServletRequest;

@RestController
@RequestMapping(value = "/users")
public class UserController {


    @Autowired
    private TokenUtilss tokenUtils;

    @Autowired
    private UserService userService;

    @CrossOrigin(origins = "http://localhost:4200")
    @GetMapping(value = "/getBasicInf")
    public User getUserInformation(HttpServletRequest request) {
        String token = tokenUtils.getToken(request);
        String username=tokenUtils.getEmailFromToken(token);
        User user = userService.findByUsername(username);
        return user;
    }

    
}
