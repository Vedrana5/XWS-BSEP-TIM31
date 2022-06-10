
import { DatePipe } from '@angular/common';
import { HttpErrorResponse } from '@angular/common/http';
import { Component, EventEmitter, OnInit, Output } from '@angular/core';
import { FormBuilder, FormControl, FormGroup, Validators } from '@angular/forms';
import { MatDialogRef } from '@angular/material/dialog';
import { MatSnackBar } from '@angular/material/snack-bar';
import { Router } from '@angular/router';
import { JobOffer } from 'src/app/interfaces/jobOffer';
import { CompanyService } from 'src/app/services/company.service';

@Component({
  selector: 'app-job-offer',
  templateUrl: './job-offer.component.html',
  styleUrls: ['./job-offer.component.css']
})
export class JobOfferComponent implements OnInit {


  jobOfferRequest: JobOffer;
  createForm!: FormGroup;
  requirements!: string[];
  allOffers!: JobOffer;
  cid!: string;
  constructor(public dialogRef: MatDialogRef<JobOfferComponent>,
    private _snackBar: MatSnackBar,
    private router: Router,
    private _formBuilder: FormBuilder,
    private companyService: CompanyService,
    private datePipe: DatePipe) {

    this.jobOfferRequest = {} as JobOffer
    this.allOffers = {} as JobOffer
    this.requirements = [];
    this.createForm = this._formBuilder.group({
      Requirements: new FormControl('', [
        Validators.required,
      ]),
      Position: new FormControl('', [
        Validators.required
      ]),
      Description: new FormControl('', [
        Validators.required
      ]),
      DueDate: new FormControl('',
        [
          Validators.required
        ])
    })

  }

  ngOnInit(): void {
    this.cid = this.router.url.substring(9);
  }

  addItem() {



  }
  add() {
    this.createJobOffer();
    console.log(this.createJobOffer);

    this.companyService.createOffer(this.jobOfferRequest).subscribe({
      next: (res) => {

        this.dialogRef.close({ event: "Created Job offer", data: res });
        this._snackBar.open(
          'You have created a job offer.',
          'Dismiss', {
          duration: 3000
        });
      },
      error: (err: HttpErrorResponse) => {

        this._snackBar.open(err.error.message + "!", 'Dismiss', {
          duration: 3000
        });
      },
      complete: () => console.info('complete')
    });
  }

  createJobOffer() {


    //this.jobOfferRequest.requirements = this.requirements;
    this.jobOfferRequest.position = this.createForm.value.Position;
    this.jobOfferRequest.description = this.createForm.value.Description;
    this.jobOfferRequest.companyId = parseInt(this.cid);
    const dateParsed = this.datePipe.transform(this.createForm.value.DueDate, 'MM/dd/yyyy');
    // this.jobOfferRequest.dueDate = dateParsed;

  }


}
