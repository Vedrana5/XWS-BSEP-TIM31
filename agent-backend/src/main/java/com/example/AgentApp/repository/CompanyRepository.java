package com.example.AgentApp.repository;

import com.example.AgentApp.model.Company;
import com.example.AgentApp.model.CompanyStatus;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

import java.util.List;
import java.util.Optional;

@Repository
public interface CompanyRepository extends JpaRepository<Company, Long> {

    List<Company> findAllByCompanyStatus(CompanyStatus status);
}
