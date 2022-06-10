
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
  getAllPendingCompanies(): Observable<CompanyDto[]> {
    return this.http.get<CompanyDto[]>(`http://localhost:8082/companies/pendingCompanies`);
  }

  approveRequest(id: any) {
    return this.http.get<CompanyDto>(
      `http://localhost:8082/companies/approve/${id}`);
  }

  rejectRequest(id: any) {
    return this.http.get<CompanyDto>(
      `http://localhost:8082/companies/reject/${id}`);
  }

  getAllUsersCompanies(email: string): Observable<CompanyDto[]> {
    return this.http.get<CompanyDto[]>(`http://localhost:8082/companies/getAllByUser/${email}`);
  }

  createCompany(newCompany: CompanyDto) {
    return this.http.post(`http://localhost:8082/companies/createNew`, newCompany, {
      responseType: 'text',
    });
  }
}
