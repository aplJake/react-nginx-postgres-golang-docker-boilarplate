import React from 'react';
import logo from './logo.svg';
import './App.css';
import Axios from 'axios';

const BACKEND_URL = 'http://localhost:5000';

class App extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      msg: "",
    }
  }

  componentWillMount() {
    Axios
      .get("http://localhost:5000/api/msg")
      .then(res => {
        const message = res.data;
        this.setState({
          msg: message
        }, console.log(this.state.msg));
      })
      .catch(err => console.log(err))
  }

  render() {
    return(
      <div className="App">
        <header className="App-header">
          <img src={logo} className="App-logo" alt="logo" />
          <p>
            Edit <code>src/App.js</code> and save to reload.
          </p>
          <h4>
            {this.state.msg}
          </h4>
        </header>
      </div>
    );
  }
}

export default App;
