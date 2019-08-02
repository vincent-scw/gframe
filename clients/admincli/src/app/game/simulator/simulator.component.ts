import { Component, OnInit } from '@angular/core';
import { SimulatorService } from 'src/app/services/simulator.service';

@Component({
  selector: 'app-simulator',
  templateUrl: './simulator.component.html',
  styleUrls: ['./simulator.component.scss'],
  providers: [SimulatorService]
})
export class SimulatorComponent implements OnInit {
  amount: number;

  constructor(private simulatorSvc: SimulatorService) { }

  ngOnInit() {
  }

  inject() {
    this.simulatorSvc.injectPlayers(this.amount).toPromise();
  }
}
