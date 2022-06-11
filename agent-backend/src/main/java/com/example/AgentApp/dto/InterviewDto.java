package com.example.AgentApp.dto;


import lombok.AllArgsConstructor;
import lombok.Getter;
import lombok.NoArgsConstructor;
import lombok.Setter;

@Getter
@Setter
@AllArgsConstructor
@NoArgsConstructor
public class InterviewDto {


    public String comment;
    public String difficulty;
    public int rating;
    public String userName;
    public Long companyId;



}
