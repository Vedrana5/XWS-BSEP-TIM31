package com.example.AgentApp.dto;


import lombok.AllArgsConstructor;
import lombok.Getter;
import lombok.NoArgsConstructor;
import lombok.Setter;

import javax.validation.constraints.NotBlank;

@Getter
@Setter
@AllArgsConstructor
@NoArgsConstructor
public class InterviewDto {

    @NotBlank
    public String comment;

    @NotBlank
    public String difficulty;

    @NotBlank
    public int rating;

    @NotBlank
    public String userName;

    @NotBlank
    public Long companyId;



}
