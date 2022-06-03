package com.example.AgentApp.dto;

import lombok.AllArgsConstructor;
import lombok.Getter;
import lombok.NoArgsConstructor;
import lombok.Setter;

import javax.validation.constraints.Email;
import javax.validation.constraints.Pattern;

@Getter
@Setter
@NoArgsConstructor
@AllArgsConstructor
public class RegisterDto {

    private String username;

    @Pattern(regexp = "[a-zA-Z]+")
    private String firstName;

    @Pattern(regexp = "[a-zA-Z]+")
    private String lastName;

    @Pattern(regexp= "^(?=.*[a-z])(?=.*[A-Z])(?=.*[0-9])(?=.*[!\"#$@%&()*<>+_|~]).*$", message =  "Password format not valid")
    private String password;

    @Email
    private String email;

    @Email
    private String recoveryEmail;

    private String dateOfBirth;

}
