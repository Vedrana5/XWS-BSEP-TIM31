import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { JobOfferPublish } from '../interfaces/jobOfferPublish-dto';

@Injectable({
  providedIn: 'root'
})
export class JobOfferService {

  constructor(private http: HttpClient) { }

  publishJobOfferOnDislinkt(jobOffer: JobOfferPublish): Observable<any> {
    return this.http.post('http://localhost:9090/users/share/jobOffer', jobOffer);
  }
}
