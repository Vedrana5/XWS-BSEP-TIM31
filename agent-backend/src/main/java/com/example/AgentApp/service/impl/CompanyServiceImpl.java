package com.example.AgentApp.service.impl;

import com.example.AgentApp.dto.CreateCompanyDto;
import com.example.AgentApp.model.*;
import com.example.AgentApp.repository.CompanyRepository;
import com.example.AgentApp.repository.UserRepository;
import com.example.AgentApp.service.CompanyService;
import com.example.AgentApp.service.UserService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.List;
import java.util.Optional;

@Service
public class CompanyServiceImpl implements CompanyService {


    @Autowired
    private CompanyRepository companyRepository;

    @Autowired
    private UserService userService;

    @Autowired
    private UserRepository userRepository;


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
       company.setCompanyStatus(CompanyStatus.PENDING);

       User owner=userService.findByUsername(companyDto.getOwnerUsername());
       company.setOwner(owner);

       companyRepository.save(company);
       return company;
    }

    @Override
    public Company approveCompany(Long id)
    {
       Optional<Company> company=companyRepository.findById(id);
       company.get().setCompanyStatus(CompanyStatus.APPROVED);
       User owner=company.get().getOwner();
       owner.setRole(UserRole.OWNER);
       userRepository.save(owner);
       companyRepository.save(company.get());

        return company.get();
    }

    @Override
    public Company rejectCompany(Long id) {
        Optional<Company> company=companyRepository.findById(id);
        company.get().setCompanyStatus(CompanyStatus.DECLINED);

        companyRepository.save(company.get());
        return company.get();
    }

    @Override
    public Company editCompany(CreateCompanyDto companyDto) {
        Optional<Company> company=companyRepository.findById(companyDto.getCompanyId());
        company.get().setName(companyDto.getName());
        company.get().setWebsite(companyDto.getWebsite());
        company.get().setEmail(companyDto.getEmail());
        company.get().setPhoneNumber(companyDto.getPhoneNumber());
        company.get().setCountryOfOrigin(companyDto.getCountryOfOrigin());
        company.get().setFounder(companyDto.getFounder());
        company.get().setNumberOfEmpl(companyDto.getNumberOfEmpl());

        return company.get();

    }



    @Override
    public List<Company> getAllStatusCompanies(CompanyStatus status) {
        List<Company> companies = companyRepository.findAllByCompanyStatus(status);
        return companies;
    }

    @Override
    public List<Company> getAll() {
        return companyRepository.findAll();
    }


}
