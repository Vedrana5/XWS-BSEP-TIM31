package com.example.AgentApp.service;

import com.example.AgentApp.dto.ChangePasswordDto;
import com.example.AgentApp.dto.RegisterDto;
import com.example.AgentApp.model.User;

import javax.management.relation.Relation;
import java.text.ParseException;

public interface UserService {
    User findByUsername(String username);

    User addUser(RegisterDto userRequest) throws ParseException;

    User activateAccount(User user);

    User findByEmail(String email);


    void resetPassword(String email, String newPassword);

    void changePassword(String emailFromToken, ChangePasswordDto changePasswordDto);
}
