export enum Shape {
  NotSet = 0,
  Rock = 1,
  Paper = 2,
  Scissors = 3
}

export interface PlayEvent {
  player: string;
  group: string;
  shape: Shape;
}