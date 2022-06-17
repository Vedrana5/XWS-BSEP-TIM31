package com.example.AgentApp.controller;

import com.example.AgentApp.dto.ChangePasswordDto;
import com.example.AgentApp.model.User;

import com.example.AgentApp.security.TokenUtilss;
import com.example.AgentApp.service.UserService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

import javax.servlet.http.HttpServletRequest;
import javax.validation.Valid;

@RestController
@RequestMapping(value = "/users")
public class UserController {


    @Autowired
    private TokenUtilss tokenUtils;

    @Autowired
    private UserService userService;

    @CrossOrigin(origins = "https://localhost:4200")
    @GetMapping(value = "/getBasicInf")
    public User getUserInformation(HttpServletRequest request) {
        String token = tokenUtils.getToken(request);
        String username=tokenUtils.getEmailFromToken(token);
        User user = userService.findByUsername(username);
        return user;
    }


}
