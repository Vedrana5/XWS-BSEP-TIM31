package com.example.AgentApp.service;

import com.example.AgentApp.dto.CreateCompanyDto;
import com.example.AgentApp.model.Company;

public interface CompanyService {
    Company createCompany(CreateCompanyDto companyDto);

    Company approveCompany(Long id);


    Company rejectCompany(Long id);

    Company editCompany(CreateCompanyDto companyDto);
}
