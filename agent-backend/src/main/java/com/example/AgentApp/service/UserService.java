package com.example.AgentApp.service;

import com.example.AgentApp.dto.RegisterDto;
import com.example.AgentApp.model.User;

import javax.management.relation.Relation;
import java.text.ParseException;

public interface UserService {
    User findByUsername(String username);

    User addUser(RegisterDto userRequest) throws ParseException;
}
