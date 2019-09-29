import { Injectable } from '@angular/core';
import { Apollo } from 'apollo-angular';
import gql from 'graphql-tag';
import { Observable } from 'rxjs';
import { GameListResponse, GameResponse, GameModel } from '../models/game.model';

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

const getGame = gql`
query games($id: String!) {
  getGame(id: $id) {
    ...GameInfo
  }
}
${gameFragment}
`;

const createGame = gql`
mutation games($game: GameInputType!) {
  createGame(game: $game) {
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

  getGame(id: string) {
    return this.apollo.watchQuery<GameResponse>({
      query: getGame,
      variables: {id: id}
    });
  }

  createGame(name: string) {
    const parameters = {
      game: {
        name: name,
        createdBy: 'testg'
      }
    }
    return this.apollo.mutate({
      mutation: createGame,
      variables: parameters
    });
  }
}
