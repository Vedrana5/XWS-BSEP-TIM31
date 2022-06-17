package com.example.AgentApp.dto;

import lombok.Getter;
import lombok.Setter;

@Getter
@Setter
public class JwtAuthenticationRequestDto {
    private String email;
    private String password;
    private String code;
}
