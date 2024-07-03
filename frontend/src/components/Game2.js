import React, { useEffect, useState } from 'react';
import { useParams } from 'react-router-dom';
import CharacterSelect from './CharacterSelect';

const Game2 = () => {
    const [socket, setSocket] = useState(null);
    const [playerList, setPlayerList] = useState([]);
    const playerID = localStorage.getItem('playerId');
    const { roomID } = useParams();
    const [selectedCharacter, setSelectedCharacter] = useState(null);
    const [isReady, setIsReady] = useState(false);

    const username = new URLSearchParams(window.location.search).get('username');

    useEffect(() => {
        const ws = new WebSocket(`ws://localhost:8080/ws?playerID=${playerID}&roomID=${roomID}&username=${encodeURIComponent(username)}`);

        ws.onopen = () => {
            console.log('connected to websocket');
        };

        ws.onmessage = (e) => {
            console.log('Received: ', e.data);
            try {
                const data = JSON.parse(e.data)
                if (data.type === 'playerList') {
                    setPlayerList(data.playerList);
                }
            } catch (error) {
                console.error('Error parsing websocket message:', error);
            }
        };

        ws.onclose = () => {
            console.log('disconnected from websocket');
        };

        setSocket(ws);

        return () => {
            if (ws.readyState === 1) {
                ws.close();
            };
        }
    }, [roomID]);

    const handleSelectCharacter = (character) => {
        setSelectedCharacter(character);
    };

    const handleReadyUp = () => {
        setIsReady(true);
        if (selectedCharacter && socket && socket.readyState === WebSocket.OPEN) {
            socket.send(JSON.stringify({
                type: 'selectCharacter',
                playerID,
                character: selectedCharacter,
                isReady: true,
            }));
        }
    };

    return (
        <div>
            <h2>Game Room: {roomID}</h2>
            <h3>Players in this room:</h3>
            <ul>
                {playerList.map((player, index) => (
                    <li key={index}>
                        {player.username} - {player.character} {player.isReady ? '(Ready)' : '(Not Ready)'}
                    </li>
                ))}
            </ul>
            <CharacterSelect onSelectCharacter={handleSelectCharacter} />
            <button onClick={handleReadyUp}>Ready Up</button>
            {/* Rest of your game UI */}
        </div>
    );
};

export default Game2;