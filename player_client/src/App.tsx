import React from 'react';
import logo from './logo.svg';
import './App.scss';
import * as account from './app/account';

const App: React.FC = () => {
  return (
    <div>
      <account.Register />
    </div>
  );
}

export default App;
