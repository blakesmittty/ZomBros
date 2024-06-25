import { useState } from 'react';
import { BrowserRouter as Router, Route, Routes, Navigate } from 'react-router-dom';
import './App.css';
import Game from './components/Game';
import Register from './components/Register';
import Login from './components/Login';
import Items from './components/Items';

function App() {
  const [loggedIn, setLoggedIn] = useState(false);
  const [playerId, setPlayerId] = useState(null);

  const handleLogin = (id) => {
    setPlayerId(id);
    setLoggedIn(true);
  };

  return (
    <Router>
      <div className='App'>
        <h1>Game</h1>
        <Routes>
          <Route path='/register' element={<Register />} />
          <Route path='/login' element={<Login setLoggedIn={handleLogin} />} />
          <Route path='/items' element={loggedIn ? <Items playerId={playerId} /> : <Navigate to='/login' />} />
          <Route path='/' element={loggedIn ? <Navigate to='/items' /> : <Navigate to='/login' />} />
        </Routes>
      </div>
    </Router>
  );
}

export default App;
