
import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';

import { environment } from 'src/environments/environment';
import { CompanyDto } from '../interfaces/company-dto';


@Injectable({
  providedIn: 'root'
})
export class CompanyService {

  constructor(private http: HttpClient) { }

  getAllCompanies(): Observable<CompanyDto[]> {
    return this.http.get<CompanyDto[]>(`http://localhost:8082/companies/getAllforUser`);

  }
}
