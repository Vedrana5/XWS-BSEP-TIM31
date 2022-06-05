import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { LandingPageComponent } from './components/LandingPage/landing-page/landing-page.component';
import { RegistrationComponent } from './components/Registration/registration/registration.component';
import { AuthGuard } from './AuthGuard/AuthGuard';
import { UserPageComponent } from './components/user-page/user-page.component';
import { AdminPageComponent } from './components/admin-page/admin-page.component';


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

];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
