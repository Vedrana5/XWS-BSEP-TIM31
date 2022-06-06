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
    path: 'companies', canActivate: [AuthGuard],
    component: CompaniesListComponent
  },
  {
    path: 'companyRequests', canActivate: [AuthGuard],
    component: CompanyRequestsComponent
  },
  {
    path: 'ownerHome', canActivate: [AuthGuard],
    component: OwnerPageComponent
  },

];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
