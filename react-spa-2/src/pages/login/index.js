import React from 'react';
import '../../App.css';
import Login from '../../components/Login/Login'
import Register from '../register';

function App() {
  return (
    <div className="App">
      <Switch>
        <Route path='/register' component={Register} />
        <Route exact path='/' component={Login} />
        <Redirect to='/' />
      </Switch>
    </div>
  );
}

export default App;
