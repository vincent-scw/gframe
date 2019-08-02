import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { environment } from '../../environments/environment';

@Injectable({
  providedIn: 'root'
})
export class SimulatorService {
  private serviceUrl: string;

  constructor(private http: HttpClient) { 
    this.serviceUrl = `${environment.defaultProtocol}://${environment.services.admin}/api/simulator`;
  }

  injectPlayers(count: number) {
    return this.http.post(`${this.serviceUrl}/inject-players/${count}`, null);
  }
}
