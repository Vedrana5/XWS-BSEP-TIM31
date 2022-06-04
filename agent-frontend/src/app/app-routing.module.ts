import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { LandingPageComponent } from './components/LandingPage/landing-page/landing-page.component';
import { RegistrationComponent } from './components/Registration/registration/registration.component';
import { AuthGuard } from './AuthGuard/AuthGuard';


const routes: Routes = [
  { path: '', component: LandingPageComponent },

  { path: 'registration', component: RegistrationComponent },

];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
