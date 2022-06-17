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
public class NewOfferDto {
    public Long companyId;

    @NotBlank
    public String position;

    @NotBlank
    public String description;

    @NotBlank
    public String dateCreated;

    @NotBlank
    public String dueDate;
}
