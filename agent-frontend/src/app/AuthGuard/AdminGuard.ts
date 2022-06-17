import { Injectable } from "@angular/core";
import { CanActivate, Router } from "@angular/router";
import { UserService } from "../services/user.service";

@Injectable({ providedIn: 'root' })
export class AdminGuard implements CanActivate {
    constructor(
        private router: Router,
        private userService: UserService
    ) { }

    canActivate() {
        let role = localStorage.getItem('role');
        if (role === 'ADMIN') {
            return true;
        }

        alert('You do not have admin permissions.')
        this.userService.logout();
        this.router.navigate(['/']);
        return false;
    }
}
