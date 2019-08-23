import React from 'react';
import './game-pad.scss';

export class GamePad extends React.Component {
  constructor(prop: any) {
    super(prop);

    this.onStart = this.onStart.bind(this);
  }

  onStart() {

  }

  render() {
    return (
      <div className="content has-text-centered start-game-div">
        <button className="button is-primary is-large" onClick={this.onStart}>Start Game</button>
      </div>
    );
  }
}