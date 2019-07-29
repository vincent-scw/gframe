import { Routes, RouterModule } from "@angular/router";
import { NgModule } from '@angular/core';

import { GameCenterComponent } from './game-center/game-center.component';

const gameRoutes: Routes = [
  {
    path: 'game',
    children: [
      {
        path: '',
        component: GameCenterComponent
      }
    ]
  }
]

@NgModule({
  imports: [
    RouterModule.forChild(gameRoutes)
  ],
  exports: [
    RouterModule
  ]
})
export class GameRoutingModule {}