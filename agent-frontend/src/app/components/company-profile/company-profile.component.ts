import { Component, Input, OnInit } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { MatDialog, MatDialogConfig } from '@angular/material/dialog';
import { ActivatedRoute, Router } from '@angular/router';
import { Subscription } from 'rxjs';
import { CompanyDto } from 'src/app/interfaces/company-dto';
import { CompanyService } from 'src/app/services/company.service';

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


  constructor(
    private companyService: CompanyService,
    private router: ActivatedRoute,
    private route: Router,
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



}
