import React, { useState, useEffect } from 'react';
import { useParams, useLocation } from 'react-router-dom';

function Room2() {
    const { roomID } = useParams();
    const [players, setPlayers] = useState([]);
    const location = useLocation();
    const queryParams = new URLSearchParams(location.search);
    const username = queryParams.get('username');
    const playerID = localStorage.getItem('playerId');

    useEffect(() => {
        if (!username) {
            console.error('No username provided.');
            return;
        }

        const ws = new WebSocket(`ws://localhost:8080/ws?playerId=${new Date().getTime()}&roomID=${roomID}&username=${username}`);

        ws.onopen = () => {
            console.log('WebSocket connection opened');
        };

        ws.onmessage = (event) => {
            const playerList = JSON.parse(event.data);
            console.log('Received player list:', playerList);
            setPlayers(playerList);
        };

        ws.onclose = (event) => {
            console.error('WebSocket closed unexpectedly:', event);
        };

        ws.onerror = (error) => {
            console.error('WebSocket error:', error);
        };

        return () => {
            if (ws.readyState === 1) {
                ws.close();
            }
        };
    }, [username, roomID]);

    return (
        <div>
            <h2>Game Room: {roomID}</h2>
            <h3>playerID: {playerID}</h3>
            <h3>Players:</h3>
            <ul>
                {players.map((player) => (
                    <li key={player.ID}>{player.Username}</li>
                ))}
            </ul>
        </div>
    );
}

export default Room2;
