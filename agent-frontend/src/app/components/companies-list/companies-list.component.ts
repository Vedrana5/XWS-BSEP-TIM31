import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { CompanyDto } from 'src/app/interfaces/company-dto';
import { LogedUser } from 'src/app/interfaces/loged-user';
import { CompanyService } from 'src/app/services/company.service';
import { UserService } from 'src/app/services/user.service';

@Component({
  selector: 'app-companies-list',
  templateUrl: './companies-list.component.html',
  styleUrls: ['./companies-list.component.css']
})
export class CompaniesListComponent implements OnInit {

  currentUser!: LogedUser;
  role!: string;
  items: CompanyDto[] = [];
  constructor(private companyService: CompanyService,
    private userService: UserService,
    private router: Router
  ) { }

  ngOnInit(): void {
    this.currentUser = this.userService.getValue();
    this.role = this.currentUser.role;
    console.log(this.role);
    this.companyService.getAllCompanies().subscribe((data) => { this.items = data });

  }

  see(id: any) {
    this.router.navigate(['/companyView/' + id]);
  }

}
