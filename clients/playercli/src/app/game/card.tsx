import React from 'react';

export class Card extends React.Component<any, any> {
  constructor(props: any) {
    super(props);
  }

  render() {
    return (
      <div className="card">
        <header className="card-header">
          <p className="card-header-title">
            {this.props.player}
          </p>
        </header>
        <div className="card-image">
          <figure className="image is-4by3">
            <img src="https://bulma.io/images/placeholders/1280x960.png" alt="Placeholder image" />
          </figure>
        </div>
          <div className="card-footer">
            <a className="card-footer-item">Rock</a>
            <a className="card-footer-item">Paper</a>
            <a className="card-footer-item">Scissors</a>
          </div>
        </div>
        );
      }
}