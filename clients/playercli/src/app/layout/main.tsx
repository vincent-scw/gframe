import React from 'react';
import { Console } from '../game';

interface MainState {
  showConsole: boolean;
}

export class Main extends React.Component<any, MainState> {
  constructor(props: any) {
    super(props);

    this.state = { showConsole: false }
    this.onStart = this.onStart.bind(this);
    this.onExit = this.onExit.bind(this);
  }

  render() {
    return (
      <div>
        <button className="button is-primary" onClick={this.onStart}>Start</button>
        <button className="button" onClick={this.onExit}>Exit</button>
        {this.state.showConsole &&
          <div>
            <Console></Console>
          </div>
        }
      </div>
    );
  }

  onStart() {
    this.setState({showConsole: true});
  }

  onExit() {
    this.setState({showConsole: false});
  }
}