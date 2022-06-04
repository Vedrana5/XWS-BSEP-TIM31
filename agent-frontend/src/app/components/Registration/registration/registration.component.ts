import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup, Validators } from '@angular/forms';
import { MatSnackBar } from '@angular/material/snack-bar';
import { Router } from '@angular/router';
import { SubjectData } from 'src/app/interfaces/subject-data';
import { UserService } from 'src/app/services/user.service';

@Component({
  selector: 'app-registration',
  templateUrl: './registration.component.html',
  styleUrls: ['./registration.component.css']
})
export class RegistrationComponent implements OnInit {

  errorMessage!: string;
  createForm!: FormGroup;
  formData!: FormData;
  uploaded: boolean = false;
  passMatch: boolean = false;
  newSubject!: SubjectData;

  constructor(private formBuilder: FormBuilder,
    private _snackBar: MatSnackBar,
    private router: Router,
    private authService: UserService) {
    this.newSubject = {} as SubjectData;
  }

  ngOnInit(): void {
    this.createForm = this.formBuilder.group({
      firstName: new FormControl('', [
        Validators.required,
        Validators.pattern('^[A-ZŠĐŽČĆ][a-zšđćčžA-ZŠĐŽČĆ ]*$'),
      ]),
      lastName: new FormControl(null, [
        Validators.required,
        Validators.pattern('^[A-ZŠĐŽČĆ][a-zšđćčžA-ZŠĐŽČĆ ]*$'),
      ]),

      email: new FormControl(null, [Validators.required, Validators.email]),
      recoveryMail: new FormControl(null, [Validators.required, Validators.email]),
      password: new FormControl(null, [
        Validators.required,
        Validators.minLength(10),
        Validators.maxLength(30),
        Validators.pattern(
          '^(?=.*[a-z])(?=.*[A-Z])(?=.*[0-9])(?=.*[!"#$@%&()*<>+_|~]).*$'
        )]),
      passConfirmed: new FormControl(null, [
        Validators.required,
        Validators.minLength(10),
        Validators.maxLength(30),
        Validators.pattern(
          '^(?=.*[a-z])(?=.*[A-Z])(?=.*[0-9])(?=.*[!"#$@%&()*<>+_|~]).*$'
        )]),
    });

  }

  onPasswordInput(): void {
    this.passMatch =
      this.createForm.value.password === this.createForm.value.passConfirmed;
  }

  onSubmit(): void {
    this.createUser();
    this.authService.createSubject(this.newSubject).subscribe(
      (res) => {
        console.log(res);
        if (res.isPawned) alert("The choosen password is weak and already pawned. Please, go change it.");
        this.router.navigate(['/']);
        this._snackBar.open(
          'Your registration request has been sumbitted. Please check your email and confirm your email address to activate your account.',
          'Dismiss'
        );

      },
      (err) => {
        let parts = err.error.split(':');
        let mess = parts[parts.length - 1];
        let description = mess.substring(1, mess.length - 4);
        this._snackBar.open(description, 'Dismiss');
      }
    );
  }

  createUser(): void {

    this.newSubject.firstName = this.createForm.value.firstName;
    this.newSubject.lastName = this.createForm.value.lastName;
    this.newSubject.email = this.createForm.value.email;
    this.newSubject.recoveryMail = this.createForm.value.recoveryMail;
    this.newSubject.password = this.createForm.value.password;
  }
}
