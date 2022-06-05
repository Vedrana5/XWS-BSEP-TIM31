package com.example.AgentApp.service.impl;

import com.example.AgentApp.model.CustomToken;
import com.example.AgentApp.model.TokenType;
import com.example.AgentApp.model.User;
import com.example.AgentApp.repository.CustomTokenRepository;
import com.example.AgentApp.service.EmailSenderService;
import com.example.AgentApp.service.VerificationTokenService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.UUID;

@Service
public class VerificationTokenServiceImpl implements VerificationTokenService {

    @Autowired
    private CustomTokenRepository customTokenRepository;

    @Autowired
    private EmailSenderService emailSenderService;

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

    private CustomToken createConfirmationToken(User user) {
        CustomToken token = new CustomToken(UUID.randomUUID().toString(),user, TokenType.Confirmation);
        customTokenRepository.save(token);
        return token;

    }

}
