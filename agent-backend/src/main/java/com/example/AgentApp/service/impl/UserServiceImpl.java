package com.example.AgentApp.service.impl;

import com.example.AgentApp.dto.ChangePasswordDto;
import com.example.AgentApp.dto.RegisterDto;
import com.example.AgentApp.model.User;
import com.example.AgentApp.model.UserRole;
import com.example.AgentApp.repository.UserRepository;
import com.example.AgentApp.service.UserService;
import com.example.AgentApp.service.VerificationTokenService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.security.crypto.password.PasswordEncoder;
import org.springframework.stereotype.Service;

import java.text.ParseException;
import java.text.SimpleDateFormat;

@Service
public class UserServiceImpl implements UserService {

   @Autowired
   private UserRepository userRepository;



   @Autowired
   private PasswordEncoder passwordEncoder;

   @Autowired
   private VerificationTokenService verificationTokenService;

    @Override
    public User findByUsername(String username) {
        return userRepository.findByUsername(username);
    }

    @Override
    public User addUser(RegisterDto userRequest) throws ParseException {
        User user = new User();
        System.out.print(userRequest.getRecoveryEmail());
        System.out.print(userRequest.getUsername());
        user.setUsername(userRequest.getUsername());
        user.setFirstName(userRequest.getFirstName());
        user.setLastName(userRequest.getLastName());
        user.setDateOfBirth(new SimpleDateFormat("MM/dd/yyyy").parse(userRequest.getDateOfBirth()));
        user.setPassword(passwordEncoder.encode(userRequest.getPassword()));
        user.setEmail(userRequest.getEmail());
        user.setRecoveryEmail(userRequest.getRecoveryEmail());
        user.setPhoneNumber(userRequest.getPhoneNumber());
        user.setConfirmed(false);
        user.setRole(UserRole.REGISTERED_USER);
        userRepository.save(user);
        verificationTokenService.sendVerificationToken(user);
        return user;
    }

    @Override
    public User activateAccount(User user) {
        User userDb = findByEmail(user.getEmail());
        userDb.setConfirmed(true);
        User saved = userRepository.save(userDb);
        return saved;
    }

    public User findByEmail(String email) {
        return userRepository.findByEmail(email);
    }

    @Override
    public void resetPassword(String email, String newPassword) {
        User user = findByEmail(email);
        user.setPassword(passwordEncoder.encode(newPassword));
        userRepository.save(user);
    }

    @Override
    public void changePassword(String emailFromToken, ChangePasswordDto changePasswordDto) {
        User user = findByEmail(emailFromToken);
        if (passwordEncoder.matches(changePasswordDto.getCurrentPassword(),user.getPassword())) {
            user.setPassword(passwordEncoder.encode(changePasswordDto.getNewPassword()));
        }

        userRepository.save(user);
    }

    @Override
    public boolean check2FAStatus(String email) {
      /* return findByEmail(email).isUsing2FA();*/
        return true;

    }

    @Override
    public String change2FAStatus(String email, Boolean status) {
      /* User user = findByEmail(email);
        user.setUsing2FA(status);
        user.setSecret();
        userRepository.save(user);
        return status ? user.getSecret() : "";*/
        return null;

    }
}
