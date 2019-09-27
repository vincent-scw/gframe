import { Component, OnInit, Input } from '@angular/core';
import { GameModel } from 'src/app/models/game.model';

@Component({
  selector: 'app-setting',
  templateUrl: './setting.component.html',
  styleUrls: ['./setting.component.scss']
})
export class SettingComponent implements OnInit {
  @Input() game: any = {};

  constructor() { 
  }

  ngOnInit() {
  }

}
