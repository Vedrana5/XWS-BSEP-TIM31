
import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';

import { environment } from 'src/environments/environment';
import { Comment } from '../interfaces/comment-dto';
import { CompanyDto } from '../interfaces/company-dto';
import { Interview } from '../interfaces/interview-dto';
import { JobOffer } from '../interfaces/jobOffer';
import { Salary } from '../interfaces/salaty-dto';


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
  getCompanyById(id: any) {

    return this.http.get<CompanyDto>(`http://localhost:8082/companies/findCompany/${id}`);
  }
  editCompany(company: CompanyDto) {

    return this.http.put<CompanyDto>(`http://localhost:8082/companies/editCompany`, company);
  }

  createOffer(offer: JobOffer) {
    console.log(offer)
    return this.http.post<JobOffer>(`http://localhost:8082/offers/newOffer`, offer);

  }
  getOfferbyCompany(id: any) {

    return this.http.get<JobOffer[]>(`http://localhost:8082/offers/getAllByCompany/${id}`);

  }

  getCommentById(id: any) {
    return this.http.get<Comment[]>(`http://localhost:8082/comments/getAllByCompany/${id}`);

  }

  getSalaryById(id: any) {
    return this.http.get<Salary[]>(`http://localhost:8082/comments/getAllSalaryByCompany/${id}`);


  }

  getInterviewById(id: any) {
    return this.http.get<Interview[]>(`http://localhost:8082/comments/getAllInterviewByCompany/${id}`);


  }

  createComment(newComment: Comment) {
    return this.http.post(`http://localhost:8082/comments/createNewComment`, newComment, {
      responseType: 'text',
    });


  }

  createInterview(newInterview: Interview) {
    console.log(newInterview)
    return this.http.post(`http://localhost:8082/comments/createInterview`, newInterview, {
      responseType: 'text',
    });


  }

  createSallary(newSalary: Salary) {
    console.log(newSalary)
    return this.http.post(`http://localhost:8082/comments/createSalary`, newSalary, {
      responseType: 'text',
    });


  }



}

