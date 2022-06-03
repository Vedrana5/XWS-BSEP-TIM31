package com.example.AgentApp.controller;

import com.example.AgentApp.dto.NewOfferDto;
import com.example.AgentApp.model.Offer;
import com.example.AgentApp.service.OfferService;
import org.springframework.web.bind.annotation.*;

import java.util.List;

@RestController
@RequestMapping(value = "/offers")
public class OfferController {

    private OfferService offerService;

    @PostMapping(value = "/newOffer")
    public Offer createNewOffer(@RequestBody NewOfferDto newOfferDto){
        Offer offer = offerService.addOffer(newOfferDto);
        return offer;
    }

    @GetMapping(value = "/getAllOffers")
    public List<Offer> allJobOffers(){
        List<Offer> offers = offerService.getAllOffers();
        return offers;

    }
}
