import * as React from 'react';
import { gameService, authService } from '../services';
import { Subscription } from 'rxjs';
import { GroupFormed } from '../services/events.model';

interface ConsoleState {
  latestMsg: string;
}

export class Console extends React.Component<any, ConsoleState> {
  groupSub: Subscription | null = null;
  gameSub: Subscription | null = null;

  constructor(props: any) {
    super(props);
    this.state = { latestMsg: 'Waiting for your opponent(s)...' };
  }

  componentDidMount() {
    this.groupSub = gameService.onGroup.subscribe(e => {
      if (e != null) {
        switch (e.status)
        {
          case GroupFormed:
            let opponents = gameService.opponents.map(x => x.name);
            this.setState({latestMsg: 
              `Game start. Your opponent(s) is/are ${opponents.join(', ')}`});
            break;
          default:
            break;
        }
      }
    });
    this.gameSub = gameService.onGame.subscribe(e => {
      if (e != null) {
        if (e.winner === -1) {
          this.setState({latestMsg: "Draw Game."});
          return;
        }
        let winner = e.moves[e.winner];
        this.setState({latestMsg:
          winner.player.id === authService.user.id ?
          `You WIN!`: `You lose.`
        });
      }
    });
  }

  componentWillUnmount() {
    if (!!this.groupSub) this.groupSub.unsubscribe();
    if (!!this.gameSub) this.gameSub.unsubscribe();
  }

  render() {
    return (
      <div>{this.state.latestMsg}</div>
    );
  }
}