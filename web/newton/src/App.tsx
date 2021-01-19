import React from 'react';
import './App.css';
import {LightTheme, ThemeProvider} from 'baseui';
import {Foo} from './Foo';
import Home from './app/page/home/Home';
import DefaultShell from './app/shared/component/default-shell/Default-Shell';

function App() {
  return (
    <ThemeProvider theme={LightTheme}>
      <div>
        <DefaultShell />
      </div>
      <div>
        <Foo />
      </div>
      <div>
        <Home />
      </div>
    </ThemeProvider>
  );
}

export default App;
