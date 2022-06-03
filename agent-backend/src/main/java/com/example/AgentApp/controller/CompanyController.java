package com.example.AgentApp.controller;

import com.example.AgentApp.dto.CreateCompanyDto;
import com.example.AgentApp.model.Company;
import com.example.AgentApp.model.User;
import com.example.AgentApp.service.CompanyService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
@RequestMapping(value = "/companies")
public class CompanyController {

    @Autowired
    private CompanyService companyService;


    @PostMapping("/createNew")
    public ResponseEntity<String> createCompany(@RequestBody CreateCompanyDto companyDto){
        Company company = companyService.createCompany(companyDto);
        if (company != null) {
            return new ResponseEntity<>("SUCCESS!", HttpStatus.CREATED);
        }
        return new ResponseEntity<>("ERROR!", HttpStatus.INTERNAL_SERVER_ERROR);
    }
}
