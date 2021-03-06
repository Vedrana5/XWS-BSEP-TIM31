package com.example.AgentApp.model;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;
import org.apache.commons.codec.binary.Base32;
import org.springframework.security.core.GrantedAuthority;

import javax.persistence.*;
import java.security.SecureRandom;
import java.util.Collection;
import java.util.Date;
import java.util.HashSet;
import java.util.Set;

@Entity
@AllArgsConstructor
@NoArgsConstructor
@Data
@Table(name = "users")
public class User {

    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;

    @Column(unique = true, nullable = false)
    private String username;

    @Column(nullable = false)
    private String firstName;

    @Column(nullable = false)
    private String lastName;

    @Column(nullable = false)
    private boolean isConfirmed;

    @Column(nullable = false)
    private String phoneNumber;

    @Column( nullable = false)
    private String password;

    @Column(nullable = false)
    private Date dateOfBirth;

    @Column(unique = true, nullable = false)
    private String email;

    @Column(nullable = false)
    private String recoveryEmail;

    @Column(nullable = false)
    private UserRole role;


    private boolean isUsingFa;

    private String secret;










}
