package com.example.AgentApp.repository;

import com.example.AgentApp.model.Offer;
import org.springframework.data.jpa.repository.JpaRepository;

import java.util.List;

public interface OfferRepository extends JpaRepository<Offer,Long> {
    List<Offer> getAllByCompanyId(Long companyId);
}
