import * as React from 'react';
import { gameService } from '../services';
import { Subscription } from 'rxjs';
import { GroupFormed } from '../services/server-events.model';

interface ConsoleState {
  latestMsg: string;
}

export class Console extends React.Component<any, ConsoleState> {
  groupSub: Subscription | null = null;

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
  }

  componentWillUnmount() {
    if (!!this.groupSub)
      this.groupSub.unsubscribe();
  }

  render() {
    return (
      <div>{this.state.latestMsg}</div>
    );
  }
}