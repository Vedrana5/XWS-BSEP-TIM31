package com.example.AgentApp.service;

import org.springframework.stereotype.Service;

@Service
public interface LoggerService {
    void loginSuccess(String email);

    void loginFailed(String email);

    void userSignedUp(String email);

    void userSigningUpFailed(String saving_new_user_failed, String email);

    void passwordChanged(String name);

    void accountConfirmed(String email);

    void accountConfirmedFailed(String token);

    void expiredMail(String email);

    void passwordChangingFailed(String message, String name);

    void resetPasswordSuccessfully(String email);

    void resetPasswordFailed(String email);

    void sendResetEmail(String email);

    void sendResetEmailFailed(String email);

    void checkCodeSuccessfully(String email);


    void checkCodeFailed(String email);
}
