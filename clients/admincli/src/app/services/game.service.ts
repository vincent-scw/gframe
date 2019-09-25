import { Injectable } from '@angular/core';
import { Apollo } from 'apollo-angular';
import gql from 'graphql-tag';
import { Observable } from 'rxjs';
import { GameListResponse } from '../models/game.model';

const gameFragment = gql`
fragment GameInfo on GameType {
  id
  name
  createdBy
  createdTime
  registerTime
  startTime
  winner {
    id
    name
  }
  type
  isCancelled
}
`;

const getGames = gql`
query games($owner: String!) {
  getGames(owner: $owner) {
    ...GameInfo
  }
}
${gameFragment}
`;

@Injectable({
  providedIn: 'root'
})
export class GameService {

  constructor(private apollo: Apollo) { }

  getGames() {
    const owner = 'testg';
    return this.apollo.watchQuery<GameListResponse>({
      query: getGames,
      variables: {owner: owner}
    });
  }
}
