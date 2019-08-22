import { Injectable } from '@angular/core';
import { Apollo } from 'apollo-angular';
import gql from 'graphql-tag';
import { Observable } from 'rxjs';

const InjectPlayers = gql`
mutation simulator($amount: Int!) {
  injectPlayers(amount: $amount) {
    amount
  }
}
`

@Injectable({
  providedIn: 'root'
})
export class SimulatorService {

  constructor(private apollo: Apollo) { }

  injectPlayers(amount: number): Observable<any> {
    return this.apollo.mutate({
      mutation: InjectPlayers,
      variables: { amount }
    });
  }
}
