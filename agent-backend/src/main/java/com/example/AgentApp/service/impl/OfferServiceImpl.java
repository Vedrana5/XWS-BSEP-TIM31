package com.example.AgentApp.service.impl;

import com.example.AgentApp.dto.NewOfferDto;
import com.example.AgentApp.model.Company;
import com.example.AgentApp.model.Offer;
import com.example.AgentApp.repository.CompanyRepository;
import com.example.AgentApp.repository.OfferRepository;
import com.example.AgentApp.service.OfferService;
import com.example.AgentApp.service.UserService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.time.LocalDate;
import java.util.List;
import java.util.Optional;

@Service
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

        offer.setDescription(newOfferDto.getDescription());
        offer.setPosition(newOfferDto.getPosition());
        offer.setDateCreated(LocalDate.now());
       // offer.setDueDate(LocalDate.parse(newOfferDto.dueDate));
        Optional<Company> company=companyRepository.findById(newOfferDto.getCompanyId());
        offer.setCompany(company.get());
        System.out.print("sdfesdfesdfeesdf"+ offer);
        offerRepository.save(offer);
        return offer;
    }

    @Override
    public List<Offer> getAllOffers() {
        List<Offer> offers = offerRepository.findAll();
        return offers;
    }

    @Override
    public List<Offer> getAllByCompany(Long companyId) {
     List<Offer> offers=offerRepository.getAllByCompanyId(companyId);
     return offers;
    }
}
