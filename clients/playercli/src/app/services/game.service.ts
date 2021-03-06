import { BehaviorSubject } from 'rxjs';
import * as neffos from 'neffos.js';
import { authService, env } from '../services';
import { GroupEvent, Player, PlayEvent } from './events.model';

export class GameService {
  private wsConn: neffos.Conn | null = null;
  onGroup = new BehaviorSubject<GroupEvent | null>(null);
  onPlayer = new BehaviorSubject<null>(null);
  onGame = new BehaviorSubject<PlayEvent | null>(null);
  connected = new BehaviorSubject<boolean>(false);

  private _opponents: Player[] = [];
  get opponents(): Player[] {
    return this._opponents;
  }

  async startListening() {
    try {
      this.wsConn = await neffos.dial(`${env.wsGameSvc}/console?token=${authService.accessToken}`, {
        default: { // "default" namespace.
          _OnNamespaceConnected: (nsConn: neffos.NSConn, msg: neffos.Message) => {
            if (nsConn.conn.wasReconnected()) {
              console.log("re-connected after " + nsConn.conn.reconnectTries.toString() + " trie(s)");
            }
            console.log("connected to namespace: " + msg.Namespace);
            this.connected.next(true);
          },
          _OnNamespaceDisconnect: (nsConn: neffos.NSConn, msg: neffos.Message) => {
            console.log("disconnected from namespace: " + msg.Namespace);
            this.connected.next(false);
          },
          group: (nsConn: neffos.NSConn, msg: neffos.Message) => {
            console.log(msg.Body);
            // Extract opponents
            let ge = JSON.parse(msg.Body) as GroupEvent;
            this._opponents = ge.players.filter(x => x.id !== authService.user.id);
            this.onGroup.next(ge);
          },
          player: (nsConn: neffos.NSConn, msg: neffos.Message) => {
            console.log(msg.Body);
            this.onPlayer.next(JSON.parse(msg.Body));
          },
          game: (nsConn: neffos.NSConn, msg: neffos.Message) => {
            console.log(msg.Body);
            this.onGame.next(JSON.parse(msg.Body));
          }
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

  ask(event: string, content: any): null | Promise<void | neffos.Message> {
    if (this.wsConn && !this.wsConn.isClosed()) {
      let msg = new neffos.Message();
      msg.Namespace = "default";
      msg.Event = event;
      msg.Body = JSON.stringify(content);
      console.log(msg)
      return this.wsConn.ask(msg).catch(err => console.error(err));
    }
    return null;
  }
}