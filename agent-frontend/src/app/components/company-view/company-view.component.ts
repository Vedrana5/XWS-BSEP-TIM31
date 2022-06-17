import { Component, Input, OnInit } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { MatDialog, MatDialogConfig } from '@angular/material/dialog';
import { ActivatedRoute, Router } from '@angular/router';
import { Subscription } from 'rxjs';
import { Comment } from 'src/app/interfaces/comment-dto';
import { CompanyDto } from 'src/app/interfaces/company-dto';
import { Interview } from 'src/app/interfaces/interview-dto';
import { JobOffer } from 'src/app/interfaces/jobOffer';
import { LogedUser } from 'src/app/interfaces/loged-user';
import { Salary } from 'src/app/interfaces/salaty-dto';
import { CompanyService } from 'src/app/services/company.service';
import { UserService } from 'src/app/services/user.service';
import { CommentComponent } from '../comment/comment.component';
import { InterviewComponent } from '../interview/interview.component';
import { JobOfferComponent } from '../job-offer/job-offer.component';
import { SalaryComponent } from '../salary/salary.component';


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
  comments: Comment[] = []
  salary: Salary[] = []
  interview: Interview[] = []
  constructor(private userService: UserService, private companyService: CompanyService, private router: Router, private route: ActivatedRoute, public matDialog: MatDialog) { }

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

    this.companyService.getCommentById(this.id).subscribe((data) => {
      this.comments = data;
      console.log(this.comments)
    });

    this.companyService.getSalaryById(this.id).subscribe((data) => {
      this.salary = data;
      console.log(this.jobOffers)
    });
    this.companyService.getInterviewById(this.id).subscribe((data) => {
      this.interview = data;
      console.log(this.jobOffers)
    });
  }

  createComment() {
    const dialogConfig = new MatDialogConfig();
    dialogConfig.disableClose = false;
    dialogConfig.id = 'modal-component';
    dialogConfig.height = 'fit-content';
    dialogConfig.width = '500px';
    const dialogRef = this.matDialog.open(CommentComponent, dialogConfig);
    dialogRef.afterClosed().subscribe({
      next: (res) => {
        console.log(res);
        this.comments = res.data;
      }
    })

  }

  createSallary() {

    const dialogConfig = new MatDialogConfig();
    dialogConfig.disableClose = false;
    dialogConfig.id = 'modal-component';
    dialogConfig.height = 'fit-content';
    dialogConfig.width = '500px';
    const dialogRef = this.matDialog.open(SalaryComponent, dialogConfig);
    dialogRef.afterClosed().subscribe({
      next: (res) => {
        console.log(res);
        this.comments = res.data;
      }
    })


  }

  createInterview() {
    const dialogConfig = new MatDialogConfig();
    dialogConfig.disableClose = false;
    dialogConfig.id = 'modal-component';
    dialogConfig.height = 'fit-content';
    dialogConfig.width = '500px';
    const dialogRef = this.matDialog.open(InterviewComponent, dialogConfig);
    dialogRef.afterClosed().subscribe({
      next: (res) => {
        console.log(res);
        this.comments = res.data;
      }
    })

  }

}
