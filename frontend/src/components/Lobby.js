import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';

function Lobby() {
    const [roomID, setRoomID] = useState('');
    const [username, setUsername] = useState('');
    const navigate = useNavigate();
    const playerID = localStorage.getItem('playerId');

    const handleCreateRoom = async () => {
        try {
            const response = await fetch('http://localhost:8080/createRoom', {
                method: 'POST',
            });
            const newRoomID = await response.text();
            navigate(`/room/${newRoomID}?username=${username}`);
        } catch (error) {
            console.error('error creating room:', error);
        }
    };

    const handleJoinRoom = () => {
        if (!roomID) {
            alert('Please enter a room ID');
            return;
        }
        navigate(`/room/${roomID}?username=${username}`);
    };

    return (
        <div>
            <h2>Lobby</h2>
            <h3>playerID : {playerID}</h3>
            <input
                type='text'
                placeholder='enter username'
                value={username}
                onChange={(e) => setUsername(e.target.value)}
                required
            />
            <button onClick={handleCreateRoom}>create room</button>
            <input
                type='text'
                placeholder='enter room ID'
                value={roomID}
                onChange={(e) => setRoomID(e.target.value)}
            />
            <button onClick={handleJoinRoom}>Join Room</button>
        </div>
    );
}

export default Lobby;