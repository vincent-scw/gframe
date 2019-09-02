import React from 'react';
import { gameService, authService } from '../services';
import { Card, Console } from '../game';
import './game-pad.scss';
import { Subscription } from 'rxjs';
import { GroupFormed, Player } from '../services/server-events.model';

interface GamePadState {
  started: boolean;
  opponent: Player;
  groupId: string;
}

export class GamePad extends React.Component<any, GamePadState> {
  groupSub: Subscription | null = null;
  
  constructor(prop: any) {
    super(prop);

    this.state = { started: false, opponent: {id: "unknown", name: "unknown"}, groupId: '' };

    this.onStart = this.onStart.bind(this);
  }

  onStart() {
    gameService.startListening().then(
      () => this.setState({ started: true })
    );
  }

  componentDidMount() {
    this.groupSub = gameService.onGroup.subscribe(e => {
      if (e != null) {
        if (e.status === GroupFormed && gameService.opponents.length === 1) {
          this.setState({opponent: gameService.opponents[0], groupId: e.id});
        }
      }
    });
  }

  componentWillUnmount() {
    if (!!this.groupSub)
      this.groupSub.unsubscribe();
  }

  render() {
    return (
      <div>
        {
          this.state.started ?
            <div>
              <div className="columns is-vcentered">
                <div className="column is-5">
                  <Card player={authService.user} groupId={this.state.groupId} readonly={false}/>
                </div>
                <div className="column has-text-centered"><strong>VS.</strong></div>
                <div className="column is-5">
                  <Card player={this.state.opponent} groupId={this.state.groupId} readonly={true}/>
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