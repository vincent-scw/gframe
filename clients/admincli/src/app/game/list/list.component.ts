import { Component, OnInit, OnDestroy } from '@angular/core';
import { GameModel } from 'src/app/models/game.model';
import { GameService } from 'src/app/services/game.service';
import { Subscription } from 'rxjs';

@Component({
  selector: 'app-list',
  templateUrl: './list.component.html',
  styleUrls: ['./list.component.scss']
})
export class ListComponent implements OnInit, OnDestroy {
  games: GameModel[];
  displayedColumns: string[] = ['id', 'name', 'registerTime', 'startTime', 
    'completed', 'winner', 'cancelled'];

  private gameSub: Subscription;
  
  constructor(private gameSvc: GameService) { }

  ngOnInit() {
    this.gameSub = this.gameSvc.getGames().valueChanges.subscribe(({data}) => {
      this.games = data.getGames;
    });
  }

  ngOnDestroy() {
    if (!!this.gameSub) { this.gameSub.unsubscribe(); }
  }
}
