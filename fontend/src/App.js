import React from 'react';
import {
  BrowserRouter as Router,
  Switch,
  Route,
} from "react-router-dom";
import 'react-notifications/lib/notifications.css';
import {NotificationContainer} from 'react-notifications';

import Login from './pages/Login'
import Register from './pages/Register'
import Shift from './pages/Shift'
import './App.css';

function App() {
  return (
    <Router>
      <Switch>
        <Route path="/login">
          <Login/>
        </Route>
        <Route path="/register">
          <Register/>
        </Route>
        <Route path="/">
          <Shift/>
        </Route>
      </Switch>
      <NotificationContainer/>
    </Router>
  );
}

export default App;
