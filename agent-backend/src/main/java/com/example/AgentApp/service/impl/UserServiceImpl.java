package com.example.AgentApp.service.impl;

import com.example.AgentApp.dto.RegisterDto;
import com.example.AgentApp.model.User;
import com.example.AgentApp.model.UserRole;
import com.example.AgentApp.repository.UserRepository;
import com.example.AgentApp.service.CustomTokenService;
import com.example.AgentApp.service.UserService;
import net.bytebuddy.utility.nullability.AlwaysNull;
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
   private CustomTokenService customTokenService;

   @Autowired
   private PasswordEncoder passwordEncoder;

    @Override
    public User findByUsername(String username) {
        return userRepository.findByUsername(username);
    }

    @Override
    public User addUser(RegisterDto userRequest) throws ParseException {
        User user = new User();
        user.setUsername(userRequest.getUsername());
        user.setFirstName(userRequest.getFirstName());
        user.setLastName(userRequest.getLastName());
        user.setDateOfBirth(new SimpleDateFormat("MM/dd/yyyy").parse(userRequest.getDateOfBirth()));
        user.setPassword(passwordEncoder.encode(userRequest.getPassword()));
        user.setEmail(userRequest.getEmail());
        user.setRecoveryEmail(userRequest.getRecoveryEmail());
        user.setConfirmed(false);
        user.setRole(UserRole.REGISTERED_USER);
        userRepository.save(user);
        return user;
    }
}
