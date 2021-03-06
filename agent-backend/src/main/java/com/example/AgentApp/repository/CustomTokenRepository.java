package com.example.AgentApp.repository;

import com.example.AgentApp.model.Company;
import com.example.AgentApp.model.CustomToken;
import com.example.AgentApp.model.User;
import org.springframework.data.jpa.repository.JpaRepository;

public interface CustomTokenRepository  extends JpaRepository<CustomToken, Long> {
    CustomToken findByToken(String token);

    CustomToken findByUser(User user);
}
