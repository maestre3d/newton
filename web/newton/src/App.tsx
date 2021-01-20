import './App.css'
import { BrowserRouter as Router } from 'react-router-dom'
import DefaultShell from './app/shared/component/default-shell/Default-Shell'

// App root component
function App() {
  return (
    <Router>
      <DefaultShell />
    </Router>
  );
}

export default App
