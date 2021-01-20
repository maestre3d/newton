import './App.css'
import { BrowserRouter as Router } from 'react-router-dom'
import DefaultShell from './app/shared/component/default-shell/Default-Shell'
import { BaseProvider, LightTheme } from 'baseui'
import { Provider as StyletronProvider, DebugEngine } from 'styletron-react'
import { Client as Styletron } from 'styletron-engine-atomic'

const debug = process.env.NODE_ENV === 'production' ? void 0 : new DebugEngine()
const engine = new Styletron()

// App root component
function App() {
  return (
    <StyletronProvider value={engine} debug={debug} debugAfterHydration>
      <BaseProvider theme={LightTheme}>
        <Router>
          <DefaultShell />
        </Router>
      </BaseProvider>
    </StyletronProvider>
  )
}

export default App
