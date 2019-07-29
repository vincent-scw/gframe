import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { SharedModule } from '../shared.module';
import { SettingComponent } from './setting/setting.component';
import { SimulatorComponent } from './simulator/simulator.component';
import { GameCenterComponent } from './game-center/game-center.component';
import { GameRoutingModule } from './game-routing.module';



@NgModule({
  declarations: [SettingComponent, SimulatorComponent, GameCenterComponent],
  imports: [
    CommonModule,
    SharedModule,
    GameRoutingModule
  ],
  exports: [
    GameCenterComponent
  ]
})
export class GameModule { }
