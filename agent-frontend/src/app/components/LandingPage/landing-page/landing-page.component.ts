import { Component, OnInit } from '@angular/core';
import { Form, FormBuilder, FormControl, FormGroup, NgForm, Validators } from '@angular/forms';
import { Subscription } from 'rxjs';
import { Router } from '@angular/router';
import { MatSnackBar } from '@angular/material/snack-bar';
import { UserService } from 'src/app/services/user.service';

@Component({
  selector: 'app-landing-page',
  templateUrl: './landing-page.component.html',
  styleUrls: ['./landing-page.component.css'],
})
export class LandingPageComponent implements OnInit {
  sub!: Subscription;
  tfaEnabled = false;
  ngForm!: FormGroup
  constructor(
    private authService: UserService,
    private _router: Router,
    private _snackBar: MatSnackBar,
    private formBuilder: FormBuilder,

  ) { }
  emaill!: string;

  ngOnInit(): void {
    this.ngForm = this.formBuilder.group({
      emaill: new FormControl('', [
        Validators.required,

      ]),
      password: new FormControl('', [
        Validators.required,
        Validators.minLength(10),
        Validators.maxLength(30),
        Validators.pattern(
          '^(?=.*[a-z])(?=.*[A-Z])(?=.*[0-9])(?=.*[!"#$@%&()*<>+_|~]).*$'
        )])
    });
  }

  forgotPass() {
    console.log(this.emaill);
    this.authService.sendCode(this.emaill).subscribe(
      res => {
        localStorage.setItem('emailForReset', this.emaill);
        this._router.navigate(['/resetPassword']);

      },
      err => {
        this._snackBar.open("User with this email does not exist!", "", {
          duration: 3000
        });
      }
    );
  }
  onSubmit(f: NgForm) {

    const loginObserver = {
      next: (x: any) => {
        this._snackBar.open('     Welcome', 'Dismiss');
        if (localStorage.getItem('role') == 'ADMIN') {
          this._router.navigate(['/adminHome']);
        } else if (localStorage.getItem('role') == 'REGISTERED_USER') {
          this._router.navigate(['/userHome']);
        }
        else {
          this._router.navigate(['/ownerHome']);

        }


      },
      error: (err: any) => {
        this._snackBar.open(
          'Email or password are incorrect.Try again,please.',
          'Dismiss'
        );
      },
    };

    this.authService.login(f.value).subscribe(loginObserver);
  }


  check2FAStatus() {
    this.authService.check2FAStatus(this.emaill).subscribe(

      res => {
        this.tfaEnabled = res
        console.log(res)
      }
    )
  }

  passwordless() {
    console.log(this.emaill);
    this.authService.sendLink(this.emaill).subscribe(
      res => {
        this._snackBar.open('Check your email. Click on the magic link to log in.', '', {
          duration: 3000,
        });
      },
      err => {
        this._snackBar.open("User with this email does not exist!", "", {
          duration: 3000
        });
      }
    );

  }
}