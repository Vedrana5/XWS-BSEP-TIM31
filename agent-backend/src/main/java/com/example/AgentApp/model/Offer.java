package com.example.AgentApp.model;

import lombok.Getter;
import lombok.Setter;

import javax.persistence.*;


@Getter
@Setter
@Entity
@Table
public class Offer {

    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;


    @Column
    private String name;

    @Column
    private String description;

    @Column
    private String position;

    @ManyToOne
    @JoinColumn(name = "company_id")
    private Company company;



}
