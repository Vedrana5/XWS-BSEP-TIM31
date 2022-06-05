import { Component, OnInit } from '@angular/core';
import { SubjectData } from 'src/app/interfaces/subject-data';

@Component({
  selector: 'app-user-page',
  templateUrl: './user-page.component.html',
  styleUrls: ['./user-page.component.css']
})
export class UserPageComponent implements OnInit {

  userInfo!: SubjectData;
  constructor() { }

  ngOnInit(): void {
  }

}
