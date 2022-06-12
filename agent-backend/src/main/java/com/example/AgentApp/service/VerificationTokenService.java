package com.example.AgentApp.service;

import com.example.AgentApp.model.CustomToken;
import com.example.AgentApp.model.User;

public interface VerificationTokenService {
    void sendVerificationToken(User user);

    CustomToken findByToken(String token);

    void deleteById(Long id);

    void sendRecoveryMail(User user);

    CustomToken findByUser(User user);

    boolean checkResetPasswordCode(String code, String token);
}
