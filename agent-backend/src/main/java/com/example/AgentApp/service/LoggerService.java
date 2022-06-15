package com.example.AgentApp.service;

import org.springframework.stereotype.Service;

@Service
public interface LoggerService {
    void loginSuccess(String email);

    void loginFailed(String email, String remoteAddr);

    void userSignedUp(String email, String remoteAddr);

    void userSigningUpFailed(String saving_new_user_failed, String email, String remoteAddr);

    void passwordChanged(String name, String remoteAddr);

    void accountConfirmed(String email, String remoteAddr);

    void accountConfirmedFailed(String token, String remoteAddr);

    void expiredMail(String email, String remoteAddr);

    void passwordChangingFailed(String message, String name, String remoteAddr);

    void resetPasswordSuccessfully(String email);

    void resetPasswordFailed(String email);

    void sendResetEmail(String email);

    void sendResetEmailFailed(String email);

    void checkCodeSuccessfully(String email);


    void checkCodeFailed(String email);

    void approveCompanyFailed(String name, Long id);

    void approveCompanySuccess(String name, Long id);

    void rejectCompanyFailed(String name, Long id);

    void rejectCompanySuccess(String name, Long id);

    void createCompanySuccess(String name);

    void createCompanyFailure(String name);

    void updateCompanyFailed(String name, Long id);

    void updateCompanySuccessfully(String name, Long id);


    void createOfferSuccess(String name, Long companyId);

    void createOfferFailed(String name, Long companyId);

    void createSalaryFailure(String name, Long companyId);

    void createSalarySuccess(String name, Long companyId);

    void createCommentFailure(String name, Long companyId);

    void createCommentSuccess(String name, Long companyId);

    void createInterviewSuccess(String name, Long companyId);

    void createInterviewFailure(String name, Long companyId);
}
