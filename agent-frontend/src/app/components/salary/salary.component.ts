import { HttpErrorResponse } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup, Validators } from '@angular/forms';
import { MatDialogRef } from '@angular/material/dialog';
import { MatSnackBar } from '@angular/material/snack-bar';
import { Router } from '@angular/router';
import { Salary } from 'src/app/interfaces/salaty-dto';
import { CompanyService } from 'src/app/services/company.service';

@Component({
  selector: 'app-salary',
  templateUrl: './salary.component.html',
  styleUrls: ['./salary.component.css']
})
export class SalaryComponent implements OnInit {

  salary!: Salary;
  createForm!: FormGroup;
  cid!: string;
  constructor(private _formBuilder: FormBuilder,
    private companyService: CompanyService,
    private router: Router,
    public dialogRef: MatDialogRef<SalaryComponent>,
    private _snackBar: MatSnackBar) {

    this.salary = {} as Salary
    this.createForm = this._formBuilder.group({
      Salary: new FormControl('', [
        Validators.required,
        Validators.pattern('^(?=.*[a-zA-Z])(?=.*[0-9]).*$'
        )
      ]),
      Position: new FormControl('', [Validators.required]),
    })
  }

  ngOnInit(): void {
    this.cid = this.router.url.substring(13);
  }

  add() {
    this.createSallary();
    console.log(this.createSallary);

    this.companyService.createSallary(this.salary).subscribe({
      next: (res) => {

        this.dialogRef.close({ event: "Created Sallary comment", data: res });
        this._snackBar.open(
          'You have created a job offer.',
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


  createSallary() {



    this.salary.position = this.createForm.value.Position;
    this.salary.salary = this.createForm.value.Salary;
    this.salary.email = localStorage.getItem("email")!;
    this.salary.companyId = parseInt(this.cid);
    console.log(this.cid)

  }
}
