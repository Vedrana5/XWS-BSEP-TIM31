import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { CompanyDto } from 'src/app/interfaces/company-dto';
import { LogedUser } from 'src/app/interfaces/loged-user';
import { CompanyService } from 'src/app/services/company.service';
import { UserService } from 'src/app/services/user.service';

@Component({
  selector: 'app-my-companies',
  templateUrl: './my-companies.component.html',
  styleUrls: ['./my-companies.component.css']
})
export class MyCompaniesComponent implements OnInit {

  items!: CompanyDto[];
  currentUser: LogedUser;
  constructor(
    private companyService: CompanyService,
    private userService: UserService,
    private router: ActivatedRoute,
    private route: Router,
  ) {
    this.currentUser = {} as LogedUser;
  }


  ngOnInit(): void {
    this.currentUser = this.userService.getValue();
    this.companyService.getAllUsersCompanies(this.currentUser.email).subscribe(res => {
      this.items = res;
      console.log(res);

    });
  }

  edit(id: any) {
    this.route.navigate(['/company/' + id]);
  }

}
