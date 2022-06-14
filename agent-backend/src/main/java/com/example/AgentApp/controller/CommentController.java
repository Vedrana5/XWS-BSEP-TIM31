package com.example.AgentApp.controller;


import com.example.AgentApp.dto.CommentDto;
import com.example.AgentApp.dto.CreateCompanyDto;
import com.example.AgentApp.dto.InterviewDto;
import com.example.AgentApp.dto.SalaryDto;
import com.example.AgentApp.model.*;
import com.example.AgentApp.service.CommentService;
import com.example.AgentApp.service.CompanyService;
import com.example.AgentApp.service.LoggerService;
import com.example.AgentApp.service.impl.LoggerServiceImpl;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.security.access.prepost.PreAuthorize;
import org.springframework.security.core.context.SecurityContextHolder;
import org.springframework.web.bind.annotation.*;

import java.util.List;

@RestController
@RequestMapping(value = "/comments")
public class CommentController {

    @Autowired
    private CommentService commentService;

    @Autowired
    private CompanyService companyService;

    private  final LoggerService loggerService;

    public CommentController() {

        this.loggerService = new LoggerServiceImpl(this.getClass());
    }



   // @PreAuthorize("hasAuthority('CREATE_COMMENT_PERMISSION')")
    @CrossOrigin(origins = "https://localhost:4200")
    @PostMapping(value = "/createNewComment")
    public ResponseEntity<String> createCompany(@RequestBody CommentDto commentDto) {
        Comment comment = commentService.createComment(commentDto);
        if (comment != null) {
            loggerService.createCommentSuccess(SecurityContextHolder.getContext().getAuthentication().getName(),commentDto.getCompanyId());

            return new ResponseEntity<>("SUCCESS!", HttpStatus.CREATED);
        }
        loggerService.createCommentFailure(SecurityContextHolder.getContext().getAuthentication().getName(),commentDto.getCompanyId());

        return new ResponseEntity<>("ERROR!", HttpStatus.INTERNAL_SERVER_ERROR);
    }

   // @PreAuthorize("hasAuthority('CREATE_SALARY_PERMISSION')")
    @CrossOrigin(origins = "https://localhost:4200")
    @PostMapping(value = "/createSalary")
    public ResponseEntity<String> createSalary(@RequestBody SalaryDto salaryDto) {
        Salary salary = commentService.createSalary(salaryDto);
        if (salary != null) {
            loggerService.createSalarySuccess(SecurityContextHolder.getContext().getAuthentication().getName(),salaryDto.getCompanyId());

            return new ResponseEntity<>("SUCCESS!", HttpStatus.CREATED);
        }
        loggerService.createSalaryFailure(SecurityContextHolder.getContext().getAuthentication().getName(),salaryDto.getCompanyId());

        return new ResponseEntity<>("ERROR!", HttpStatus.INTERNAL_SERVER_ERROR);
    }

  //  @PreAuthorize("hasAuthority('CREATE_INTERVIEW_PERMISSION')")
    @CrossOrigin(origins = "https://localhost:4200")
    @PostMapping(value = "/createInterview")
    public ResponseEntity<String> createInterview(@RequestBody InterviewDto interviewDto) {
        Interview interview = commentService.createInterview(interviewDto);
        if (interview != null) {
            loggerService.createInterviewSuccess(SecurityContextHolder.getContext().getAuthentication().getName(),interviewDto.getCompanyId());
            return new ResponseEntity<>("SUCCESS!", HttpStatus.CREATED);
        }
        loggerService.createInterviewFailure(SecurityContextHolder.getContext().getAuthentication().getName(),interviewDto.getCompanyId());
        return new ResponseEntity<>("ERROR!", HttpStatus.INTERNAL_SERVER_ERROR);
    }




    @CrossOrigin(origins = "https://localhost:4200")
    @GetMapping("/getAllByCompany/{id}")
    public List<Comment> getCompanyComment(@PathVariable Long id){
        Company company = companyService.findById(id);
        List<Comment> comments =commentService.getAllByCompany(company.getId());

        return comments;
    }

    @CrossOrigin(origins = "https://localhost:4200")
    @GetMapping("/getAllSalaryByCompany/{id}")
    public List<Salary> getCompanySalary(@PathVariable Long id){
        Company company = companyService.findById(id);
        List<Salary> salaries =commentService.getAllSalaryByCompany(company.getId());

        return salaries;
    }


    @CrossOrigin(origins = "https://localhost:4200")
    @GetMapping("/getAllInterviewByCompany/{id}")
    public List<Interview> getCompanyInterview(@PathVariable Long id){
        Company company = companyService.findById(id);
        List<Interview> interviews =commentService.getAllInterviewByCompany(company.getId());

        return interviews;
    }

}
