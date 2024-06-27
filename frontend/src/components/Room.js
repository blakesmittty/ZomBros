import React, { useState, useEffect } from "react";
import { useParams } from 'react-router-dom';

function Room() {
    const { roomID } = useParams();
    const [players, setPlayers] = useState([]);
    const playerId = localStorage.getItem('playerId');

    useEffect(() => {
        const ws = new WebSocket(`ws://localhost:8080/ws?playerId=${playerId}&roomID=${roomID}`);

        ws.onopen = () => {
            console.log('websocket connection opened');
            if (roomID) {
                ws.send(`join:${roomID}`);
            } else {
                ws.send('create');
            }
        };

        ws.onmessage = (e) => {
            try {
                const data = JSON.parse(e.data);
                console.log('Received player list: ', data);
                setPlayers(data);
            } catch (error) {
                console.error('Received non JSON message: ', e.data);
            }
        };

        ws.onclose = (e) => {
            console.error('websocket closed unexpectedly:', e);
        };

        ws.onerror = (error) => {
            console.error('websocket error: ', error);
        }

        return () => {
            if (ws.readyState === 1) {
                ws.close();
            }
        };
    }, [playerId, roomID]);

    return (
        <div>
            <h2>Game Room</h2>
            <h2>your playerId: {playerId}</h2>
            <h3>Players:</h3>
            <ul>
                {players.map((player) => (
                    <li key={player.id}>{player.username}</li>
                ))}
            </ul>
        </div>
    );

}
export default Room;