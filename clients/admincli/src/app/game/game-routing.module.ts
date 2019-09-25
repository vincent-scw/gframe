import { Routes, RouterModule } from "@angular/router";
import { NgModule } from '@angular/core';

import { GameCenterComponent } from './game-center/game-center.component';
import { ListComponent } from './list/list.component';

const gameRoutes: Routes = [
  {
    path: 'game',
    children: [
      {
        path: 'list',
        component: ListComponent
      },
      {
        path: ':id',
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