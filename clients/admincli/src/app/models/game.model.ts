import { ResponseBase } from './response-base.model';

export interface GameModel {
  id: string;
  name: string;
  createdBy: string;
  createdTime: Date;
  registerTime?: Date;
  startTime?: Date;
  winner: Player;
  type: number;
  isCancelled: boolean;
}

export interface GameListResponse extends ResponseBase {
  getGames: GameModel[];
}