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
import GameRoom from './components/GameRoom';
import Game2 from './components/Game2';

function App() {
  const [loggedIn, setLoggedIn] = useState(localStorage.getItem('loggedIn') === 'true');
  const [playerId, setPlayerId] = useState(localStorage.getItem('playerId'));
  const [username, setUsername] = useState(localStorage.getItem('username'));

  const handleLogin = (id, username) => {
    console.log('handle login called with: ', id, username);
    setPlayerId(id);
    setUsername(username);
    setLoggedIn(true);
  };

  useEffect(() => {
    if (loggedIn && playerId && username) {
      localStorage.setItem('loggedIn', 'true');
      localStorage.setItem('playerId', playerId);
      localStorage.setItem('username', username);
    } else {
      localStorage.removeItem('loggedIn');
      localStorage.removeItem('playerId');
      localStorage.removeItem('username');
    }
  }, [loggedIn, playerId, username])

  return (
    <Router>
      <div className='App'>
        <h1>Game</h1>
        <Routes>
          <Route path='/register' element={<Register />} />
          <Route path='/login' element={<Login handleLogin={handleLogin} />} />
          <Route path='/items' element={loggedIn ? <Items playerId={playerId} /> : <Navigate to='/login' />} />
          {/*<Route path='/game' element={loggedIn ? <Game /> : <Navigate to='/login' />} />*/}
          {/*<Route path='/' element={loggedIn ? <Navigate to='/game' /> : <Navigate to='/login' />} />*/}
          <Route path='/' element={loggedIn ? <GameRoom /> : <Navigate to='/login' />} />
          <Route path='/gameRoom' element={<GameRoom />} />
          <Route path='/room/:roomID' element={<Game2 />} />
        </Routes>
      </div>
    </Router>
  );
}

export default App;
