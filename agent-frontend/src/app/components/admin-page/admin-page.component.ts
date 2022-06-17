
import { ChangeDetectorRef, Component, Input, OnInit } from '@angular/core';
import { SubjectData } from 'src/app/interfaces/subject-data';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { UserService } from 'src/app/services/user.service';
import { Subscription } from 'rxjs';
import { MatDialog, MatDialogConfig } from '@angular/material/dialog';
import { ChangePasswordComponent } from '../change-password/change-password.component';
import { TfaComponent } from 'src/app/tfa/tfa/tfa.component';

@Component({
  selector: 'app-admin-page',
  templateUrl: './admin-page.component.html',
  styleUrls: ['./admin-page.component.css']
})
export class AdminPageComponent implements OnInit {

  personalData!: SubjectData;
  sub!: Subscription;
  constructor(private userService: UserService, public matDialog: MatDialog) {
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
        console.log(data.firstName)

      },
    });
  }

  openModalForChangePassword(): void {
    const dialogConfig = new MatDialogConfig();
    dialogConfig.disableClose = false;
    dialogConfig.id = 'modal-component';
    dialogConfig.height = 'fit-content';
    dialogConfig.width = '500px';
    this.matDialog.open(ChangePasswordComponent, dialogConfig);
  }

  opet2fa() {
    const dialogConfig = new MatDialogConfig();
    dialogConfig.disableClose = false;
    dialogConfig.id = 'modal-component';
    dialogConfig.height = 'fit-content';
    dialogConfig.width = '800px';
    this.matDialog.open(TfaComponent, dialogConfig);
  }
}
