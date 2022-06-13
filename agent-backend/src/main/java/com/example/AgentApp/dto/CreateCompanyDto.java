package com.example.AgentApp.dto;

import com.example.AgentApp.model.Company;
import lombok.AllArgsConstructor;
import lombok.Getter;
import lombok.NoArgsConstructor;
import lombok.Setter;

import javax.validation.constraints.Email;
import javax.validation.constraints.NotBlank;
import javax.validation.constraints.Pattern;

@Getter
@Setter
@AllArgsConstructor
@NoArgsConstructor
public class CreateCompanyDto {

    private Long id;

    @Pattern(regexp = "[a-zA-Z]+")
    private String name;


    private String website;

    @Email
    private String email;

    private String phoneNumber;

    @Pattern(regexp = "[a-zA-Z]+")
    private String countryOfOrigin;

    @Pattern(regexp = "[a-zA-Z]+")
    private String founder;


    private String numberOfEmpl;

    public String ownerUsername;

    public CreateCompanyDto(Company c) {
      this.id=c.getId();
      this.name=c.getName();
      this.website=c.getWebsite();
      this.email=c.getEmail();
      this.phoneNumber=c.getPhoneNumber();
      this.countryOfOrigin=c.getCountryOfOrigin();
      this.founder=c.getFounder();
      this.numberOfEmpl=c.getNumberOfEmpl();
      this.ownerUsername=Long.toString(c.getOwner().getId());

    }

}
