import { Component, Input, OnInit } from '@angular/core';
import { SubjectData } from 'src/app/interfaces/subject-data';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { UserService } from 'src/app/services/user.service';
import { Subscription } from 'rxjs';

@Component({
  selector: 'app-owner-page',
  templateUrl: './owner-page.component.html',
  styleUrls: ['./owner-page.component.css']
})
export class OwnerPageComponent implements OnInit {

  personalData!: SubjectData;
  sub!: Subscription;
  constructor(private userService: UserService) {
    this.personalData = {} as SubjectData;
  }

  @Input()
  userDetails: any
  initialDetails: any
  editMode = false
  id!: number
  ngOnInit(): void {
    this.getPersonalData();
  }

  detailsForm = new FormGroup({
    firstName: new FormControl('', Validators.required),
    lastName: new FormControl('', Validators.required),
    email: new FormControl('', Validators.required),
    recoveryEmail: new FormControl('', Validators.required),
    phoneNumber: new FormControl('', Validators.required),
    userName: new FormControl('', Validators.required),
    dateOfBirth: new FormControl('', Validators.required),

  })

  getPersonalData() {
    this.sub = this.userService.getPersonalData().subscribe({
      next: (data: SubjectData) => {
        this.userDetails = data
        this.initialDetails = JSON.parse(JSON.stringify(data));
        this.detailsForm.controls['firstName'].setValue(data.firstName)
        this.detailsForm.controls['lastName'].setValue(data.lastName)
        this.detailsForm.controls['email'].setValue(data.email)
        this.detailsForm.controls['phoneNumber'].setValue(data.phoneNumber)
        this.detailsForm.controls['recoveryEmail'].setValue(data.recoveryEmail)
        this.detailsForm.controls['userName'].setValue(data.username)
        this.detailsForm.controls['dateOfBirth'].setValue(data.dateOfBirth)

      },
    });
  }

}
