package com.example.AgentApp.service;

import com.example.AgentApp.dto.CommentDto;
import com.example.AgentApp.model.Comment;

import java.util.List;

public interface CommentService {
    Comment createComment(CommentDto commentDto);

    List<Comment> getAllByCompany(Long id);
}
