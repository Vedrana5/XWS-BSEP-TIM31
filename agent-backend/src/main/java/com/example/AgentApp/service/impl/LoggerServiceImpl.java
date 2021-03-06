package com.example.AgentApp.service.impl;

import com.example.AgentApp.service.LoggerService;

import org.apache.logging.log4j.LogManager;
import org.apache.logging.log4j.Logger;

public class LoggerServiceImpl implements LoggerService {
    private final Logger logger;

    public LoggerServiceImpl(Class<?> parentClass) {
        this.logger = LogManager.getLogger(parentClass);
    }



    public void loginFailed(String email, String remoteAddr) {

        logger.warn("Login failed. Email: " + email,". Ip address: " + remoteAddr);
    }

    @Override
    public void userSignedUp(String email, String remoteAddr) {
        logger.info("New user successfully signed up. Email: {}", email);

    }

    @Override
    public void userSigningUpFailed(String saving_new_user_failed, String email, String remoteAddr) {
        logger.warn("New user signing up failed: {}. Email: {}", saving_new_user_failed, email);

    }

    @Override
    public void passwordChanged(String name, String remoteAddr) {
        logger.info("Password successfully changed. Email: {}", name);

    }

    @Override
    public void accountConfirmed(String email, String remoteAddr) {
        logger.info("Account confirmed successfully. Email: {}", email);
    }

    @Override
    public void accountConfirmedFailed(String token, String remoteAddr) {
        logger.warn("Failed to confirm account, token {} not found", token);
    }

    @Override
    public void expiredMail(String email, String remoteAddr) {
        logger.info("Email send successfully. Email: {}", email);

    }

    @Override
    public void passwordChangingFailed(String message, String name, String remoteAddr) {
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
        logger.warn("Code failed. Email: {}" + email);


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
    public void createOfferSuccess(String name, Long companyId) {
        logger.info("Offer successfully created for company {}. Email: {}",companyId, name);

    }

    @Override
    public void createOfferFailed(String name, Long companyId) {
        logger.warn("Offer creation failed for company {}. Email: {}",companyId, name);
    }

    @Override
    public void createSalaryFailure(String name, Long companyId) {
        logger.warn("Sallary creation failed for company {}. Email: {}",companyId, name);

    }

    @Override
    public void createSalarySuccess(String name, Long companyId) {
        logger.info("Sallary  successfully create for company {}. Email: {}",companyId, name);

    }

    @Override
    public void createCommentFailure(String name, Long companyId) {
        logger.warn("Comment creation failed for company {}. Email: {}",companyId, name);

    }

    @Override
    public void createCommentSuccess(String name, Long companyId) {
        logger.info("Comment  successfully create for company {}. Email: {}",companyId, name);

    }

    @Override
    public void createInterviewSuccess(String name, Long companyId) {
        logger.info("Interview  successfully create for company {}. Email: {}",companyId, name);

    }

    @Override
    public void createInterviewFailure(String name, Long companyId) {
        logger.warn("Interview creation failed for company {}. Email: {}",companyId, name);

    }

    @Override
    public void sendLinkForPasswordlessFailed(String email) {
        logger.warn("Send passwordless code failed. Email: " + email );

    }

    @Override
    public void sendLinkForPasswordlessSuccess(String email) {
        logger.info("Send passwordless code succeed. Email: " + email);

    }

    @Override
    public void passwordlessLoginSuccess(String username) {
        logger.info("Passwordless login succeed. Username: " + username);

    }

    @Override
    public void passwordlessLoginFailed(String username, String remoteAddr) {
        logger.error("Passwordless login failed. Username: " + username + ". Ip address: " + remoteAddr);
    }

    


    @Override
    public void loginSuccess(String email) {
        logger.info("Login successful. Email: " + email);

    }
}
