import { Component, OnInit, OnDestroy } from '@angular/core';
import { GameModel } from 'src/app/models/game.model';
import { GameService } from 'src/app/services/game.service';
import { Subscription } from 'rxjs';
import { MatTableDataSource } from '@angular/material';

@Component({
  selector: 'app-list',
  templateUrl: './list.component.html',
  styleUrls: ['./list.component.scss']
})
export class ListComponent implements OnInit, OnDestroy {
  displayedColumns: string[] = ['id', 'name', 'registerTime', 'startTime', 
    'status', 'winner'];
  dataSource: MatTableDataSource<GameModel>;

  private gameSub: Subscription;
  
  constructor(private gameSvc: GameService) { }

  ngOnInit() {
    this.gameSub = this.gameSvc.getGames().valueChanges.subscribe(({data}) => {
      this.dataSource = new MatTableDataSource(data.getGames);
      this.dataSource.data.forEach(d => {
        if (d.isCancelled) {
          d.status = "Cancelled";
        } else if (d.isCompleted) {
          d.status = "Completed";
        } else if (d.isStarted) {
          d.status = "Started";
        } else {
          d.status = "Created";
        }
      });
      this.dataSource.filterPredicate = (data, filter): boolean => {
        return data.id.includes(filter) || data.name.includes(filter);
      };
    });
  }

  ngOnDestroy() {
    if (!!this.gameSub) { this.gameSub.unsubscribe(); }
  }

  applyFilter(filterValue: string) {
    this.dataSource.filter = filterValue.trim().toLowerCase();
  }
}
