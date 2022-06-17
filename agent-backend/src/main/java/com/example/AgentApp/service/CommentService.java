package com.example.AgentApp.service;

import com.example.AgentApp.dto.CommentDto;
import com.example.AgentApp.dto.InterviewDto;
import com.example.AgentApp.dto.SalaryDto;
import com.example.AgentApp.model.Comment;
import com.example.AgentApp.model.Interview;
import com.example.AgentApp.model.Salary;

import java.util.List;

public interface CommentService {
    Comment createComment(CommentDto commentDto);

    List<Comment> getAllByCompany(Long id);

    Salary createSalary(SalaryDto salaryDto);

    List<Salary> getAllSalaryByCompany(Long id);

    List<Interview> getAllInterviewByCompany(Long id);

    Interview createInterview(InterviewDto interviewDto);
}
