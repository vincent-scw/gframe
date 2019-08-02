import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { ConsoleComponent } from './console/console.component';
import { SharedModule } from '../shared.module';
import { GameModule } from '../game/game.module';

@NgModule({
  declarations: [ConsoleComponent],
  imports: [
    CommonModule,
    SharedModule,
    GameModule
  ],
  exports: [
    ConsoleComponent
  ]
})
export class LayoutModule { }
