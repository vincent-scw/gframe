import React from  'react';
import playerReceptionService from '../services/player_reception.service';

export class Main extends React.Component {
  constructor(props: any) {
    super(props);

    this.onStart = this.onStart.bind(this);
  }

  render() {
    return (
      <div>
        <button className="button is-primary" onClick={this.onStart}>Start</button>
      </div>
    );
  }

  onStart() {
    playerReceptionService.in()
      .then(_ => {
        
      })
      .catch(err => {
        console.error(err);
      });
  }
}