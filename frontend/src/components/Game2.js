/* eslint-disable no-undef */
import React, { useEffect, useState, useRef } from 'react';
import { useParams } from 'react-router-dom';
import CharacterSelect from './CharacterSelect';
//import CharacterSprite from './CharacterSprite';
import GameCanvas from './GameCanvas';
import * as _ from '../utils/gamestate_pb'; // crazy strange error, module doesnt work if imported as proto

const Game2 = () => {
    const [socket, setSocket] = useState(null);
    const [playerList, setPlayerList] = useState([]);
    const playerID = localStorage.getItem('playerId');
    const { roomID } = useParams();
    const [selectedCharacter, setSelectedCharacter] = useState(null);
    const [isReady, setIsReady] = useState(false);
    const [allPlayersReady, setAllPlayersReady] = useState(false);
    const [movement, setMovement] = useState({ x: 0, y: 0 });
    const [gameState, setGameState] = useState(null);
    const [initGameState, setInitGameState] = useState(null);
    const [isShooting, setIsShooting] = useState(false);
    const [aimAngle, setAimAngle] = useState(0);
    const username = new URLSearchParams(window.location.search).get('username');
    const canvasRef = useRef(null);

    useEffect(() => {

        const ws = new WebSocket(`ws://localhost:8080/ws?playerID=${playerID}&roomID=${roomID}&username=${encodeURIComponent(username)}`);

        ws.onopen = () => {
            console.log('connected to websocket');
        };

        ws.onmessage = async (e) => {
            if (typeof e.data === 'string') {
                try {
                    const data = JSON.parse(e.data);
                    if (data.type === 'playerList') {
                        setPlayerList(data.playerList);
                        checkAllPlayersReady(data.playerList);
                    } else {
                        console.error('unknown json message type:', data.type);
                    }
                } catch (error) {
                    console.error('Error parsing JSON message: ', error);
                }
            } else {
                try {
                    const arrayBuffer = await e.data.arrayBuffer();
                    //console.log("arrayBuffer: ", arrayBuffer);
                    const array = new Uint8Array(arrayBuffer);
                    //console.log("array", array);
                    const protoGameState = proto.gamestate.GameState.deserializeBinary(array);
                    setGameState(protoGameState);
                } catch (error) {
                    console.error('error decoding gamestate: ', error)
                }
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

    const checkAllPlayersReady = (players) => {
        const allReady = players.every(player => player.isReady);
        setAllPlayersReady(allReady);
    };

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
    if (allPlayersReady) {
        return (
            <GameCanvas playerID={playerID} initialGameState={gameState} socket={socket} />
        );
    }

    return (
        <div>
            <h2>Game Room: {roomID}</h2>
            <h3>Players in this room:</h3>
            <ul>
                {playerList.map((player, index) => (
                    <li key={index}>
                        {player.username} - {player.character} {player.isReady ? '(Ready)' : '(Not Ready)'}
                        {/*<CharacterSprite character={player.character} />*/}
                    </li>
                ))}
            </ul>
            <CharacterSelect onSelectCharacter={handleSelectCharacter} />
            <button onClick={handleReadyUp} disabled={!selectedCharacter}>Ready Up</button>
            {/* Rest of your game UI */}
        </div>
    );
};

export default Game2;