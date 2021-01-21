import React from 'react'
import './App.css'
import { BrowserRouter as Router } from 'react-router-dom'
import DefaultShell from './app/shared/component/default-shell/Default-Shell'
import { BaseProvider, DarkTheme, LightTheme } from 'baseui'
import { Provider as StyletronProvider, DebugEngine } from 'styletron-react'
import { Client as Styletron } from 'styletron-engine-atomic'
import { THEME, ThemeContext, useTheme } from './internal/shared/infrastructure/theme'


const debug = process.env.NODE_ENV === 'production' ? void 0 : new DebugEngine()
const engine = new Styletron()

// App root component
function App() {
  const theme = useTheme()
  return (
    <StyletronProvider value={engine} debug={debug} debugAfterHydration>
      <ThemeContext.Provider value={theme}>
        <BaseProvider theme={theme.theme === THEME.light ? LightTheme : DarkTheme}>
          <Router>
            <DefaultShell />
          </Router>
        </BaseProvider>
      </ThemeContext.Provider>
    </StyletronProvider>
  )
}

export default App
