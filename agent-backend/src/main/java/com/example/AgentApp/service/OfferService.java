package com.example.AgentApp.service;

import com.example.AgentApp.dto.NewOfferDto;
import com.example.AgentApp.model.Offer;

public interface OfferService {
    Offer addOffer(NewOfferDto newOfferDto);
}
