import { Component, OnInit, Input, ViewEncapsulation } from '@angular/core';
import * as neffos from 'neffos.js';
import { environment } from '../../../environments/environment';

@Component({
  selector: 'app-console',
  templateUrl: './console.component.html',
  styleUrls: ['./console.component.scss'],
  encapsulation: ViewEncapsulation.None,
})
export class ConsoleComponent implements OnInit {
  @Input()
  set listening(l: boolean) {
    if (l) {
      this.startListener();
    } else {
      if (this.conn != null && !this.conn.isClosed) {
        this.conn.close();
      }
      this.messages = [
        "Console is sleeping..."
      ];
    }
  }

  conn: neffos.Conn;
  messages: Array<any>;

  constructor() {
  }

  ngOnInit() {
  }

  async startListener() {
    try {
      this.conn = await neffos.dial(`${environment.wsProtocol}://${environment.services.admin}/console`, {
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
            this.messages.push(msg.Body)
          }
        }
      }, { // optional.
          reconnect: 2000,
          // set custom headers.
          headers: {
            // 'X-Username': 'kataras',
          }
        });

      await this.conn.connect("default");
      this.messages = ["Start listening..."];
    } catch (err) {
      console.error(err);
      this.messages = [
        `<i class='red'>Error</i>: ${err}`
      ];
    }
  }
}
