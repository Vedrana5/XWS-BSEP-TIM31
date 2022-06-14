package com.example.AgentApp.controller;

import com.example.AgentApp.dto.CreateCompanyDto;
import com.example.AgentApp.model.*;
import com.example.AgentApp.security.TokenUtilss;
import com.example.AgentApp.service.CompanyService;
import com.example.AgentApp.service.LoggerService;
import com.example.AgentApp.service.UserService;
import com.example.AgentApp.service.impl.CompanyServiceImpl;
import com.example.AgentApp.service.impl.LoggerServiceImpl;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.security.access.prepost.PreAuthorize;
import org.springframework.security.core.context.SecurityContextHolder;
import org.springframework.web.bind.annotation.*;

import javax.servlet.http.HttpServletRequest;
import java.util.List;

@RestController
@RequestMapping(value = "/companies")
public class CompanyController {

    @Autowired
    private CompanyService companyService;

    @Autowired
    private UserService userService;

    @Autowired
    private TokenUtilss tokenUtils;


    private  final LoggerService loggerService;

    public CompanyController() {

        this.loggerService = new LoggerServiceImpl(this.getClass());
    }


   // @PreAuthorize("hasAuthority('CREATE_COMPANY_PERMISSION')")
    @CrossOrigin(origins = "https://localhost:4200")
    @PostMapping(value = "/createNew")
    public ResponseEntity<String> createCompany(@RequestBody CreateCompanyDto companyDto) {
        Company company = companyService.createCompany(companyDto);
        if (company != null) {
            loggerService.createCompanySuccess(SecurityContextHolder.getContext().getAuthentication().getName());
            return new ResponseEntity<>("SUCCESS!", HttpStatus.CREATED);
        }
        loggerService.createCompanyFailure(SecurityContextHolder.getContext().getAuthentication().getName());
        return new ResponseEntity<>("ERROR!", HttpStatus.INTERNAL_SERVER_ERROR);
    }

   // @PreAuthorize("hasAuthority('ACTIVATE_COMPANY_PERMISSION')")
    @CrossOrigin(origins = "https://localhost:4200")
    @GetMapping(value = "approve/{id}")
    public Company approveCompany(@PathVariable Long id) {
        Company company = companyService.approveCompany(id);
        if(company==null) {
            loggerService.approveCompanyFailed(SecurityContextHolder.getContext().getAuthentication().getName(), id);
        }
        loggerService.approveCompanySuccess(SecurityContextHolder.getContext().getAuthentication().getName(), id);
        return company;

    }


  //  @PreAuthorize("hasAuthority('ACTIVATE_COMPANY_PERMISSION')")
    @CrossOrigin(origins = "https://localhost:4200")
    @GetMapping(value = "reject/{id}")
    public Company rejectCompany(@PathVariable Long id) {
        Company company = companyService.rejectCompany(id);
        if(company==null) {
            loggerService.rejectCompanyFailed(SecurityContextHolder.getContext().getAuthentication().getName(), id);
        }
        loggerService.rejectCompanySuccess(SecurityContextHolder.getContext().getAuthentication().getName(), id);


        return company;
    }

    @CrossOrigin(origins = "https://localhost:4200")
    @GetMapping("/getAllforUser")
    public List<Company> getAllForUser(HttpServletRequest request){
        String username = tokenUtils.getEmailFromToken(tokenUtils.getToken(request));
        User user = userService.findByEmail(username);

        List<Company> companies=companyService.getAllStatusCompanies(CompanyStatus.APPROVED);


      return companies;
    }

    @CrossOrigin(origins = "https://localhost:4200")
    @GetMapping("/getAllByUser/{email}")
    public List<Company> getUsersCompany(@PathVariable String email){
        User user = userService.findByEmail(email);
        List<Company>companies=null;
        companies=companyService.getAllByUser(user.getId());

         return companies;
    }

  //  @PreAuthorize("hasAuthority('UPDATE_COMPANY_PERMISSION')")
    @CrossOrigin(origins = "https://localhost:4200")
    @PutMapping(value = "/editCompany")
    public Company editCottage(@RequestBody CreateCompanyDto companyDto) {
        Company company = companyService.editCompany(companyDto);
       if(company==null) {
           loggerService.updateCompanyFailed(SecurityContextHolder.getContext().getAuthentication().getName(),companyDto.getId());
       }
        loggerService.updateCompanySuccessfully(SecurityContextHolder.getContext().getAuthentication().getName(),companyDto.getId());

        return company;

    }

    @CrossOrigin(origins = "https://localhost:4200")
    @GetMapping(value = "/pendingCompanies")
    public List<Company> getAllPendingCompanies(){
        List<Company> companies = companyService.getAllStatusCompanies(CompanyStatus.PENDING);
        return companies;
    }

   @CrossOrigin(origins = "https://localhost:4200")
    @GetMapping(value = "/approvedCompanies")
    public List<Company> getAllApprovedCompanies(){
        List<Company> companies = companyService.getAllStatusCompanies(CompanyStatus.APPROVED);
        return companies;
    }


    @GetMapping(value = "/all")
    public List<Company> getAll(){
        return companyService.getAll();
    }

    @CrossOrigin(origins = "https://localhost:4200")
    @GetMapping(value = "findCompany/{id}")
    public Company findCompanybyId(@PathVariable Long id) {
        Company company = companyService.findById(id);
        return company;

    }

}
