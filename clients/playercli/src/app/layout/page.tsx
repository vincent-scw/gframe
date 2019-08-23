import React from 'react';
import { Header } from './header';
import { Main } from './main';
import { Register } from '../account/index';
import authService from '../services/auth.service';
import './page.scss';

interface PageState {
  content: any;
}

export class Page extends React.Component<any, PageState> {
  constructor(prop: any) {
    super(prop)
    this.state = { content: null };
  }

  componentDidMount() {
    authService.userSubject.subscribe(x => {
      this.setState({
        content:
          authService.isLoggedin ?
            <div className="container">
              <div><Header /></div>
              <div><Main /></div>
            </div>
            :
            <div className="full-page-size">
              <div className="columns is-centered is-vcentered">
                <div className="column register">
                  <Register />
                </div>
              </div>
            </div>
      })
    });
  }

  render() {
    return this.state.content;
  }
}