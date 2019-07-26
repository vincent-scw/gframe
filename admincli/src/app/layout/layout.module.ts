import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { ConsoleComponent } from './console/console.component';
import { SharedModule } from '../shared.module';

@NgModule({
  declarations: [ConsoleComponent],
  imports: [
    CommonModule,
    SharedModule
  ],
  exports: [
    ConsoleComponent
  ]
})
export class LayoutModule { }
