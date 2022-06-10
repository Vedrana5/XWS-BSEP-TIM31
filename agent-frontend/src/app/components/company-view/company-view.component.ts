import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { CompanyDto } from 'src/app/interfaces/company-dto';
import { JobOffer } from 'src/app/interfaces/jobOffer';
import { LogedUser } from 'src/app/interfaces/loged-user';
import { CompanyService } from 'src/app/services/company.service';
import { UserService } from 'src/app/services/user.service';
import { JobOfferComponent } from '../job-offer/job-offer.component';

@Component({
  selector: 'app-company-view',
  templateUrl: './company-view.component.html',
  styleUrls: ['./company-view.component.css']
})
export class CompanyViewComponent implements OnInit {

  company!: CompanyDto;
  currentUser!: LogedUser;
  role!: string;
  id!: any
  jobOffers: JobOffer[] = []
  constructor(private userService: UserService, private companyService: CompanyService, private router: Router, private route: ActivatedRoute) { }

  ngOnInit(): void {
    this.currentUser = this.userService.getValue();
    this.role = this.currentUser.role;
    this.id = +this.route.snapshot.paramMap.get('id')!;
    this.companyService.getCompanyById(this.id).subscribe((data) => {
      this.company = data;
      console.log(this.company.id)
    });
    this.companyService.getOfferbyCompany(this.id).subscribe((data) => {
      this.jobOffers = data;
      console.log(this.jobOffers)
    });
  }

}
