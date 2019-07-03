import React from 'react';
import './App.scss';
import * as layout from './app/layout';

const App: React.FC = () => {
  return (
    <div className="container">
      <div>
        <layout.Header />
      </div>
      <div>
        <layout.Main />
      </div>
    </div>
  );
}

export default App;
