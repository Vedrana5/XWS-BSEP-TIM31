package com.example.AgentApp.service.impl;

import com.example.AgentApp.dto.NewOfferDto;
import com.example.AgentApp.model.Company;
import com.example.AgentApp.model.Offer;
import com.example.AgentApp.model.User;
import com.example.AgentApp.repository.CompanyRepository;
import com.example.AgentApp.repository.OfferRepository;
import com.example.AgentApp.repository.UserRepository;
import com.example.AgentApp.service.CompanyService;
import com.example.AgentApp.service.OfferService;
import com.example.AgentApp.service.UserService;
import org.springframework.beans.factory.annotation.Autowired;

import java.util.List;
import java.util.Optional;

public class OfferServiceImpl  implements OfferService {

   @Autowired
   private UserService userService;

   @Autowired
   private CompanyRepository companyRepository;

   @Autowired
   private OfferRepository offerRepository;

    @Override
    public Offer addOffer(NewOfferDto newOfferDto) {
        Offer offer=new Offer();
        offer.setName(newOfferDto.getName());
        offer.setDescription(newOfferDto.getDescription());
        offer.setPosition(newOfferDto.getPosition());
        Company company=companyRepository.getById(newOfferDto.getCompanyId());
        offer.setCompany(company);
        offerRepository.save(offer);
        return offer;
    }

    @Override
    public List<Offer> getAllOffers() {
        List<Offer> offers = offerRepository.findAll();
        return offers;
    }
}
