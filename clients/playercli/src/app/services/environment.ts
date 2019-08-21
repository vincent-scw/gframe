declare var process : {
  env: {
    NODE_ENV: string
  }
}

class environment {
  private apiUrl: string = "api.gframe.fun";

  get authSvc(): string {
    if (process.env.NODE_ENV === 'production') {
      return `https://${this.apiUrl}/oauth`;
    } else {
      return 'http://localhost:8440';
    }
  }

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

const env = new environment();
export default env;