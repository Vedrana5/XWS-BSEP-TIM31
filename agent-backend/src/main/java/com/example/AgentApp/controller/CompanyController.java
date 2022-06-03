package com.example.AgentApp.controller;

import com.example.AgentApp.dto.CreateCompanyDto;
import com.example.AgentApp.model.Company;
import com.example.AgentApp.model.User;
import com.example.AgentApp.service.CompanyService;
import com.example.AgentApp.service.impl.CompanyServiceImpl;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

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

    @GetMapping("approve/{id}")
    public ResponseEntity<?> approveCompany(@PathVariable Long id) {
        Company company = companyService.approveCompany(id);
      /*  List<Company> companies = companyService.getAllCompaniesWithStatus(CompanyStatus.PENDING);
        if (company != null){
            return new ResponseEntity<List<CompanyResponseDto>>(companyMapper.mapToDtos( companies), HttpStatus.OK);
        }
        return new ResponseEntity<>("Failed to approve company!", HttpStatus.CONFLICT);*/

        if (company != null) {
            return new ResponseEntity<>("SUCCESS!", HttpStatus.OK);
        }
        return new ResponseEntity<>("ERROR!", HttpStatus.INTERNAL_SERVER_ERROR);

    }
}
