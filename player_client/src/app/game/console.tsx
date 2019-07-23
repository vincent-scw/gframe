import * as React from 'react';
import io from 'socket.io-client';

export class Console extends React.Component {
  private socket: any;
  state = {
    contents: []
  }

  constructor(props: any) {
    super(props);
  }

  componentDidMount() {
    this.socket = io('http://localhost:9010',
      {
        path: '/ws',
        'transports': [
          "websocket"
        ]
      });
    this.socket.on('console', (msg: string) => this.setState({ contents: <div>{msg}</div> }))
  }

  render() {
    return (
      <div>
        {this.state.contents}
      </div>
    );
  }
}