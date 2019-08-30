import * as React from 'react';
import { gameService } from '../services';
import { Subscription } from 'rxjs';

interface ConsoleState {
  latestMsg: string;
}

export class Console extends React.Component<any, ConsoleState> {
  msgSub: Subscription | null = null;

  constructor(props: any) {
    super(props);
    this.state = { latestMsg: '' };
  }

  componentDidMount() {
    this.msgSub = gameService.onMsg.subscribe(msg => {
      this.setState({latestMsg: msg});
    });
  }

  componentWillUnmount() {
    if (!!this.msgSub)
      this.msgSub.unsubscribe();
  }

  render() {
    return (
      <div>{this.state.latestMsg}</div>
    );
  }
}