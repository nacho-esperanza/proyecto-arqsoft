import logo from './logo.svg';
import './App.css';

function App() {
  return (
    <div className="App">
      <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
        <p>
          Edit <code>src/App.js</code> and save to reload.
        </p>
        <a>
          Hello
        </a>
        <a>
          contraseña 
        <input
            variant="outlined"
            margin="normal"
            required
            fullWidth
            name="password"
            label="Password"
            type="password"
            id="password"
          >
          
          </input>
        <button
        type="button"
        fullWidth
        variant="contained"
        color="primary"
        >login
        </button>

        </a>
      </header>
    </div>
  );
}

export default App;
