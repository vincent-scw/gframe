import React from 'react';
import { GamePad } from '../game';

interface MainState {

}

const divStyle = {
  minHeight: '60vh'
};

export class Main extends React.Component<any, MainState> {
  constructor(props: any) {
    super(props);

    this.state = {}
  }

  render() {
    return (
      <div style={divStyle}>
        <GamePad />
      </div>
    );
  }
}