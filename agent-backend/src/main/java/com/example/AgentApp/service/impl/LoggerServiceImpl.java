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
    public void loginSuccess(String email) {
        logger.info("Login successful. Email: " + email);

    }
}
