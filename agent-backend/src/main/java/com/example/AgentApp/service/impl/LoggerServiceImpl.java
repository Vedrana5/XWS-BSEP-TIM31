package com.example.AgentApp.service.impl;

import com.example.AgentApp.service.LoggerService;

import org.apache.logging.log4j.LogManager;
import org.apache.logging.log4j.Logger;

public class LoggerServiceImpl implements LoggerService {
    private final Logger logger;

    public LoggerServiceImpl(Class<?> parentClass) {
        this.logger = LogManager.getLogger(parentClass);
    }



    public void loginFailed(String email) {

        logger.warn("Login failed. Email: " + email);
    }

    @Override
    public void userSignedUp(String email) {
        logger.info("New user successfully signed up. Email: {}", email);

    }

    @Override
    public void userSigningUpFailed(String saving_new_user_failed, String email) {
        logger.warn("New user signing up failed: {}. Email: {}", saving_new_user_failed, email);

    }

    @Override
    public void passwordChanged(String name) {
        logger.info("Password successfully changed. Email: {}", name);

    }

    @Override
    public void accountConfirmed(String email) {
        logger.info("Account confirmed successfully. Email: {}", email);
    }

    @Override
    public void accountConfirmedFailed(String token) {
        logger.warn("Failed to confirm account, token {} not found", token);
    }

    @Override
    public void expiredMail(String email) {
        logger.info("Email send successfully. Email: {}", email);

    }

    @Override
    public void passwordChangingFailed(String message, String name) {
        logger.warn("Change password failed. Email: {}.", name);
    }

    @Override
    public void resetPasswordSuccessfully(String email) {
        logger.info("Reset password successfully. Email: {}", email);
    }

    @Override
    public void resetPasswordFailed(String email) {
        logger.info("Reset password failed. Email: {}.", email);

    }

    @Override
    public void sendResetEmail(String email) {
        logger.info("Send mail for reset password. Email: {}", email);

    }

    @Override
    public void sendResetEmailFailed(String email) {
        logger.warn("User not found. Email: " + email);

    }

    @Override
    public void checkCodeSuccessfully(String email) {
        logger.info("Code accepted. Email: {}", email);

    }

    @Override
    public void checkCodeFailed(String email) {
        logger.warn("Code failed. Email: " + email);


    }

    @Override
    public void approveCompanyFailed(String name, Long id) {
        logger.warn("Company {} unsuccessfully activated by {}", id, name);
    }

    @Override
    public void approveCompanySuccess(String name, Long id) {
        logger.info("Company {} successfully activated by {}", id, name);
    }

    @Override
    public void rejectCompanyFailed(String name, Long id) {
        logger.warn("Company {} unsuccessfully reject by {}", id, name);

    }

    @Override
    public void rejectCompanySuccess(String name, Long id) {
        logger.info("Company {} successfully activated by {}", id, name);

    }

    @Override
    public void createCompanySuccess(String name) {
        logger.info("Company request successfully created. Email: {}", name);
    }

    @Override
    public void createCompanyFailure(String name) {
        logger.warn("Company request creation failed. Email: {}", name);
    }

    @Override
    public void updateCompanyFailed(String name, Long id) {
        logger.info("Company {} update failed. Email: {}",id, name);

    }

    @Override
    public void updateCompanySuccessfully(String name, Long id) {
        logger.info("Company {}  successfully update. Email: {}",id, name);

    }


    @Override
    public void loginSuccess(String email) {
        logger.info("Login successful. Email: " + email);

    }
}
