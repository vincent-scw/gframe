import { Component, OnInit, OnDestroy } from '@angular/core';
import { GameModel } from 'src/app/models/game.model';
import { ActivatedRoute, Router } from '@angular/router';
import { Subscription } from 'rxjs';
import { GameService } from 'src/app/services/game.service';

@Component({
  selector: 'app-game-center',
  templateUrl: './game-center.component.html',
  styleUrls: ['./game-center.component.scss']
})
export class GameCenterComponent implements OnInit, OnDestroy {
  game: any = {};
  isNew = false;
  showSimulator = false;

  private routeSub: Subscription;
  private gameSub: Subscription;

  constructor(
    private route: ActivatedRoute,
    private router: Router,
    private gameSvc: GameService
    ) { }

  ngOnInit() {
    this.routeSub = this.route.paramMap.subscribe(p => {
      const id = p.get('id');
      if (id === 'create') {
        this.game = {id: 'New Game', name: 'New Game'};
        this.isNew = true;
      } else {
        this.gameSub = this.gameSvc.getGame(id).valueChanges
          .subscribe(({data}) => {
            this.game = data.getGame;
            this.isNew = false;
            this.showSimulator = this.game.isStarted;
          }, err => console.error(err));
      }
    });
  }

  ngOnDestroy() {
    if (!!this.routeSub) this.routeSub.unsubscribe();
    if (!!this.gameSub) this.gameSub.unsubscribe();
  }

  submit() {
    if (this.isNew) {
      this.gameSvc.createGame(this.game.name).subscribe(({data}) => {
        this.router.navigateByUrl(`/game/${data.createGame.id}`);
      }, (error) => console.error(error));
    } else {

    }
  }
}
