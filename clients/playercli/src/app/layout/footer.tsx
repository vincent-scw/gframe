import React from 'react';

export class Footer extends React.Component {
  render() {
    return (
      <div className="content has-text-centered">
        <p>
          Follow <strong>gframe</strong> at&nbsp;
            <button className="button is-small" onClick={() => window.open("https://github.com/vincent-scw/gframe")}>
            <span className="icon">
              <i className="fab fa-github"></i>
            </span>
            <span>GitHub</span>
          </button>.
          </p>
      </div>
    );
  }
}