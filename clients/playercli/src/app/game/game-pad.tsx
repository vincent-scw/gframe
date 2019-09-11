import React from 'react';
import { gameService, authService } from '../services';
import { Card, Console } from '../game';
import './game-pad.scss';
import { Subscription } from 'rxjs';
import { GroupFormed, Player, GroupEvent } from '../services/server-events.model';

interface GamePadState {
  started: boolean;
  opponent: Player;
  group: GroupEvent | null;
}

export class GamePad extends React.Component<any, GamePadState> {
  groupSub: Subscription | null = null;
  connSub: Subscription | null = null;
  
  constructor(prop: any) {
    super(prop);

    this.state = { started: false, opponent: {id: "unknown", name: "unknown"}, group: null };

    this.onStart = this.onStart.bind(this);
  }

  onStart() {
    gameService.startListening();
  }

  componentDidMount() {
    this.groupSub = gameService.onGroup.subscribe(e => {
      if (e != null) {
        if (e.status === GroupFormed && gameService.opponents.length === 1) {
          this.setState({opponent: gameService.opponents[0], group: e});
        }
      }
    });
    this.connSub = gameService.connected.subscribe(c => this.setState({started: c}));
  }

  componentWillUnmount() {
    if (!!this.groupSub) this.groupSub.unsubscribe();
    if (!!this.connSub) this.connSub.unsubscribe();
  }

  render() {
    return (
      <div>
        {
          this.state.started ?
            <div>
              <div className="columns is-vcentered">
                <div className="column is-5">
                  <Card player={authService.user} group={this.state.group} readonly={false}/>
                </div>
                <div className="column has-text-centered"><strong>VS.</strong></div>
                <div className="column is-5">
                  <Card player={this.state.opponent} group={this.state.group} readonly={true}/>
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