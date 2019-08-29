import React from 'react';
import { Welcome } from '../account';
import { authService } from '../services';
import User from '../account/user.model';

export class Header extends React.Component<any, User> {
  constructor(props: any) {
    super(props);
    this.state = { username: authService.user.username };
  }

  render() {
    return (
      <div className="section">
        <Welcome username={this.state.username} />
      </div>
    );
  }
}