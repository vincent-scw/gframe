import React from 'react';
import User from './user.model';
import { authService } from '../services';
import './welcome.scss';

export class Welcome extends React.Component<User, any> {
  render() {
    return (
      <div className="title">
        <span>Hello, {this.props.username}!</span>
        <a onClick={authService.logout} className="button logout" title="Logout">
          <span className="icon has-text-grey-light">
            <i className="fas fa-lg fa-sign-out-alt"></i>
          </span>
        </a>
      </div>
    );
  }
}