package com.example.AgentApp.model;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

import javax.persistence.*;


@Entity
@AllArgsConstructor
@NoArgsConstructor
@Data
@Table(name = "salary")
public class Salary {

    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;

    @Column(nullable = false)
    private String salary;

    @Column(nullable = false)
    private String position;


    @ManyToOne
    @JoinColumn(name = "company_id")
    private Company company;


    @ManyToOne
    @JoinColumn(name = "user_id")
    private User user;
}
