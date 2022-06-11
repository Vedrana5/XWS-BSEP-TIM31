package com.example.AgentApp.service.impl;

import com.example.AgentApp.dto.CommentDto;
import com.example.AgentApp.model.Comment;
import com.example.AgentApp.model.Company;
import com.example.AgentApp.model.CompanyStatus;
import com.example.AgentApp.model.User;
import com.example.AgentApp.repository.CommentRepository;
import com.example.AgentApp.repository.CompanyRepository;
import com.example.AgentApp.repository.UserRepository;
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


    @Override
    public Comment createComment(CommentDto commentDto) {

        Comment comment = new Comment();
        User u = userRepository.findByUsername(commentDto.getUserName());
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

}