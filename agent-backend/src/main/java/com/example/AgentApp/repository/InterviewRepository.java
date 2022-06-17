package com.example.AgentApp.repository;

import com.example.AgentApp.model.CustomToken;
import com.example.AgentApp.model.Interview;
import org.springframework.data.jpa.repository.JpaRepository;

public interface InterviewRepository  extends JpaRepository<Interview, Long> {
}
