import { HttpErrorResponse } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup, Validators } from '@angular/forms';
import { MatDialogRef } from '@angular/material/dialog';
import { MatSnackBar } from '@angular/material/snack-bar';
import { Router } from '@angular/router';
import { Interview } from 'src/app/interfaces/interview-dto';
import { CompanyService } from 'src/app/services/company.service';

@Component({
  selector: 'app-interview',
  templateUrl: './interview.component.html',
  styleUrls: ['./interview.component.css']
})
export class InterviewComponent implements OnInit {

  interview!: Interview;
  createForm!: FormGroup;
  cid!: string;
  constructor(private _formBuilder: FormBuilder,
    private companyService: CompanyService,
    private router: Router,
    public dialogRef: MatDialogRef<InterviewComponent>,
    private _snackBar: MatSnackBar) {

    this.interview = {} as Interview
    this.createForm = this._formBuilder.group({
      Comment: new FormControl('', [Validators.required]),
      Difficulty: new FormControl('', [Validators.required]),
      Rating: new FormControl('', [Validators.required, Validators.min(1), Validators.max(5)])
    })
  }

  ngOnInit(): void {
    console.log(this.router.url);
    this.cid = this.router.url.substring(13);

  }

  add() {

    if (this.createForm.invalid) return;
    this.createInterview();
    console.log(this.createInterview);

    this.companyService.createInterview(this.interview).subscribe({
      next: (res) => {

        this.dialogRef.close({ event: "Created Interview", data: res });
        this._snackBar.open(
          'You have created a interview comment.',
          'Dismiss', {
          duration: 3000
        });
        window.location.reload();
      },
      error: (err: HttpErrorResponse) => {

        this._snackBar.open(err.error.message + "!", 'Dismiss', {
          duration: 3000
        });
      },
      complete: () => console.info('complete')
    });
  }

  createInterview() {
    console.log(this.createForm.value.Name);

    this.interview.comment = this.createForm.value.Comment;
    this.interview.email = localStorage.getItem("email")!;
    this.interview.companyId = parseInt(this.cid);
    this.interview.difficulty = this.createForm.value.Difficulty;
    this.interview.rating = this.createForm.value.Rating;
    console.log(this.cid)
  }


}
