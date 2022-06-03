package com.example.AgentApp.service;

import com.example.AgentApp.model.User;

import javax.management.relation.Relation;

public interface UserService {
    User findByUsername(String username);
}
