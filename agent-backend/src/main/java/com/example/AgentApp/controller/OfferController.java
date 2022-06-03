package com.example.AgentApp.controller;

import com.example.AgentApp.dto.NewOfferDto;
import com.example.AgentApp.model.Offer;
import com.example.AgentApp.service.OfferService;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
@RequestMapping(value = "/offers")
public class OfferController {

    private OfferService offerService;

    @PostMapping("/newOffer")
    public Offer createNewOffer(@RequestBody NewOfferDto newOfferDto){
        Offer offer = offerService.addOffer(newOfferDto);
        return offer;
    }
}
