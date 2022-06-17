package com.example.AgentApp.service.impl;

import com.example.AgentApp.dto.CommentDto;
import com.example.AgentApp.dto.InterviewDto;
import com.example.AgentApp.dto.SalaryDto;
import com.example.AgentApp.model.*;
import com.example.AgentApp.repository.*;
import com.example.AgentApp.service.CommentService;
import com.example.AgentApp.service.CompanyService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.ArrayList;
import java.util.List;
import java.util.Optional;


@Service
public class CommentServiceImpl implements CommentService {

    @Autowired
    private UserRepository userRepository;

    @Autowired
    private CompanyService companyService;

    @Autowired
    private CompanyRepository companyRepository;

    @Autowired
    private CommentRepository commentRepository;

    @Autowired
    private SalaryRepository salaryRepository;

    @Autowired
    private InterviewRepository interviewRepository;


    @Override
    public Comment createComment(CommentDto commentDto) {

        Comment comment = new Comment();
        User u = userRepository.findByEmail(commentDto.getEmail());
        Optional<Company> company = companyRepository.findById(commentDto.companyId);
        comment.setCompany(company.get());
        comment.setUser(u);
        comment.setComment(commentDto.getComment());
        commentRepository.save(comment);


        return comment;
    }


    @Override
    public List<Comment> getAllByCompany(Long id) {
        List<Comment> myComments = new ArrayList<>();
        List<Comment> comments = commentRepository.findAll();
        for (Comment comment : comments) {
            if (comment.getCompany().getId() == id) {
                myComments.add(comment);
                return myComments;
            }
        }
        return myComments;
    }

    @Override
    public Salary createSalary(SalaryDto salaryDto) {

        Salary salary = new Salary();
        User u = userRepository.findByUsername(salaryDto.getUserName());
        Optional<Company> company = companyRepository.findById(salaryDto.companyId);
        salary.setCompany(company.get());
        salary.setUser(u);
        salary.setSalary(salaryDto.getSalary());
        salary.setPosition(salaryDto.getPosition());
        salaryRepository.save(salary);


        return salary;

    }

    @Override
    public List<Salary> getAllSalaryByCompany(Long id) {
        List<Salary> mySalary = new ArrayList<>();
        List<Salary> salaries = salaryRepository.findAll();
        for (Salary salary : salaries) {
            if (salary.getCompany().getId() == id) {
                mySalary.add(salary);
                return mySalary;
            }
        }
        return mySalary;
    }

    @Override
    public List<Interview> getAllInterviewByCompany(Long id) {
        List<Interview> myInterview = new ArrayList<>();
        List<Interview> interviews = interviewRepository.findAll();
        for (Interview interview : interviews) {
            if (interview.getCompany().getId() == id) {
                myInterview.add(interview);
                return myInterview;
            }
        }
        return myInterview;
    }

    @Override
    public Interview createInterview(InterviewDto interviewDto) {
        Interview interview = new Interview();
        User u = userRepository.findByUsername(interviewDto.getUserName());
        Optional<Company> company = companyRepository.findById(interviewDto.companyId);
        interview.setCompany(company.get());
        interview.setUser(u);
        interview.setComment(interviewDto.getComment());
        interview.setRating(interviewDto.getRating());
        if (interviewDto.difficulty.equals("HARD")) interview.setDifficulty(Interview.Difficulty.EASY.HARD);
        else if (interviewDto.difficulty.equals("INTERMEDIATE")) interview.setDifficulty(Interview.Difficulty.INTERMEDIATE);
        else interview.setDifficulty(Interview.Difficulty.INTERMEDIATE.EASY);
        interviewRepository.save(interview);
        return interview;
    }
}