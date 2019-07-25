import * as React from 'react';
import * as neffos from 'neffos.js';

interface ConsoleState {
  messages: string[];
}

export class Console extends React.Component<any, ConsoleState> {
  constructor(props: any) {
    super(props);
    this.state = { messages: [] };
  }

  async componentDidMount() {
    try {
      const conn = await neffos.dial("ws://localhost:9010/console", {
        default: { // "default" namespace.
          _OnNamespaceConnected: (nsConn: neffos.NSConn, msg: neffos.Message) => {
            if (nsConn.conn.wasReconnected()) {
              console.log("re-connected after " + nsConn.conn.reconnectTries.toString() + " trie(s)");
            }
            console.log("connected to namespace: " + msg.Namespace);
          },
          _OnNamespaceDisconnect: (nsConn: neffos.NSConn, msg: neffos.Message) => {
            console.log("disconnected from namespace: " + msg.Namespace);
          },
          console: (nsConn: neffos.NSConn, msg: neffos.Message) => { // "chat" event.
            console.log(msg.Body);
            this.setState({ messages: this.state.messages.concat([msg.Body]) })
          }
        }
      }, { // optional.
          reconnect: 2000,
          // set custom headers.
          headers: {
            // 'X-Username': 'kataras',
          }
        });

      await conn.connect("default");
    } catch (err) {
      console.error(err);
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