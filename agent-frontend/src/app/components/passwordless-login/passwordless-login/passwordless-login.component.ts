import { Component, OnInit } from '@angular/core';
import { MatSnackBar } from '@angular/material/snack-bar';
import { ActivatedRoute, Router } from '@angular/router';
import { UserService } from 'src/app/services/user.service';

@Component({
  selector: 'app-passwordless-login',
  templateUrl: './passwordless-login.component.html',
  styleUrls: ['./passwordless-login.component.css']
})
export class PasswordlessLoginComponent implements OnInit {

  constructor(private _router: Router, private _route: ActivatedRoute, private _service: UserService,
    private _snackBar: MatSnackBar) { }

  ngOnInit(): void {
    let token = decodeURI(this._route.snapshot.paramMap.get('token') || "")
    this._service.checkPasswordlessToken(token).subscribe(
      res => {
        this._snackBar.open('Welcome!', '', {
          duration: 3000,
        });
        if (localStorage.getItem('role') == 'ADMIN') {
          this._router.navigate(['/adminHome']);
        } else if (localStorage.getItem('role') == 'REGISTERED_USER') {
          this._router.navigate(['/userHome']);
        }
        else {
          this._router.navigate(['/ownerHome']);

        }
      },
      err => {
        this._snackBar.open('Your link has expired. We have sent you a new one, please check your mail.', '', {
          duration: 3000,
        });
        this._router.navigate(['/']);
      }
    )
  }

}
