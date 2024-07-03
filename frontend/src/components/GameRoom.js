import React, { useState } from "react";
import axios from 'axios';
import { useNavigate, useParams } from "react-router-dom";

const GameRoom = () => {
    const [roomID, setRoomID] = useState('');
    const [error, setError] = useState('');
    const navigate = useNavigate();
    const playerID = localStorage.getItem('playerId');
    //const { username } = useParams();
    const username = localStorage.getItem('username');

    const createRoom = async () => {
        try {
            const response = await axios.post('http://localhost:8080/createRoom');
            setRoomID(response.data.roomID);
        } catch (error) {
            setError('Failed to create room');
        }
    };

    const joinRoom = async () => {
        try {
            const response = await axios.get(`http://localhost:8080/joinRoom?roomID=${roomID}`);
            if (response.data.success) {
                navigate(`/room/${roomID}?username=${username}`);
            } else {
                setError('Room not found');
            }
        } catch (error) {
            setError('failed to join room');
        }
    }

    return (
        <div>
            <h1>Zombie Game Room</h1>
            <h2>playerID: {playerID} </h2>
            <h2>username: {username} </h2>
            <button onClick={createRoom}>Create New Room</button>
            <div>
                <input
                    type="text"
                    value={roomID}
                    onChange={(e) => setRoomID(e.target.value)}
                    placeholder="Enter Room ID"
                />
                <button onClick={joinRoom}>Join Room</button>
            </div>
            {error && <p style={{ color: 'red' }}>{error}</p>}
        </div>
    );
};

export default GameRoom;