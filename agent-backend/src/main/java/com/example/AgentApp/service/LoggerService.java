package com.example.AgentApp.service;

import org.springframework.stereotype.Service;

@Service
public interface LoggerService {
    void loginSuccess(String email);
}
