import React from 'react';
import { gameService, authService } from '../services';
import { Card, Console } from '../game';
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
              <div className="columns is-vcentered">
                <div className="column is-5">
                  <Card player={authService.user.username}/>
                </div>
                <div className="column has-text-centered"><strong>VS.</strong></div>
                <div className="column is-5">
                  <Card player={"unknown"}/>
                </div>
              </div>
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