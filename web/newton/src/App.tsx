import React from 'react';
import './App.css';
import {LightTheme, ThemeProvider} from 'baseui';

function App() {
  return (
    <ThemeProvider theme={LightTheme}>
      I can use themed Base Web components here!
    </ThemeProvider>
  );
}

export default App;
