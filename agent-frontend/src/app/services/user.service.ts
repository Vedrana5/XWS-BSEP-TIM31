import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { BehaviorSubject, Observable, Subject } from 'rxjs';
import { LogedUser } from 'src/app/interfaces/loged-user';
import { SubjectData } from 'src/app/interfaces/subject-data';
import { map } from 'rxjs/operators';
import { Router } from '@angular/router';

@Injectable({
  providedIn: 'root',
})
export class UserService {
  private currentUserSubject: BehaviorSubject<LogedUser>;
  public currentUser: Observable<LogedUser>;
  private user!: LogedUser;



  private loginStatus = new BehaviorSubject<boolean>(false);

  constructor(private _http: HttpClient, private router: Router) {
    this.currentUserSubject = new BehaviorSubject<LogedUser>(
      JSON.parse(localStorage.getItem('currentUser')!)
    );
    this.currentUser = this.currentUserSubject.asObservable();
  }
  createSubject(newUser: SubjectData) {
    return this._http.post(`http://localhost:8082/auth/register`, newUser, {
      responseType: 'text',
    });
  }
  changePassword(data: any) {
    console.log(data)
    return this._http.put(`http://localhost:8082/auth/changePassword`, data);
  }
  login(model: any): Observable<LogedUser> {
    return this._http.post(`http://localhost:8082/auth/login`, model).pipe(
      map((response: any) => {
        if (response && response.token) {
          localStorage.setItem('token', response.token.accessToken);
          localStorage.setItem('currentUser', JSON.stringify(response));
          localStorage.setItem('role', response.role);
          localStorage.setItem('email', response.email);
          this.currentUserSubject.next(response);
        }
        return this.user;
      })
    );
  }
  public getValue(): LogedUser {
    console.log("Token" + this.currentUserSubject.value.token.accessToken);
    return this.currentUserSubject.value;
  }

  getPersonalData(): Observable<SubjectData> {
    return this._http.get<SubjectData>(
      `http://localhost:8082/users/getBasicInf`
    );
  }

  checkCode(verCode: string): Observable<any> {

    console.log(verCode)
    return this._http.post<any>('http://localhost:8082/auth/checkCode', {
      email: localStorage.getItem('emailForReset'),
      code: verCode,
    });

  }

  resetPassword(newPassword: string): Observable<any> {
    return this._http.post<any>('http://localhost:8082/auth/resetPassword', {
      email: localStorage.getItem('emailForReset'),
      newPassword: newPassword,
    });
  }

  sendCode(email: string): Observable<any> {
    return this._http.post<any>('http://localhost:8082/auth/sendCode', email);
  }

  public get currentUserValue(): LogedUser {
    return this.currentUserSubject.value;
  }

  logout() {
    this.loginStatus.next(false);
    localStorage.removeItem('token');
    localStorage.removeItem('role');
    localStorage.removeItem('currentUser');
    localStorage.removeItem('email');
    this.router.navigate(['/']);

  }

  checkPasswordlessToken(token: string): Observable<any> {
    return this._http.get(`http://localhost:8082/auth/password-less-login/${token}`).pipe(
      map((response: any) => {
        if (response && response.token) {
          this.loginStatus.next(true);
          localStorage.setItem('token', response.token.accessToken);
          localStorage.setItem('currentUser', JSON.stringify(response));
          localStorage.setItem('role', response.role)
          localStorage.setItem('username', response.username)
          this.currentUserSubject.next(response);
        }
        return this.user;
      })
    );
  }

  sendLink(email: string): Observable<any> {
    console.log(email)
    return this._http.post(`http://localhost:8082/auth/password-less-login/`, email)
  }

  check2FAStatus(email: string): Observable<any> {

    return this._http.get(`http://localhost:8082/auth/two-factor-auth-status/` + email)
  }

  enable2FA(email: string, status: boolean): Observable<any> {
    console.log(email, status)
    return this._http.put(`http://localhost:8082/auth/two-factor-auth/`, {
      email,
      status
    })
  }
}