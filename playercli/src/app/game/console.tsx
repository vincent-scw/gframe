import * as React from 'react';
import { Ws } from '../services/websocket';

interface ConsoleState {
  messages: string[];
}

export class Console extends React.Component<any, ConsoleState> {
  constructor(props: any) {
    super(props);
    this.state = { messages: [] };
  }

  componentDidMount() {
    const conn = new Ws("ws://localhost:9010/console");
    conn.On("console", (msg) => {
      console.log(msg)
      this.setState({ messages: this.state.messages.concat([msg]) })
    })
  }

  render() {
    return (
      <ul>
        {this.state.messages.map((msg, index) => <li key={index}>{msg}</li>)}
      </ul>
    );
  }
}