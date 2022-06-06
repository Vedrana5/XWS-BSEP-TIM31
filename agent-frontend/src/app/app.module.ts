import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { DatePipe } from '@angular/common';

import { HttpClientModule, HTTP_INTERCEPTORS } from '@angular/common/http';
import { NgbModule } from '@ng-bootstrap/ng-bootstrap';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { MatCardModule } from '@angular/material/card';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { LandingPageComponent } from './components/LandingPage/landing-page/landing-page.component';
import { MaterialModule } from 'src/material.module';
import { JwtInterceptor } from './JwtInterceptor/jwt-interceptor';
import { RegistrationComponent } from './components/Registration/registration/registration.component';
import { UserPageComponent } from './components/user-page/user-page.component';
import { AdminPageComponent } from './components/admin-page/admin-page.component';
import { UserNavbarComponent } from './components/user-navbar/user-navbar.component';
import { AdminNavbarComponent } from './components/admin-navbar/admin-navbar.component';
import { OwnerPageComponent } from './components/owner-page/owner-page.component';
import { OwnerNavbarComponent } from './components/owner-navbar/owner-navbar.component';
import { CompaniesListComponent } from './components/companies-list/companies-list.component';
import { CompanyRequestsComponent } from './components/company-requests/company-requests.component';
import { MyCompaniesComponent } from './components/my-companies/my-companies.component';


@NgModule({
  declarations: [
    AppComponent,
    LandingPageComponent,
    RegistrationComponent,
    UserPageComponent,
    AdminPageComponent,
    UserNavbarComponent,
    AdminNavbarComponent,
    OwnerPageComponent,
    OwnerNavbarComponent,
    CompaniesListComponent,
    CompanyRequestsComponent,
    MyCompaniesComponent,

  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    HttpClientModule,
    FormsModule,
    ReactiveFormsModule,
    NgbModule,
    BrowserAnimationsModule,
    MaterialModule,
    MatCardModule,
  ],
  providers: [HttpClientModule,
    { provide: HTTP_INTERCEPTORS, useClass: JwtInterceptor, multi: true }, DatePipe],
  bootstrap: [AppComponent]
})
export class AppModule { }
