import { HttpErrorResponse } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { MatDialogRef } from '@angular/material/dialog';
import { MatSnackBar } from '@angular/material/snack-bar';
import { Router } from '@angular/router';
import { ChangePasswordDto } from 'src/app/interfaces/change-password-dto';
import { LogedUser } from 'src/app/interfaces/loged-user';
import { CompanyService } from 'src/app/services/company.service';
import { UserService } from 'src/app/services/user.service';

@Component({
  selector: 'app-change-password',
  templateUrl: './change-password.component.html',
  styleUrls: ['./change-password.component.css']
})
export class ChangePasswordComponent implements OnInit {

  passMatch: boolean = false;
  changedPassword!: ChangePasswordDto;
  constructor(private userService: UserService, private router: Router, private _snackBar: MatSnackBar, private dialogRef: MatDialogRef<ChangePasswordComponent>) {
    this.changedPassword = {} as ChangePasswordDto;
  }
  ngOnInit(): void {

  }
  form = new FormGroup({
    currentPassword: new FormControl(null, [
      Validators.required,
      Validators.minLength(8),
      Validators.maxLength(30),
      Validators.pattern(
        '^(?=.*[a-z])(?=.*[A-Z])(?=.*[0-9])(?=.*[!"#$@%&()*<>+_|~]).*$'
      )]),
    newPassword: new FormControl(null, [
      Validators.required,
      Validators.minLength(10),
      Validators.maxLength(30),
      Validators.pattern(
        '^(?=.*[a-z])(?=.*[A-Z])(?=.*[0-9])(?=.*[!"#$@%&()*<>+_|~]).*$'
      )]),
    repeatedPassword: new FormControl(null, [
      Validators.required,
      Validators.minLength(10),
      Validators.maxLength(30),
      Validators.pattern(
        '^(?=.*[a-z])(?=.*[A-Z])(?=.*[0-9])(?=.*[!"#$@%&()*<>+_|~]).*$'
      )]),
  })

  onPasswordInput(): void {
    this.passMatch =
      this.form.value.newPassword === this.form.value.repeatedPassword;
  }

  createNewPassword(): void {
    this.changedPassword.newPassword = this.form.value.newPassword;
    this.changedPassword.currentPassword = this.form.value.currentPassword;
  }

  submit() {
    if (this.form.invalid) {
      return;
    }

    this.createNewPassword();

    this.userService.changePassword(this.changedPassword).subscribe({
      next: () => {
        this.router.navigate(['/']);
        this._snackBar.open(
          'Password is changed!',
          'Dismiss', {
          duration: 3000
        });

        this.router.navigate(['/']);
        this.dialogRef.close();

      },
      error: (err: HttpErrorResponse) => {
        let parts = err.error.split(':');
        let mess = parts[parts.length - 1];
        let description = mess.substring(1, mess.length - 4);
        this._snackBar.open(description, 'Dismiss', {
          duration: 3000
        });
      },
      complete: () => console.info('complete')
    });

    this.userService.changePassword(this.form.value).subscribe();


  }

}
