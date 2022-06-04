import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { BehaviorSubject, Observable, Subject } from 'rxjs';
import { LogedUser } from 'src/app/interfaces/loged-user';

import { SubjectData } from 'src/app/interfaces/subject-data';
import { map } from 'rxjs/operators';

@Injectable({
  providedIn: 'root',
})
export class UserService {
  private currentUserSubject: BehaviorSubject<LogedUser>;
  public currentUser: Observable<LogedUser>;
  private user!: LogedUser;





  constructor(private _http: HttpClient) {
    this.currentUserSubject = new BehaviorSubject<LogedUser>(
      JSON.parse(localStorage.getItem('currentUser')!)
    );
    this.currentUser = this.currentUserSubject.asObservable();
  }

  createSubject(newSubject: SubjectData): Observable<any> {
    console.log(newSubject.username)
    return this._http.post<any>(
      'http://localhost:8082/auth/register',
      newSubject
    );
  }

  changePassword(data: any) {
    console.log(data)
    return this._http.put(`http://localhost:8081/api/user/changePassword`, data);
  }
  login(model: any): Observable<LogedUser> {
    return this._http.post(`http://localhost:8081/api/user/login`, model).pipe(
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

  checkCode(verCode: string): Observable<any> {
    return this._http.post<any>('http://localhost:8081/api/user/checkCode', {
      email: localStorage.getItem('emailForReset'),
      code: verCode,
    });
  }

  resetPassword(newPassword: string): Observable<any> {
    return this._http.post<any>('http://localhost:8081/api/user/resetPassword', {
      email: localStorage.getItem('emailForReset'),
      newPassword: newPassword,
    });
  }

  sendCode(email: string): Observable<any> {
    return this._http.post<any>('http://localhost:8081/api/user/sendCode', email);
  }

  public get currentUserValue(): LogedUser {
    return this.currentUserSubject.value;
  }
}