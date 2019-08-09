declare var process : {
  env: {
    NODE_ENV: string
  }
}

class environment {
  get authSvc(): string {
    if (process.env.NODE_ENV === 'production') {
      return 'http://localhost:8440';
    } else {
      return 'http://localhost:8440';
    }
  }

  get receptionSvc(): string {
    if (process.env.NODE_ENV === 'production') {
      return 'http://localhost:8441';
    } else {
      return 'http://localhost:8441';
    }
  }

  get notificationSvc(): string {
    if (process.env.NODE_ENV === 'production') {
      return 'ws://localhost:8442';
    } else {
      return 'ws://localhost:8442';
    }
  }
}

const env = new environment();
export default env;