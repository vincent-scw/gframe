import { Component, OnInit, OnDestroy } from '@angular/core';
import { GameModel } from 'src/app/models/game.model';
import { ActivatedRoute } from '@angular/router';
import { Subscription } from 'rxjs';
import { GameService } from 'src/app/services/game.service';

@Component({
  selector: 'app-game-center',
  templateUrl: './game-center.component.html',
  styleUrls: ['./game-center.component.scss']
})
export class GameCenterComponent implements OnInit, OnDestroy {
  game: any = {};
  showSimulator = false;

  private gameSub: Subscription;

  constructor(
    private route: ActivatedRoute,
    private gameSvc: GameService
    ) { }

  ngOnInit() {
    const id = this.route.snapshot.paramMap.get('id');
    if (id === 'create') {
      this.game = {id: 'New Game', name: 'New Game'};
    } else {
      this.gameSub = this.gameSvc.getGame(id).valueChanges
        .subscribe(({data}) => {
          this.game = data.getGame;
          this.showSimulator = true;
        }, err => console.error(err));
    }
  }

  ngOnDestroy() {
    if (!!this.gameSub) this.gameSub.unsubscribe();
  }
}
