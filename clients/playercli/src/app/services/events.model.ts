export const GroupFormed = 2;

export enum Shape {
  NotSet = 0,
  Rock = 1,
  Paper = 2,
  Scissors = 3
}

export interface Player {
  id: string;
  name: string;
}

export interface GroupEvent {
  id: string;
  players: Player[];
  status: number;
}

export interface Move {
  group: string;
  shape: Shape;
}

export interface PlayEvent {
  winner: number;
  plays: Move[];
}