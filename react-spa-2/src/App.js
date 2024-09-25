import React from 'react';
import { Routes, Route } from 'react-router-dom';
import './App.css';
import Login from './components/Login/Login'
import Register from './pages/register';

function App() {
  return (
    <div className="App">
      <Routes>
        {/* подстановочный путь */}
        <Route path="/" element={<Login />} />
        <Route path="register" element={<Register />} />
      </Routes>
    </div>
  );
}

export default App;
