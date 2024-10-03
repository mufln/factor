import React from 'react';
import { Routes, Route } from 'react-router-dom';
import './App.css';
import LoginRegister from './pages/login/login';
import Base from './pages/main/base';

function App() {
  return (
    <div className="App">
      <Routes>
        {/* подстановочный путь */}
        <Route path="/" element={<LoginRegister />} />
        <Route path="/main" element={<Base/>} />
      </Routes>
    </div>
  );
}

export default App;
