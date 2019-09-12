import React from 'react';
import { gameService, EVENT_GAME } from '../services';
import { Shape, Move, Player, GroupEvent } from '../services/events.model';

interface CardProps {
  readonly: boolean;
  player: Player;
  shape: Shape;
  group: GroupEvent | null;
}

interface CardState {
  selectedShape: Shape;
  readonly: boolean;
}

export class Card extends React.Component<CardProps, CardState> {
  constructor(props: CardProps) {
    super(props);

    this.state = { selectedShape: props.shape, readonly: props.readonly }

    this.onSelected = this.onSelected.bind(this);
  }

  onSelected(shape: Shape) {
    this.setState({ selectedShape: shape });
    let ret = gameService.ask(
      EVENT_GAME,
      {
        move: {
          player: this.props.player,
          shape: shape
        },
        group: this.props.group
      }
    );
    if (ret != null) {
      ret.then(() => this.setState({ readonly: true}));
    }
  }

  render() {
    return (
      <div>
        <header>
          <p>
            <i>Player:</i>&nbsp;{this.props.player.name}
          </p>
        </header>
        <div className="field has-addons">
          <p className="control">
            <button 
              className={`button is-large ${this.state.selectedShape===Shape.Rock ? 'is-info is-selected' : ''}`} 
              disabled={this.state.readonly}
              onClick={() => this.onSelected(Shape.Rock)}>
              <span className="icon">
                <i className="far fa-hand-rock"></i>
              </span>
              <span>Rock</span>
            </button>
          </p>
          <p className="control">
            <button 
              className={`button is-large ${this.state.selectedShape===Shape.Paper ? 'is-info is-selected' : ''}`} 
              disabled={this.state.readonly}
              onClick={() => this.onSelected(Shape.Paper)}>
              <span className="icon">
                <i className="far fa-hand-paper"></i>
              </span>
              <span>Paper</span>
            </button>
          </p>
          <p className="control">
            <button 
              className={`button is-large ${this.state.selectedShape===Shape.Scissors ? 'is-info is-selected' : ''}`} 
              disabled={this.state.readonly}
              onClick={() => this.onSelected(Shape.Scissors)}>
              <span className="icon">
                <i className="far fa-hand-scissors"></i>
              </span>
              <span>Scissors</span>
            </button>
          </p>
        </div>
      </div>
    );
  }
}