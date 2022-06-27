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
import de.taimos.totp.TOTP;
import org.apache.commons.codec.binary.Base32;
import org.apache.commons.codec.binary.Hex;
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
    public ResponseEntity<Object> login(@Valid
            @RequestBody JwtAuthenticationRequestDto authenticationRequest, HttpServletRequest request) {

        try {
            Authentication authentication = authenticationManager.authenticate(new UsernamePasswordAuthenticationToken(
                    authenticationRequest.getEmail(), authenticationRequest.getPassword()));

            User u = userService.findByEmail(authenticationRequest.getEmail());
            String code = authenticationRequest.getCode();
            if (u.isUsingFa() && (code == null || !code.equals(getTOTPCode(u.getSecret())))) {
                loggerService.loginFailed(authenticationRequest.getEmail(),request.getRemoteAddr());
                return ResponseEntity.badRequest().body("Code invalid");
            }


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
            loggerService.loginFailed(authenticationRequest.getEmail(),request.getRemoteAddr());
            return ResponseEntity.badRequest().build();
        }

    }


    public static String getTOTPCode(String secretKey) {
        Base32 base32 = new Base32();
        byte[] bytes = base32.decode(secretKey);
        String hexKey = Hex.encodeHexString(bytes);
        return TOTP.getOTP(hexKey);
    }


    @CrossOrigin(origins = "https://localhost:4200")
    @PostMapping("/register")
    public ResponseEntity<String> register(@Valid @RequestBody RegisterDto userRequest, UriComponentsBuilder ucBuilder,HttpServletRequest request) throws UnknownHostException, ParseException {
        User savedUser = userService.addUser(userRequest);
        if (savedUser != null) {
            loggerService.userSignedUp(userRequest.getEmail(),request.getRemoteAddr());
            return new ResponseEntity<>("SUCCESS!", HttpStatus.CREATED);
        }
        loggerService.userSigningUpFailed("Saving new user failed", userRequest.getEmail(),request.getRemoteAddr());
        return new ResponseEntity<>("ERROR!", HttpStatus.INTERNAL_SERVER_ERROR);
    }

    @CrossOrigin(origins = "https://localhost:4200")
    @GetMapping("/confirmAccount/{token}")
    public ResponseEntity<String> confirmAccount(@PathVariable String token,HttpServletRequest request) {
        CustomToken verificationToken = customTokenService.findByToken(token);
        System.out.print(verificationToken);
        User user = verificationToken.getUser();
        if (verificationToken.getExpiryDate().isBefore(LocalDateTime.now())) {
            customTokenService.deleteById(verificationToken.getId());
            customTokenService.sendVerificationToken(user);
            loggerService.expiredMail(user.getEmail(),request.getRemoteAddr());
            return new ResponseEntity<>("Confirmation link is expired,we sent you new one.Please check you mail box.", HttpStatus.BAD_REQUEST);
        }
        User activated = userService.activateAccount(user);
        customTokenService.deleteById(verificationToken.getId());
        if (activated.isConfirmed()) {
            loggerService.accountConfirmed(user.getEmail(),request.getRemoteAddr());
            return new ResponseEntity<>("Account is activated.You can login.", HttpStatus.OK);


        } else {
            loggerService.accountConfirmedFailed(user.getEmail(),request.getRemoteAddr());
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
           loggerService.passwordChanged(SecurityContextHolder.getContext().getAuthentication().getName(),request.getRemoteAddr());

           return ResponseEntity.noContent().build();
       } catch(Exception e) {
           loggerService.passwordChangingFailed(e.getMessage(), SecurityContextHolder.getContext().getAuthentication().getName(),request.getRemoteAddr());
           return ResponseEntity.badRequest().build();
       }
    }




    @CrossOrigin(origins = "https://localhost:4200")
    @PostMapping(value = "/checkCode")
    public ResponseEntity<String> checkCode(@Valid @RequestBody CheckCodeDto checkCodeDto) {
        User user = userService.findByEmail(checkCodeDto.getEmail());
        CustomToken token = customTokenService.findByUser(user);

        if (customTokenService.checkResetPasswordCode(checkCodeDto.getCode(), token.getToken())) {
           System.out.print("Kodddd"+ checkCodeDto.getCode());
          //  customTokenService.deleteById(token.getId());
            loggerService.checkCodeSuccessfully(user.getEmail());
            return new ResponseEntity<>("\" Success \"", HttpStatus.OK);
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


    @PostMapping(value = "/password-less-login")
    public ResponseEntity<?> sendLinkForPasswordLess(@RequestBody String email) {
        User user = userService.findByEmail(email);
        System.out.print("fdfdgdg"+user);
        if (user == null) {
            loggerService.sendLinkForPasswordlessFailed(user.getEmail());
            return ResponseEntity.notFound().build();}
        customTokenService.sendMagicLink(user);
        loggerService.sendLinkForPasswordlessSuccess(user.getEmail());
        return ResponseEntity.accepted().build();
    }

    @GetMapping(value = "/password-less-login/{link}")
    public ResponseEntity<?> passwordLessLogin(@PathVariable String link,HttpServletRequest request) {
        CustomToken token  = customTokenService.findByToken(link);
        User user = token.getUser();
        if (token.getExpiryDate().isBefore(LocalDateTime.now())) {
            customTokenService.deleteById(token.getId());
            customTokenService.sendMagicLink(user);
            loggerService.passwordlessLoginFailed(user.getUsername(),request.getRemoteAddr());
            return new ResponseEntity<>("Your magic link is expired,we sent you new one. Please check you mail box.", HttpStatus.BAD_REQUEST);
        }
        Authentication authentication = new UsernamePasswordAuthenticationToken(
                user.getUsername(), null);
        SecurityContextHolder.getContext().setAuthentication(authentication);
        UserRole role = userService.findByUsername(user.getUsername()).getRole();
        String jwt = tokenUtils.generateToken(user.getUsername());
        int expiresIn = tokenUtils.getExpiredIn();
        LogUserDto loggedUserDto = new LogUserDto(user.getUsername(), role.toString(), new UserTokenState(jwt, expiresIn));
        customTokenService.deleteById(token.getId());
        loggerService.passwordlessLoginSuccess(user.getUsername());
        return ResponseEntity.ok(loggedUserDto);
    }

    @PutMapping(value = "/two-factor-auth")
    public ResponseEntity<SecretDto> change2FAStatus(@RequestBody Change2FAStatusDto dto) {
        String secret = userService.change2FAStatus(dto.email, dto.status);
        return ResponseEntity.ok(new SecretDto(secret));
    }
    @GetMapping(value= "/two-factor-auth-status/{email}")
    public ResponseEntity<Boolean> check2FAStatus(@PathVariable String email, HttpServletRequest request) {
        boolean twoFAEnabled = userService.check2FAStatus(email);
        return ResponseEntity.ok(twoFAEnabled);
    }

}
