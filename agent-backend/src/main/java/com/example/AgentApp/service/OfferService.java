package com.example.AgentApp.service;

import com.example.AgentApp.dto.NewOfferDto;
import com.example.AgentApp.model.Offer;

import java.util.List;

public interface OfferService {
    Offer addOffer(NewOfferDto newOfferDto);

    List<Offer> getAllOffers();

    List<Offer> getAllByCompany(Long companyId);
}
