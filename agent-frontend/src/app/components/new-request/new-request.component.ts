import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup, Validators } from '@angular/forms';
import { MatSnackBar } from '@angular/material/snack-bar';
import { Router } from '@angular/router';
import { HttpErrorResponse } from '@angular/common/http';
import { CompanyDto } from 'src/app/interfaces/company-dto';
import { CompanyService } from 'src/app/services/company.service';

@Component({
  selector: 'app-new-request',
  templateUrl: './new-request.component.html',
  styleUrls: ['./new-request.component.css']
})
export class NewRequestComponent implements OnInit {

  errorMessage!: string;
  createForm!: FormGroup;
  formData!: FormData;
  newCompany!: CompanyDto;
  email: any;
  constructor(private formBuilder: FormBuilder,
    private _snackBar: MatSnackBar,
    private router: Router, private companyService: CompanyService) {
    this.newCompany = {} as CompanyDto;
    this.email = localStorage.getItem('email')


  }


  ngOnInit(): void {
    this.createForm = this.formBuilder.group({
      name: new FormControl('', [
        Validators.required,
        Validators.pattern('^[A-ZŠĐŽČĆ][a-zšđćčžA-ZŠĐŽČĆ0-9 ]*$'),
      ]),
      website: new FormControl('', [
        Validators.required,
        Validators.pattern('^[A-ZŠĐŽČĆ][a-zšđćčžA-ZŠĐŽČĆ0-9 ]*$'),
      ]),
      email: new FormControl(null, [
        Validators.required,
        Validators.email,

      ]),
      phoneNumber: new FormControl(null, [
        Validators.required,
        Validators.pattern('^[0]{1}[0-9]{8}$'),
      ]),

      countryOfOrigin: new FormControl(null, [
        Validators.required,
        Validators.pattern('^[A-ZŠĐŽČĆ][a-zšđćčžA-ZŠĐŽČĆ0-9 ]*$'),
      ]),
      founder: new FormControl(null, [
        Validators.required,
        Validators.pattern('^[A-ZŠĐŽČĆ][a-zšđćčžA-ZŠĐŽČĆ0-9 ]*$'),
      ]),
      numberOfEmpl: new FormControl(null, [
        Validators.required,
        Validators.pattern(/^-?(0|[1-9]\d*)?$/)
      ]),


    })

  }

  onSubmit(): void {

    if (this.createForm.invalid) return;
    this.createCompany();

    this.companyService.createCompany(this.newCompany).subscribe({
      next: () => {
        this._snackBar.open(
          'You have successfully created a request that has been sent to administration for approval.',
          '', {
          duration: 3000
        });
        this.router.navigate(['/userHome']);


      },
      error: (err: HttpErrorResponse) => {
        this._snackBar.open(err.error.message + "!", 'Dismiss', {
          duration: 3000
        });
      },
      complete: () => console.info('complete')
    });

  }






  createCompany(): void {
    this.newCompany.name = this.createForm.value.name;
    this.newCompany.website = this.createForm.value.website;
    this.newCompany.email = this.createForm.value.email;
    this.newCompany.phoneNumber = this.createForm.value.phoneNumber;
    this.newCompany.countryOfOrigin = this.createForm.value.countryOfOrigin;
    this.newCompany.founder = this.createForm.value.founder;
    this.newCompany.numberOfEmpl = this.createForm.value.numberOfEmpl;
    this.newCompany.ownerUsername = this.email;

  }
}
