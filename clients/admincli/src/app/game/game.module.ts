import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { SharedModule } from '../shared.module';
import { SettingComponent } from './setting/setting.component';
import { SimulatorComponent } from './simulator/simulator.component';
import { GameCenterComponent } from './game-center/game-center.component';
import { GameRoutingModule } from './game-routing.module';
import { ListComponent } from './list/list.component';



@NgModule({
  declarations: [SettingComponent, SimulatorComponent, GameCenterComponent, ListComponent],
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
