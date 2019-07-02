import React from 'react';
import { Register, Welcome } from '../account';
import authService from '../services/auth.service';
import User from '../account/user.model';

export class Header extends React.Component<any, User> {
  constructor(props: any) {
    super(props);
    this.state = { username: '' };
    this.userLogin = this.userLogin.bind(this);
  }

  componentDidMount() {
    authService.userSubject.subscribe(x => {
      if (x != null)
        this.setState(x)
    })
  }

  render() {
    const isLoggedIn = this.state.username !== '';
    return (
      <div>
        {isLoggedIn ? <Welcome username={this.state.username} />
          : <Register onUserLogin={this.userLogin} />
        }
      </div>);
  }

  userLogin(user: any) {

  }

  setUser(user: any) {
    localStorage.setItem("user", JSON.stringify(user));
  }
}