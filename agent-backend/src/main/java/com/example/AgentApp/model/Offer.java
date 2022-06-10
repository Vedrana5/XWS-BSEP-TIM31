package com.example.AgentApp.model;

import lombok.Getter;
import lombok.Setter;

import javax.persistence.*;
import java.time.LocalDate;


@Getter
@Setter
@Entity
@Table
public class Offer {

    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;



    @Column
    private String description;

    @Column
    private String position;

    @Column
    private LocalDate dateCreated;

    @Column
    private LocalDate dueDate;

    @ManyToOne
    @JoinColumn(name = "company_id")
    private Company company;





}
