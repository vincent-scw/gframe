import React from 'react';

const smallIconStyle = {
  height: "50px"
}

const largeIconStyle = {
  width: "28em",
  height: "28em",
  paddingLeft: "52px"
}

enum Shape {
  NotSet = 0,
  Rock = 1,
  Paper = 2,
  Scissors = 3
}

interface CardState {
  selectedShape: Shape;
}

export class Card extends React.Component<any, CardState> {
  constructor(props: any) {
    super(props);

    this.state = { selectedShape: Shape.NotSet }

    this.onSelected = this.onSelected.bind(this);
  }

  onSelected(shape: Shape) {
    this.setState({ selectedShape: shape });
  }

  render() {
    return (
      <div className="card">
        <header className="card-header">
          <p className="card-header-title">
            <i>Player:</i>&nbsp;{this.props.player}
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
          <a className="card-footer-item" onClick={() => this.onSelected(Shape.Rock)}>
            <span className="icon" style={smallIconStyle}>
              <i className="far fa-3x fa-hand-rock"></i>
            </span>
          </a>
          <a className="card-footer-item" onClick={() => this.onSelected(Shape.Paper)}>
            <span className="icon" style={smallIconStyle}>
              <i className="far fa-3x fa-hand-paper"></i>
            </span>
          </a>
          <a className="card-footer-item" onClick={() => this.onSelected(Shape.Scissors)}>
            <span className="icon" style={smallIconStyle}>
              <i className="far fa-3x fa-hand-scissors"></i>
            </span>
          </a>
        </div>
      </div>
    );
  }
}