import React from 'react';
import './App.scss';
import * as layout from './app/layout';

const App: React.FC = () => {
  return (
    <div className="container">
      <div>
        <layout.Header />
      </div>
    </div>
  );
}

export default App;
