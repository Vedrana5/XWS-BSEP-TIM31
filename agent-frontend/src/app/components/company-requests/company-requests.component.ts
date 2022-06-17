import { Component, OnInit } from '@angular/core';
import { MatSnackBar } from '@angular/material/snack-bar';
import { CompanyDto } from 'src/app/interfaces/company-dto';
import { CompanyService } from 'src/app/services/company.service';

@Component({
  selector: 'app-company-requests',
  templateUrl: './company-requests.component.html',
  styleUrls: ['./company-requests.component.css']
})
export class CompanyRequestsComponent implements OnInit {

  items!: CompanyDto[];
  constructor(private companyService: CompanyService, private matSnackBar: MatSnackBar) { }

  ngOnInit(): void {
    this.companyService.getAllPendingCompanies().subscribe((data) => { this.items = data });


  }

  approveRequest(id: number) {

    this.companyService.approveRequest(id).subscribe(data => {
      console.log(id)

      window.location.reload();

    });

  }

  rejectRequest(id: number) {
    this.companyService.rejectRequest(id).subscribe(data => {
      console.log(id)

      window.location.reload();

    });
  }


}
