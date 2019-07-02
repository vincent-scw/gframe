import React from 'react';
import User from './user.model';

export class Welcome extends React.Component<User, any> {
  constructor(props: User) {
    super(props);
  }

  render() {
    return (
      <div>
        Hi, {this.props.username}!
      </div>
    );
  }
}