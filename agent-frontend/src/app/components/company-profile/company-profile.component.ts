import { Component, Input, OnInit } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { MatDialog, MatDialogConfig } from '@angular/material/dialog';
import { ActivatedRoute, Router } from '@angular/router';
import { Subscription } from 'rxjs';
import { CompanyDto } from 'src/app/interfaces/company-dto';
import { JobOffer } from 'src/app/interfaces/jobOffer';
import { CompanyService } from 'src/app/services/company.service';
import { JobOfferComponent } from '../job-offer/job-offer.component';

@Component({
  selector: 'app-company-profile',
  templateUrl: './company-profile.component.html',
  styleUrls: ['./company-profile.component.css']
})
export class CompanyProfileComponent implements OnInit {

  company!: CompanyDto;
  sub!: Subscription;
  @Input()
  userDetails: any
  initialDetails: any
  editMode = false
  id!: any
  jobOffers!: JobOffer[]



  constructor(
    private companyService: CompanyService,
    private router: ActivatedRoute,
    private route: Router,
    public matDialog: MatDialog
  ) {
    this.company = {} as CompanyDto
  }

  ngOnInit(): void {
    this.id = +this.router.snapshot.paramMap.get('id')!;
    this.companyService.getCompanyById(this.id).subscribe((data) => {
      this.company = data;
      console.log(this.company.id)
    });
  }

  submit() {
    this.companyService.editCompany(this.company).subscribe((data) => {

      console.log(this.company)
      this.route.navigate(['/myCompanies']);
    });
  }
  cancel() {
    this.route.navigate(['/myCompanies']);
  }

  openModal() {
    const dialogConfig = new MatDialogConfig();
    dialogConfig.disableClose = false;
    dialogConfig.id = 'modal-component';
    dialogConfig.height = 'fit-content';
    dialogConfig.width = '500px';
    const dialogRef = this.matDialog.open(JobOfferComponent, dialogConfig);
    dialogRef.afterClosed().subscribe({
      next: (res) => {
        console.log(res);
        this.jobOffers = res.data;
      }
    })
  }

}
