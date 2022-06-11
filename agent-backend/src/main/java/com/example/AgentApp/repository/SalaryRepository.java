package com.example.AgentApp.repository;

import com.example.AgentApp.model.Comment;
import com.example.AgentApp.model.Salary;
import org.springframework.data.jpa.repository.JpaRepository;

public interface SalaryRepository  extends JpaRepository<Salary, Long> {
}
