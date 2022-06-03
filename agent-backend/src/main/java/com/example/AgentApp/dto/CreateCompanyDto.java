package com.example.AgentApp.dto;

import lombok.AllArgsConstructor;
import lombok.Getter;
import lombok.NoArgsConstructor;
import lombok.Setter;

@Getter
@Setter
@AllArgsConstructor
@NoArgsConstructor
public class CreateCompanyDto {
    private Long companyId;
    private String name;
    private String website;
    private String email;
    private String phoneNumber;
    private String countryOfOrigin;
    private String founder;
    private String numberOfEmpl;
    public String ownerUsername;

}
