import { Routes, RouterModule } from "@angular/router";
import { NgModule } from '@angular/core';
import { LoginComponent } from './login/login.component';

const accountRoutes: Routes = [
  {path: 'login', component: LoginComponent }
];

@NgModule({
  imports: [
    RouterModule.forChild(accountRoutes)
  ],
  exports: [
    RouterModule
  ]
})
export class AccountRoutingModule {}