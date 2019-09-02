export interface GroupEvent {
  id: string;
  players: Player[];
  status: number;
}

export interface Player {
  id: string;
  name: string;
}

export const GroupFormed = 202;