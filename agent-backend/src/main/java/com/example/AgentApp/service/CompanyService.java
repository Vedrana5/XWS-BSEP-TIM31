package com.example.AgentApp.service;

import com.example.AgentApp.dto.CreateCompanyDto;
import com.example.AgentApp.model.Company;
import com.example.AgentApp.model.CompanyStatus;
import com.example.AgentApp.model.Offer;
import org.springframework.stereotype.Service;

import java.util.List;

public interface CompanyService {
    Company createCompany(CreateCompanyDto companyDto);

    Company approveCompany(Long id);


    Company rejectCompany(Long id);

    Company editCompany(CreateCompanyDto companyDto);


    List<Company> getAllStatusCompanies(CompanyStatus status);

    List<Company> getAll();
}
