import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { ConsoleComponent } from './console/console.component';
import { SharedModule } from '../shared.module';
import { GameModule } from '../game/game.module';
import { FooterComponent } from './footer/footer.component';

@NgModule({
  declarations: [ConsoleComponent, FooterComponent],
  imports: [
    CommonModule,
    SharedModule,
    GameModule
  ],
  exports: [
    ConsoleComponent,
    FooterComponent
  ]
})
export class LayoutModule { }
