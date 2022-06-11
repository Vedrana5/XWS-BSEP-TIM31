package com.example.AgentApp.model;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

import javax.persistence.*;



@Entity
@AllArgsConstructor
@NoArgsConstructor
@Data
@Table(name = "interview")
public class Interview {

    public enum Difficulty {
        EASY, INTERMEDIATE, HARD
    }

    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;


    @Column(nullable = false)
    private String comment;

    @ManyToOne
    @JoinColumn(name = "user_id")
    private User user;

    @ManyToOne
    @JoinColumn(name = "company_id")
    private Company company;

    @Column(nullable = false)
    private Difficulty difficulty;


    @Column(nullable = false)
    private int rating;


}
