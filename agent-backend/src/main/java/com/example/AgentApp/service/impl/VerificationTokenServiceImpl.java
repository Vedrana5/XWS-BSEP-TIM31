package com.example.AgentApp.service.impl;

import com.example.AgentApp.model.CustomToken;
import com.example.AgentApp.model.TokenType;
import com.example.AgentApp.model.User;
import com.example.AgentApp.repository.CustomTokenRepository;
import com.example.AgentApp.service.EmailSenderService;
import com.example.AgentApp.service.VerificationTokenService;
import net.bytebuddy.utility.RandomString;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.security.crypto.password.PasswordEncoder;
import org.springframework.stereotype.Service;

import java.time.LocalDateTime;
import java.util.UUID;

@Service
public class VerificationTokenServiceImpl implements VerificationTokenService {

    @Autowired
    private CustomTokenRepository customTokenRepository;

    @Autowired
    private EmailSenderService emailSenderService;

    @Autowired
    private PasswordEncoder passwordEncoder;
    @Override

    public void sendVerificationToken(User user) {
        String confirmationLink = "http://localhost:8082/auth/confirmAccount/" + createConfirmationToken(user).getToken();
        emailSenderService.sendEmail(user.getEmail(),"Confirm account", "Click on following link to confirm " +
                "your account \n" + confirmationLink);
    }

    @Override
    public CustomToken findByToken(String token) {
        return customTokenRepository.findByToken(token);
    }

    @Override
    public void deleteById(Long id) {
        customTokenRepository.deleteById(id);

    }

    @Override
    public void sendRecoveryMail(User user) {
        CustomToken customToken = createResetPasswordToken(user);
        String passwordCode = customToken.getToken();
        saveToken(customToken);
        emailSenderService.sendEmail(user.getRecoveryEmail(),"Reset password", "Following code is your new temporary " +
                "password \nCode : " + passwordCode);

    }

    @Override
    public CustomToken findByUser(User user) {
        return customTokenRepository.findByUser(user);
    }

    @Override
    public boolean checkResetPasswordCode(String code, String token) {
        return passwordEncoder.matches(code,token);
    }

    @Override
    public void sendResetPasswordToken(User user) {
        CustomToken customToken = createResetPasswordToken(user);
        String passwordCode = customToken.getToken();
        System.out.print("fsdsfdsfsdfds"+passwordCode);
        saveToken(customToken);
        emailSenderService.sendEmail(user.getRecoveryEmail(),"Reset password", "Following code is your new temporary " +
                "password \nCode : " + passwordCode);

    }

    @Override
    public void sendMagicLink(User user) {
      
        CustomToken token = createTokenForMagicLink(user);
        System.out.print("Jebeni token"+token.getToken());
        emailSenderService.sendEmail(user.getEmail(),"Password-less login",
                "Click on the following link to sign in to your account "
                        +"https://localhost:4200/passwordless-login/"
        + token.getToken()
        );
    }

    private CustomToken createTokenForMagicLink(User user) {
        CustomToken token = new CustomToken(UUID.randomUUID().toString(),user, TokenType.Verification);
        token.setExpiryDate(LocalDateTime.now().plusMinutes(5));
        customTokenRepository.save(token);
        return token;
    }


    private void saveToken(CustomToken customToken) {
        String valueOfToken = customToken.getToken();
        customToken.setToken(passwordEncoder.encode(valueOfToken));
        customTokenRepository.save(customToken);
    }

    private CustomToken createResetPasswordToken(User user) {
        CustomToken token = new CustomToken(RandomString.make(8),user,TokenType.ResetPassword);
        token.setExpiryDate(LocalDateTime.now().plusMinutes(10));
        customTokenRepository.save(token);
        return token;
    }

    private CustomToken createConfirmationToken(User user) {
        CustomToken token = new CustomToken(UUID.randomUUID().toString(),user, TokenType.Confirmation);
        customTokenRepository.save(token);
        return token;

    }

}
