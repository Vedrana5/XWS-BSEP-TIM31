import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { LandingPageComponent } from './components/LandingPage/landing-page/landing-page.component';
import { RegistrationComponent } from './components/Registration/registration/registration.component';
import { AdminComponent } from './components/admin/admin/admin.component';
import { ClientComponent } from './components/client/client/client.component';
import { ChangePasswordComponent } from './components/change-password/change-password/change-password.component';
import { AuthGuard } from './AuthGuard/AuthGuard';
import { ResetPasswordComponent } from './components/reset-password/reset-password/reset-password.component';

const routes: Routes = [
  { path: '', component: LandingPageComponent },
  { path: 'resetPassword', component: ResetPasswordComponent },
  { path: 'registration', component: RegistrationComponent },
  {
    path: 'ahome',
    component: AdminComponent, canActivate: [AuthGuard]
  },
  {
    path: 'chome',
    component: ClientComponent, canActivate: [AuthGuard],
    children: [

      { path: 'changePassword', component: ChangePasswordComponent, canActivate: [AuthGuard] },
    ],
  }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
