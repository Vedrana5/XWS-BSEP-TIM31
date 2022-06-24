import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { LandingPageComponent } from './components/LandingPage/landing-page/landing-page.component';
import { RegistrationComponent } from './components/Registration/registration/registration.component';
import { AuthGuard } from './AuthGuard/AuthGuard';
import { UserPageComponent } from './components/user-page/user-page.component';
import { AdminPageComponent } from './components/admin-page/admin-page.component';
import { CompaniesListComponent } from './components/companies-list/companies-list.component';
import { OwnerNavbarComponent } from './components/owner-navbar/owner-navbar.component';
import { OwnerPageComponent } from './components/owner-page/owner-page.component';
import { CompanyRequestsComponent } from './components/company-requests/company-requests.component';
import { MyCompaniesComponent } from './components/my-companies/my-companies.component';
import { NewRequestComponent } from './components/new-request/new-request.component';
import { CompanyProfileComponent } from './components/company-profile/company-profile.component';
import { CompanyViewComponent } from './components/company-view/company-view.component';
import { ResetPasswordComponent } from './components/comment/reset-password/reset-password.component';
import { OwnerGuard } from './AuthGuard/OwnerGuard';
import { AdminGuard } from './AuthGuard/AdminGuard';
import { UserGuard } from './AuthGuard/UserGuard';
import { PasswordlessLoginComponent } from './components/passwordless-login/passwordless-login/passwordless-login.component';


const routes: Routes = [
  { path: '', component: LandingPageComponent },
  {
    path: "passwordless-login/:token",
    component: PasswordlessLoginComponent,
  }
  ,
  { path: 'registration', component: RegistrationComponent },
  {
    path: 'userHome',
    component: UserPageComponent, canActivate: [UserGuard]
  },
  { path: 'resetPassword', component: ResetPasswordComponent },


  {
    path: 'adminHome',
    component: AdminPageComponent, canActivate: [AdminGuard]
  },
  {
    path: 'newRequest',
    component: NewRequestComponent, canActivate: [UserGuard]
  },
  {
    path: 'companies', canActivate: [AuthGuard],
    component: CompaniesListComponent
  },
  {
    path: 'myCompanies', canActivate: [OwnerGuard],
    component: MyCompaniesComponent
  },
  {
    path: 'companyRequests', canActivate: [AdminGuard],
    component: CompanyRequestsComponent
  },
  {
    path: 'ownerHome', canActivate: [OwnerGuard],
    component: OwnerPageComponent
  },
  {
    path: "company/:id", canActivate: [AuthGuard],
    component: CompanyProfileComponent
  },
  {
    path: "companyView/:id", canActivate: [AuthGuard],
    component: CompanyViewComponent
  },


];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
