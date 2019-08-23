import * as React from 'react';
import { gameService } from '../services';

interface ConsoleState {
  latestMsg: string;
}

export class Console extends React.Component<any, ConsoleState> {
  constructor(props: any) {
    super(props);
    this.state = { latestMsg: '' };
  }

  componentDidMount() {
    gameService.onMsg.subscribe(msg => {
      this.setState({latestMsg: msg});
    });  
  }

  componentWillUnmount() {
    gameService.onMsg.unsubscribe();
  }

  render() {
    return (
      <div>{this.state.latestMsg}</div>
    );
  }
}