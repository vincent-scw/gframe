import { Component, OnInit, Input, ViewEncapsulation, OnDestroy } from '@angular/core';
import { SubscriptionClient } from 'subscriptions-transport-ws';
import { WebSocketLink } from 'apollo-link-ws';
import { environment } from '../../../environments/environment';
import ApolloClient from 'apollo-client';
import { InMemoryCache } from 'apollo-cache-inmemory';
import gql from 'graphql-tag';

const GRAPHQL_ENDPOINT = `${environment.wsProtocol}://${environment.services.admin}/console`;

@Component({
  selector: 'app-console',
  templateUrl: './console.component.html',
  styleUrls: ['./console.component.scss'],
  encapsulation: ViewEncapsulation.None,
})
export class ConsoleComponent implements OnInit, OnDestroy {
  client: SubscriptionClient;
  messages: Array<any>;

  constructor() {
    this.messages = [];
  }

  ngOnInit() {
    try {
      this.client = new SubscriptionClient(GRAPHQL_ENDPOINT, {
        reconnect: true,
      });
      const link = new WebSocketLink(this.client)
      const cache = new InMemoryCache();
      const apolloClient = new ApolloClient({ cache, link });
      apolloClient.subscribe({
        query: gql`
          subscription onNewMsg {
              newMsgCreated {
                  id
              }
          }`,
        variables: {}
      }).subscribe(
        (res)=> {
          this.messages.push(res.data)
        }
      );

      this.messages = ["Start listening..."];
    } catch (err) {
      console.error(err);
      this.messages = [
        `<i class='red'>Error</i>: ${err}`
      ];
    }
  }

  ngOnDestroy() {
    if (this.client != null) {
      this.client.close();
    }
  }
}
