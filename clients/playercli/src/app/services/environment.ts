declare var process : {
  env: {
    NODE_ENV: string
  }
}

export class environment {
  private apiUrl: string = "api.gframe.fun";

  get gameSvc(): string {
    if (process.env.NODE_ENV === 'production') {
      return `https://${this.apiUrl}/game`;
    } else {
      return 'http://localhost:8441';
    }
  }

  get wsGameSvc(): string {
    if (process.env.NODE_ENV === 'production') {
      return `wss://${this.apiUrl}/game`;
    } else {
      return 'ws://localhost:8441';
    }
  }
}