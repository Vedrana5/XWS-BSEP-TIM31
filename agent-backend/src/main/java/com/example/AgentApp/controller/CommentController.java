package com.example.AgentApp.controller;


import com.example.AgentApp.dto.CommentDto;
import com.example.AgentApp.dto.CreateCompanyDto;
import com.example.AgentApp.model.Comment;
import com.example.AgentApp.model.Company;
import com.example.AgentApp.model.User;
import com.example.AgentApp.service.CommentService;
import com.example.AgentApp.service.CompanyService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

import java.util.List;

@RestController
@RequestMapping(value = "/comments")
public class CommentController {

    @Autowired
    private CommentService commentService;

    @Autowired
    private CompanyService companyService;



    @CrossOrigin(origins = "http://localhost:4200")
    @PostMapping(value = "/createNewComment")
    public ResponseEntity<String> createCompany(@RequestBody CommentDto commentDto) {
        Comment comment = commentService.createComment(commentDto);
        if (comment != null) {
            return new ResponseEntity<>("SUCCESS!", HttpStatus.CREATED);
        }
        return new ResponseEntity<>("ERROR!", HttpStatus.INTERNAL_SERVER_ERROR);
    }


    @CrossOrigin(origins = "http://localhost:4200")
    @GetMapping("/getAllByCompany/{id}")
    public List<Comment> getCompanyComment(@PathVariable Long id){
        Company company = companyService.findById(id);
        List<Comment> comments =commentService.getAllByCompany(company.getId());

        return comments;
    }

}
