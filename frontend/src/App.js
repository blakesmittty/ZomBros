import { useState, useEffect } from 'react';
import { BrowserRouter as Router, Route, Routes, Navigate } from 'react-router-dom';
import './App.css';
import Game from './components/Game';
import Register from './components/Register';
import Login from './components/Login';
import Items from './components/Items';
import Room from './components/Room';
import Lobby from './components/Lobby';
import Room2 from './components/Room2';

function App() {
  const [loggedIn, setLoggedIn] = useState(localStorage.getItem('loggedIn') === 'true');
  const [playerId, setPlayerId] = useState(localStorage.getItem('playerId'));

  const handleLogin = (id) => {
    setPlayerId(id);
    setLoggedIn(true);
  };

  useEffect(() => {
    if (loggedIn && playerId) {
      localStorage.setItem('loggedIn', 'true');
      localStorage.setItem('playerId', playerId);
    } else {
      localStorage.removeItem('loggedIn');
      localStorage.removeItem('playerId');
    }
  }, [loggedIn, playerId])

  return (
    <Router>
      <div className='App'>
        <h1>Game</h1>
        <Routes>
          <Route path='/register' element={<Register />} />
          <Route path='/login' element={<Login setLoggedIn={handleLogin} />} />
          <Route path='/items' element={loggedIn ? <Items playerId={playerId} /> : <Navigate to='/login' />} />
          {/*<Route path='/game' element={loggedIn ? <Game /> : <Navigate to='/login' />} />*/}
          {/*<Route path='/' element={loggedIn ? <Navigate to='/game' /> : <Navigate to='/login' />} />*/}
          <Route path='/' element={loggedIn ? <Lobby /> : <Navigate to='/login' />} />
          <Route path='/room/:roomID' element={<Room2 />} />
        </Routes>
      </div>
    </Router>
  );
}

export default App;
