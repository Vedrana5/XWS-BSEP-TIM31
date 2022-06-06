package com.example.AgentApp.model;

import lombok.Data;

import javax.persistence.*;
import java.util.List;


@Entity
@Data
@Table(name="companies")
public class Company {

    public enum StatusofCompany {
        APPROVED,  DECLINED,PENDING
    }

    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;

    @Column
    private String name;
    @Column
    private String website;

    @Column
    private String email;

    @Column
    private String phoneNumber;

    @Column
    private String countryOfOrigin;


    @Column
    private String founder;

    @Column
    private String numberOfEmpl;

    @ManyToOne
    @JoinColumn(name = "owner_id")
    private User owner;

   @Column
    private CompanyStatus companyStatus;

/*

    @OneToMany
    @JoinColumn(name = "company_id")
    private List<Offer> offers;
*/

}
