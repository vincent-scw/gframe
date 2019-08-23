import React from 'react';
import { gameService } from '../services';
import { Console } from './console';
import './game-pad.scss';

interface GamePadState {
  started: boolean;
}

export class GamePad extends React.Component<any, GamePadState> {
  constructor(prop: any) {
    super(prop);

    this.state = { started: false };

    this.onStart = this.onStart.bind(this);
  }

  onStart() {
    gameService.startListening().then(
      () => this.setState({ started: true })
    );
  }

  render() {
    return (
      <div>
        {
          this.state.started ?
            <div>
              <Console />
            </div>
            :
            <div className="content has-text-centered start-game-div">
              <button className="button is-primary is-large" onClick={this.onStart}>Start Game</button>
            </div>
        }
      </div>
    );
  }
}