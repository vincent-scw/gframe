import React from 'react';
import { gameService, EVENT_GAME } from '../services';
import { Shape, PlayEvent } from './event.model';
import { Player, GroupEvent } from '../services/server-events.model';

const smallIconStyle = {
  height: "50px"
}

const largeIconStyle = {
  width: "28em",
  height: "28em",
  paddingLeft: "52px"
}

interface CardProps {
  readonly: boolean;
  player: Player;
  group: GroupEvent | null;
}

interface CardState {
  selectedShape: Shape;
}

export class Card extends React.Component<CardProps, CardState> {
  constructor(props: CardProps) {
    super(props);

    this.state = { selectedShape: Shape.NotSet }

    this.onSelected = this.onSelected.bind(this);
  }

  onSelected(shape: Shape) {
    if (this.props.readonly) return;
    this.setState({ selectedShape: shape });
    gameService.ask(EVENT_GAME, {
      play: {
        player: this.props.player, 
        shape: shape
      },
      group: this.props.group
    })
  }

  render() {
    return (
      <div className="card">
        <header className="card-header">
          <p className="card-header-title">
            <i>Player:</i>&nbsp;{this.props.player.name}
          </p>
        </header>
        <div className="card-image">
          {this.state.selectedShape === Shape.NotSet && 
            <figure><i className="fas fa-question" style={largeIconStyle}></i></figure>}
          {this.state.selectedShape === Shape.Rock && 
            <figure><i className="far fa-hand-rock" style={largeIconStyle}></i></figure>}
          {this.state.selectedShape === Shape.Paper && 
            <figure><i className="far fa-hand-paper" style={largeIconStyle}></i></figure>}
          {this.state.selectedShape === Shape.Scissors && 
            <figure><i className="far fa-hand-scissors" style={largeIconStyle}></i></figure>}
        </div>
        <div className="card-footer">
          <a className="card-footer-item"
            onClick={() => this.onSelected(Shape.Rock)}>
            <span className="icon" style={smallIconStyle}>
              <i className="far fa-3x fa-hand-rock"></i>
            </span>
          </a>
          <a className="card-footer-item"
            onClick={() => this.onSelected(Shape.Paper)}>
            <span className="icon" style={smallIconStyle}>
              <i className="far fa-3x fa-hand-paper"></i>
            </span>
          </a>
          <a className="card-footer-item"
            onClick={() => this.onSelected(Shape.Scissors)}>
            <span className="icon" style={smallIconStyle}>
              <i className="far fa-3x fa-hand-scissors"></i>
            </span>
          </a>
        </div>
      </div>
    );
  }
}