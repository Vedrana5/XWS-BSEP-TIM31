package com.example.AgentApp.dto;

import com.example.AgentApp.model.Company;
import lombok.AllArgsConstructor;
import lombok.Getter;
import lombok.NoArgsConstructor;
import lombok.Setter;

@Getter
@Setter
@AllArgsConstructor
@NoArgsConstructor
public class CreateCompanyDto {
    private Long id;
    private String name;
    private String website;
    private String email;
    private String phoneNumber;
    private String countryOfOrigin;
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
