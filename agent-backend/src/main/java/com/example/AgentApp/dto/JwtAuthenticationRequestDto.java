package com.example.AgentApp.dto;

import lombok.Getter;
import lombok.Setter;

@Getter
@Setter
public class JwtAuthenticationRequestDto {
    private String username;
    private String password;
}
