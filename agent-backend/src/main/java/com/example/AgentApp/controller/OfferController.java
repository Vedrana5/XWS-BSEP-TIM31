package com.example.AgentApp.controller;

import com.example.AgentApp.dto.NewOfferDto;
import com.example.AgentApp.model.Offer;
import com.example.AgentApp.service.LoggerService;
import com.example.AgentApp.service.OfferService;
import com.example.AgentApp.service.impl.LoggerServiceImpl;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.security.access.prepost.PreAuthorize;
import org.springframework.security.core.context.SecurityContextHolder;
import org.springframework.web.bind.annotation.*;

import java.util.List;

@CrossOrigin(origins = "https://localhost:4200")
@RestController
@RequestMapping(value = "/offers")
public class OfferController {

    @Autowired
    private OfferService offerService;

    private  final LoggerService loggerService;


    public OfferController() {

        this.loggerService = new LoggerServiceImpl(this.getClass());
    }


    //  @PreAuthorize("hasAuthority('CREATE_OFFER_PERMISSION')")
    @CrossOrigin(origins = "https://localhost:4200")
    @PostMapping(value = "/newOffer")
    public Offer createNewOffer(@RequestBody NewOfferDto newOfferDto){
        Offer offer = offerService.addOffer(newOfferDto);
        if(offer==null) {
            loggerService.createOfferFailed(SecurityContextHolder.getContext().getAuthentication().getName(),newOfferDto.getCompanyId());


        }
        loggerService.createOfferSuccess(SecurityContextHolder.getContext().getAuthentication().getName(),newOfferDto.getCompanyId());

        return offer;
    }

    @CrossOrigin(origins = "https://localhost:4200")
    @GetMapping(value = "/getAllOffers")
    public List<Offer> allJobOffers(){
        List<Offer> offers = offerService.getAllOffers();
        return offers;

    }

    @CrossOrigin(origins = "https://localhost:4200")
    @GetMapping(value = "/getAllByCompany/{companyId}")
    public List<Offer> allOffersForCompany(@PathVariable Long companyId){
        List<Offer> offers = offerService.getAllByCompany(companyId);
       return offers;
    }
}
