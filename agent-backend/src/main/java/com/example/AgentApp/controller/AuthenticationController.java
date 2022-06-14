package com.example.AgentApp.controller;

import com.example.AgentApp.dto.*;
import com.example.AgentApp.model.CustomToken;
import com.example.AgentApp.model.User;
import com.example.AgentApp.model.UserDetails;
import com.example.AgentApp.model.UserRole;
import com.example.AgentApp.repository.UserRepository;

import com.example.AgentApp.security.TokenUtilss;
import com.example.AgentApp.service.LoggerService;
import com.example.AgentApp.service.UserService;
import com.example.AgentApp.service.VerificationTokenService;
import com.example.AgentApp.service.impl.LoggerServiceImpl;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.security.authentication.AuthenticationManager;
import org.springframework.security.authentication.UsernamePasswordAuthenticationToken;
import org.springframework.security.core.Authentication;
import org.springframework.security.core.context.SecurityContextHolder;
import org.springframework.web.bind.annotation.*;
import org.springframework.web.util.UriComponentsBuilder;

import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;
import javax.validation.Valid;
import java.net.URI;
import java.net.UnknownHostException;
import java.text.ParseException;
import java.time.LocalDateTime;

@RestController
@RequestMapping(value = "/auth")
@CrossOrigin(origins = "https://localhost:4200")
public class AuthenticationController {

    @Autowired
    private TokenUtilss tokenUtils;

    @Autowired
    private AuthenticationManager authenticationManager;

    @Autowired
    private UserRepository userRepository;

    @Autowired
    private VerificationTokenService customTokenService;

    @Autowired
    private UserService userService;


    private  final LoggerService loggerService;

    public AuthenticationController() {

        this.loggerService = new LoggerServiceImpl(this.getClass());
    }

    @CrossOrigin(origins = "https://localhost:4200")
    @PostMapping("/login")
    public ResponseEntity<LogUserDto> login(
            @RequestBody JwtAuthenticationRequestDto authenticationRequest, HttpServletResponse response) {

        try {
            Authentication authentication = authenticationManager.authenticate(new UsernamePasswordAuthenticationToken(
                    authenticationRequest.getEmail(), authenticationRequest.getPassword()));
            SecurityContextHolder.getContext().setAuthentication(authentication);
            UserRole role = userService.findByEmail(authenticationRequest.getEmail()).getRole();
            UserDetails user = (UserDetails) authentication.getPrincipal();
            String jwt = tokenUtils.generateToken(user.getUser().getEmail());
            int expiresIn = tokenUtils.getExpiredIn();
            LogUserDto loggedUserDto = new LogUserDto(authenticationRequest.getEmail(), role.toString(), new UserTokenState(jwt, expiresIn));
            loggerService.loginSuccess(authenticationRequest.getEmail());
            return ResponseEntity.ok(loggedUserDto);
        }
        catch (Exception ex) {
            loggerService.loginFailed(authenticationRequest.getEmail());
            return ResponseEntity.badRequest().build();
        }

    }





    @CrossOrigin(origins = "https://localhost:4200")
    @PostMapping("/register")
    public ResponseEntity<String> register(@Valid @RequestBody RegisterDto userRequest, UriComponentsBuilder ucBuilder) throws UnknownHostException, ParseException {
        User savedUser = userService.addUser(userRequest);
        if (savedUser != null) {
            loggerService.userSignedUp(userRequest.getEmail());
            return new ResponseEntity<>("SUCCESS!", HttpStatus.CREATED);
        }
        loggerService.userSigningUpFailed("Saving new user failed", userRequest.getEmail());
        return new ResponseEntity<>("ERROR!", HttpStatus.INTERNAL_SERVER_ERROR);
    }

    @CrossOrigin(origins = "https://localhost:4200")
    @GetMapping("/confirmAccount/{token}")
    public ResponseEntity<String> confirmAccount(@PathVariable String token) {
        CustomToken verificationToken = customTokenService.findByToken(token);
        User user = verificationToken.getUser();
        if (verificationToken.getExpiryDate().isBefore(LocalDateTime.now())) {
            customTokenService.deleteById(verificationToken.getId());
            customTokenService.sendVerificationToken(user);
            loggerService.expiredMail(user.getEmail());
            return new ResponseEntity<>("Confirmation link is expired,we sent you new one.Please check you mail box.", HttpStatus.BAD_REQUEST);
        }
        User activated = userService.activateAccount(user);
        customTokenService.deleteById(verificationToken.getId());
        if (activated.isConfirmed()) {
            loggerService.accountConfirmed(user.getEmail());
            return new ResponseEntity<>("Account is activated.You can login.", HttpStatus.OK);


        } else {
            loggerService.accountConfirmedFailed(token);
            return new ResponseEntity<>("Error happened!", HttpStatus.INTERNAL_SERVER_ERROR);
        }
    }

    @CrossOrigin(origins = "https://localhost:4200")
    @PostMapping(value = "/sendCode")
    public ResponseEntity<?> sendCode(@RequestBody String email) {
        User user = userService.findByEmail(email);
        if (user == null){
            loggerService.sendResetEmailFailed(user.getEmail());
            return ResponseEntity.notFound().build();}
        customTokenService.sendResetPasswordToken(user);
        loggerService.sendResetEmail(user.getEmail());
        return ResponseEntity.accepted().build();
    }


    @CrossOrigin(origins = "https://localhost:4200")
    @PutMapping(value = "/changePassword")
    public ResponseEntity<HttpStatus> changePassword(@Valid @RequestBody ChangePasswordDto changePasswordDto, HttpServletRequest request) {
        String token = tokenUtils.getToken(request);

       try {
           userService.changePassword(tokenUtils.getEmailFromToken(token), changePasswordDto);
           loggerService.passwordChanged(SecurityContextHolder.getContext().getAuthentication().getName());

           return ResponseEntity.noContent().build();
       } catch(Exception e) {
           loggerService.passwordChangingFailed(e.getMessage(), SecurityContextHolder.getContext().getAuthentication().getName());
           return ResponseEntity.badRequest().build();
       }
    }




    @CrossOrigin(origins = "https://localhost:4200")
    @PostMapping(value = "/checkCode")
    public ResponseEntity<String> checkCode(@RequestBody CheckCodeDto checkCodeDto) {
        User user = userService.findByEmail(checkCodeDto.getEmail());
        CustomToken token = customTokenService.findByUser(user);

        if (customTokenService.checkResetPasswordCode(checkCodeDto.getCode(), token.getToken())) {
            if (token.getExpiryDate().isBefore(LocalDateTime.now())) {
                customTokenService.deleteById(token.getId());
                customTokenService.sendVerificationToken(user);
                return new ResponseEntity<>("Reset password code is expired,we sent you new one.Please check you mail box.", HttpStatus.BAD_REQUEST);
            }

            customTokenService.deleteById(token.getId());
            loggerService.checkCodeSuccessfully(user.getEmail());
            return new ResponseEntity<>("Success!", HttpStatus.OK);
        }

        loggerService.checkCodeFailed(user.getEmail());
        return new ResponseEntity<>("Entered code is not valid!", HttpStatus.BAD_REQUEST);
    }

    @CrossOrigin(origins = "https://localhost:4200")
    @PostMapping(value = "/resetPassword")
    public ResponseEntity<String> resetPassword(@Valid @RequestBody ResetPasswordDto resetPasswordDto) {
       try {
           userService.resetPassword(resetPasswordDto.getEmail(), resetPasswordDto.getNewPassword());
           loggerService.resetPasswordSuccessfully(resetPasswordDto.getEmail());
           return new ResponseEntity<>("OK", HttpStatus.OK);
       }
       catch(Exception e) {
           loggerService.resetPasswordFailed(resetPasswordDto.getEmail());
           return new ResponseEntity<>("Not a reset password", HttpStatus.BAD_REQUEST);


       }
    }

}
