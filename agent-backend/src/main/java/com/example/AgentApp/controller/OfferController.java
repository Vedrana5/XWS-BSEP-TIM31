package com.example.AgentApp.controller;

import com.example.AgentApp.dto.NewOfferDto;
import com.example.AgentApp.model.Offer;
import com.example.AgentApp.service.OfferService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.*;

import java.util.List;

@RestController
@RequestMapping(value = "/offers")
public class OfferController {

    @Autowired
    private OfferService offerService;

    @CrossOrigin(origins = "http://localhost:4200")
    @PostMapping(value = "/newOffer")
    public Offer createNewOffer(@RequestBody NewOfferDto newOfferDto){

        Offer offer = offerService.addOffer(newOfferDto);
        System.out.print("fdfsdfsdfdsfdsfdsf"+offer);
        return offer;
    }

    @GetMapping(value = "/getAllOffers")
    public List<Offer> allJobOffers(){
        List<Offer> offers = offerService.getAllOffers();
        return offers;

    }

    @GetMapping(value = "/getAllByCompany/{companyId}")
    public List<Offer> allOffersForCompany(@PathVariable Long companyId){
        List<Offer> offers = offerService.getAllByCompany(companyId);
       return offers;
    }
}
