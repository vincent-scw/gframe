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
  isCompleted: boolean;
  isStarted: boolean;

  status?: string;
}

export interface GameResponse extends ResponseBase {
  getGame: GameModel;
}

export interface GameListResponse extends ResponseBase {
  getGames: GameModel[];
}