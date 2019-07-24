import * as React from 'react';

interface ConsoleState {
  messages: string[];
}

export class Console extends React.Component<any, ConsoleState> {
  constructor(props: any) {
    super(props);
    this.state = { messages: [] };
  }

  componentDidMount() {
    const conn = new WebSocket("ws://localhost:9010/ws/");
    conn.onclose = (evt) => {
      console.log("connection closed.")
    }
    conn.onmessage = (evt) => {
      this.setState({ messages: this.state.messages.concat([evt.data]) });
    }
  }

  render() {
    return (
      <ul>
        {this.state.messages.map((msg, index) => <li key={index}>{msg}</li>)}
      </ul>
    );
  }
}