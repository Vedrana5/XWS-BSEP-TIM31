package com.example.AgentApp.repository;

import com.example.AgentApp.model.Comment;
import com.example.AgentApp.model.CustomToken;
import org.springframework.data.jpa.repository.JpaRepository;

public interface CommentRepository extends JpaRepository<Comment, Long> {
}
