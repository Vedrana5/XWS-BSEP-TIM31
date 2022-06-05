package com.example.AgentApp.dto;

import lombok.AllArgsConstructor;
import lombok.Getter;
import lombok.NoArgsConstructor;
import lombok.Setter;

@Getter
@Setter

@NoArgsConstructor
public class LogUserDto {
    private String email;
    private String role;
    private UserTokenState token;

    public LogUserDto(String email, String role, UserTokenState token) {
        this.email = email;
        this.role = role;
        this.token = token;
    }

}
