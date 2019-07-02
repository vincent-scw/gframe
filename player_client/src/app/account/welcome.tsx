import React from 'react';
import User from './user.model';
import authService from '../services/auth.service';

export class Welcome extends React.Component<User, any> {
  render() {
    return (
      <div className="navbar-item">
        <div className="navbar-item is-info">Hi, {this.props.username}!</div>
        <button className="button is-light" onClick={authService.logout}>
          <strong>Log out</strong>
        </button>
      </div>
    );
  }
}