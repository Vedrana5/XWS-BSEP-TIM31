package com.example.AgentApp.controller;

import com.example.AgentApp.dto.JwtAuthenticationRequestDto;
import com.example.AgentApp.dto.LogUserDto;
import com.example.AgentApp.dto.RegisterDto;
import com.example.AgentApp.dto.UserTokenState;
import com.example.AgentApp.model.User;
import com.example.AgentApp.model.UserDetails;
import com.example.AgentApp.model.UserRole;
import com.example.AgentApp.repository.UserRepository;
import com.example.AgentApp.security.TokenUtils;
import com.example.AgentApp.service.CustomTokenService;
import com.example.AgentApp.service.UserService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.security.authentication.AuthenticationManager;
import org.springframework.security.authentication.UsernamePasswordAuthenticationToken;
import org.springframework.security.core.Authentication;
import org.springframework.security.core.context.SecurityContextHolder;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;
import org.springframework.web.util.UriComponentsBuilder;

import javax.servlet.http.HttpServletResponse;
import javax.validation.Valid;
import java.net.UnknownHostException;
import java.text.ParseException;

@RestController
@RequestMapping(value = "/auth")
public class AuthenticationController {

    @Autowired
    private TokenUtils tokenUtils;

    @Autowired
    private AuthenticationManager authenticationManager;

    @Autowired
    private UserRepository userRepository;

    @Autowired
    private CustomTokenService customTokenService;

    @Autowired
    private UserService userService;

    @PostMapping("/login")
    public ResponseEntity<LogUserDto> login(
            @RequestBody JwtAuthenticationRequestDto authenticationRequest, HttpServletResponse response) {

        Authentication authentication = authenticationManager.authenticate(new UsernamePasswordAuthenticationToken(
                authenticationRequest.getUsername(), authenticationRequest.getPassword()));
        SecurityContextHolder.getContext().setAuthentication(authentication);
        UserRole role = userService.findByUsername(authenticationRequest.getUsername()).getRole();
        UserDetails user = (UserDetails) authentication.getPrincipal();
        String jwt = tokenUtils.generateToken(user.getUser().getUsername());
        int expiresIn = tokenUtils.getExpiredIn();
        LogUserDto loggedUserDto = new LogUserDto(authenticationRequest.getUsername(), role.toString(), new UserTokenState(jwt, expiresIn));
        return ResponseEntity.ok(loggedUserDto);
    }






    @PostMapping("/register")
    public ResponseEntity<String> register(@Valid @RequestBody RegisterDto userRequest, UriComponentsBuilder ucBuilder) throws UnknownHostException, ParseException {
        User user = this.userService.findByUsername(userRequest.getUsername());
        User savedUser = userService.addUser(userRequest);
        if (savedUser != null) {
            return new ResponseEntity<>("SUCCESS!", HttpStatus.CREATED);
        }
        return new ResponseEntity<>("ERROR!", HttpStatus.INTERNAL_SERVER_ERROR);
    }
}
