package com.example.AgentApp.service.impl;

import com.example.AgentApp.dto.CreateCompanyDto;
import com.example.AgentApp.model.Company;
import com.example.AgentApp.model.User;
import com.example.AgentApp.repository.CompanyRepository;
import com.example.AgentApp.service.CompanyService;
import com.example.AgentApp.service.UserService;
import org.springframework.beans.factory.annotation.Autowired;

public class CompanyServiceImpl implements CompanyService {


    @Autowired
    private CompanyRepository companyRepository;

    @Autowired
    private UserService userService;


    @Override
    public Company createCompany(CreateCompanyDto companyDto) {

       Company company=new Company();
       
       company.setName(companyDto.getName());
       company.setWebsite(companyDto.getWebsite());
       company.setEmail(companyDto.getEmail());
       company.setPhoneNumber(companyDto.getPhoneNumber());
       company.setCountryOfOrigin(companyDto.getCountryOfOrigin());
       company.setFounder(companyDto.getFounder());
       company.setNumberOfEmpl(companyDto.getNumberOfEmpl());
       company.setStatus(Company.StatusofCompany.PENDING);

       User owner=userService.findByUsername(companyDto.getOwnerUsername());
       company.setOwner(owner);

       companyRepository.save(company);
       return company;
    }
}
