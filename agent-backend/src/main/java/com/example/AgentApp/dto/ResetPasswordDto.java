package com.example.AgentApp.dto;

import lombok.AllArgsConstructor;
import lombok.Getter;
import lombok.NoArgsConstructor;
import lombok.Setter;

import javax.validation.constraints.Email;
import javax.validation.constraints.Pattern;

@Getter
@Setter
@AllArgsConstructor
@NoArgsConstructor
public class ResetPasswordDto {

    @Email
    private String email;


    @Pattern(regexp= "^(?=.*[a-z])(?=.*[A-Z])(?=.*[0-9])(?=.*[!\"#$@%&()*<>+_|~]).*$", message =  "Password format not valid")
    private String newPassword;
}
