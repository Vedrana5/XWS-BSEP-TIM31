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
public class SalaryDto {

    @NotBlank
    public String salary;
    @NotBlank
    public String position;
    @NotBlank
    public Long companyId;
    @NotBlank
    public String userName;
}
