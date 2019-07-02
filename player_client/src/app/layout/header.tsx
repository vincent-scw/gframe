import React from 'react';
import { Register, Welcome } from '../account';
import authService from '../services/auth.service';
import User from '../account/user.model';

export class Header extends React.Component<any, User> {
  constructor(props: any) {
    super(props);
    this.state = { username: '' };
  }

  componentDidMount() {
    authService.userSubject.subscribe(x => {
      if (x != null)
        this.setState(x)
      else
        this.setState({ username: '' })
    })
  }

  render() {
    const isLoggedIn = this.state.username !== '';
    return (
      <nav className="navbar" role="navigation" aria-label="main navigation">
        <div className="navbar-brand"></div>
        <div className="navbar-menu">
          <div className="navbar-end">
            {isLoggedIn ? <Welcome username={this.state.username} />
              : <Register />
            }
          </div>
        </div>
      </nav>);
  }
}