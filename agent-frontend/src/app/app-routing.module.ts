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


const routes: Routes = [
  { path: '', component: LandingPageComponent },

  { path: 'registration', component: RegistrationComponent },
  {
    path: 'userHome',
    component: UserPageComponent, canActivate: [AuthGuard]
  },

  {
    path: 'adminHome',
    component: AdminPageComponent, canActivate: [AuthGuard]
  },
  {
    path: 'newRequest',
    component: NewRequestComponent, canActivate: [AuthGuard]
  },
  {
    path: 'companies', canActivate: [AuthGuard],
    component: CompaniesListComponent
  },
  {
    path: 'myCompanies', canActivate: [AuthGuard],
    component: MyCompaniesComponent
  },
  {
    path: 'companyRequests', canActivate: [AuthGuard],
    component: CompanyRequestsComponent
  },
  {
    path: 'ownerHome', canActivate: [AuthGuard],
    component: OwnerPageComponent
  },
  {
    path: "company/:id", canActivate: [AuthGuard],
    component: CompanyProfileComponent
  },

];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
