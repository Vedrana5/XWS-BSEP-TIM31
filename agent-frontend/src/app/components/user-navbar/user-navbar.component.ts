import { Component, OnInit } from '@angular/core';
import { LogedUser } from 'src/app/interfaces/loged-user';
import { UserService } from 'src/app/services/user.service';

@Component({
  selector: 'app-user-navbar',
  templateUrl: './user-navbar.component.html',
  styleUrls: ['./user-navbar.component.css']
})
export class UserNavbarComponent implements OnInit {

  role!: string;
  currentUser!: LogedUser;
  constructor(private userService: UserService) { }

  ngOnInit(): void {
    this.currentUser = this.userService.getValue();
    this.role = this.currentUser.role;
    console.log(this.role)
  }

  logout() {
    this.userService.logout();
  }
}
