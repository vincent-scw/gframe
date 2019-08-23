import { BehaviorSubject } from 'rxjs';
import * as neffos from 'neffos.js';
import { authService, env } from '../services';

export class GameService {
  private wsConn: neffos.Conn | null = null;
  onMsg = new BehaviorSubject<string>('');

  async startListening() {
    try {
      this.wsConn = await neffos.dial(`${env.wsGameSvc}/console?token=${authService.accessToken}`, {
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
          console: (nsConn: neffos.NSConn, msg: neffos.Message) => {
            console.log(msg.Body);
            this.onMsg.next(msg.Body);
          }
        }
      }, { // optional.
          reconnect: 2000,
          // set custom headers.
          headers: {
            // 'X-Username': 'kataras',
          }
        });

      await this.wsConn.connect("default");
    } catch (err) {
      console.error(err);
    }
  }

  close() {
    this.wsConn && this.wsConn.close();
  }
}