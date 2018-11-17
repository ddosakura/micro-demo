import React, { Component, useState } from 'react';
import logo from './logo.svg';
import './App.css';

const { UserServiceClient } = require('./proto/user/user_grpc_web_pb');
const { UserRequest } = require('./proto/user/user_pb.js');

function initClient() {

  const client = new UserServiceClient('localhost:50001');

  const request = new UserRequest();
  request.setMessage('Hello World!');

  const metadata = { 'custom-header-1': 'value1' };

  client.echo(request, metadata, (err, response) => {
    // ...
  });
}

function App() {
  const [count, setCount] = useState(0);

  initClient()

  return (
    <div className="App">
      <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
        <p>
          Edit <code>src/App.js</code> and save to reload.
          </p>
        <a
          className="App-link"
          href="https://reactjs.org"
          target="_blank"
          rel="noopener noreferrer"
        >
          Learn React
          </a>
        <p>You clicked {count} times</p>
        <button onClick={() => setCount(count + 1)}>
          Click me
        </button>
      </header>
    </div>
  );
}

export default App;
