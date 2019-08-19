declare var process : {
  env: {
    NODE_ENV: string
  }
}

class environment {
  private apiUrl: string = "gframe-api.eastasia.cloudapp.azure.com:90";

  get authSvc(): string {
    if (process.env.NODE_ENV === 'production') {
      return `http://${this.apiUrl}/oauth`;
    } else {
      return 'http://localhost:8440';
    }
  }

  get receptionSvc(): string {
    if (process.env.NODE_ENV === 'production') {
      return `http://${this.apiUrl}/reception`;
    } else {
      return 'http://localhost:8441';
    }
  }

  get notificationSvc(): string {
    if (process.env.NODE_ENV === 'production') {
      return `ws://${this.apiUrl}/notification`;
    } else {
      return 'ws://localhost:8442';
    }
  }
}

const env = new environment();
export default env;