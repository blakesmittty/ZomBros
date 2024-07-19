/* eslint-disable no-undef */
import React, { useEffect, useState, useRef } from 'react';
import { useParams } from 'react-router-dom';
import CharacterSelect from './CharacterSelect';
import CharacterSprite from './CharacterSprite';
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

    const mapWidth = 3000;
    const mapHeight = 3000;
    const viewportWidth = 800;
    const viewportHeight = 600;

    /*
    const playerRef = useRef({
        x: 1500,
        y: 1500,
        moveX: 0,
        moveY: 0,
        width: 20,
        height: 20,
        speed: 5,
        isShooting: false,
        aimAngle: 0
    });
    */
    /*
    const map = new Image();
    map.src = '/testmap.png';

    useEffect(() => {
        const handleMapLoad = () => {
            console.log('map image loaded');
            const backgroundCanvas = backgroundCanvasRef.current;
            if (backgroundCanvas) {
                const backgroundCtx = backgroundCanvas.getContext('2d');
                backgroundCanvas.width = mapWidth;
                backgroundCanvas.height = mapHeight;
                backgroundCtx.drawImage(map, 0, 0, mapWidth, mapHeight);
                startGame();
            } else {
                console.error('Background canvas is not ready');
            }
        };

        const handleMapError = (error) => {
            console.error('failed to load map', error);
        };

        map.onload = handleMapLoad;
        map.onerror = handleMapError;

        return () => {
            map.onload = null;
            map.onerror = null;
        };
    }, [map]);
    */
    /*
     const sendPlayerInput = (moveX, moveY, isShooting, aimAngle) => {
         if (socket && socket.readyState === WebSocket.OPEN) {
             const input = new proto.gamestate.PlayerInput();
             input.setMoveX(moveX);
             input.setMoveY(moveY);
             input.setIsShooting(isShooting);
             input.setAimAngle(aimAngle);
 
             const buffer = input.serializeBinary();
             socket.send(buffer);
         }
     }
     */

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
    /*
    useEffect(() => {
        const handleKeyDown = (Event) => {
            switch (Event.key.toLowerCase()) {
                case 'w': setMovement(prev => ({ ...prev, y: -1 })); break;
                case 'a': setMovement(prev => ({ ...prev, x: -1 })); break;
                case 's': setMovement(prev => ({ ...prev, y: 1 })); break;
                case 'd': setMovement(prev => ({ ...prev, x: 1 })); break;
            }
        };

        const handleKeyUp = (Event) => {
            switch (Event.key.toLowerCase()) {
                case 'w':
                case 's': setMovement(prev => ({ ...prev, y: 0 })); break;
                case 'a':
                case 'd': setMovement(prev => ({ ...prev, x: 0 })); break;
            }
            
        };

        window.addEventListener('keydown', handleKeyDown);
        window.addEventListener('keyup', handleKeyUp);

        return () => {
            window.removeEventListener('keydown', handleKeyDown);
            window.removeEventListener('keyup', handleKeyUp);
        };
    }, []);

    useEffect(() => {
        const handleMouseMove = (Event) => {

            if (!canvasRef.current) return;

            const rect = canvasRef.current.getBoundingClientRect();
            const x = Event.clientX - rect.left;
            const y = Event.clientY - rect.top;

            const centerX = canvasRef.current.width / 2;
            const centerY = canvasRef.current.height / 2;
            const angle = Math.atan2(y - centerY, x - centerX);

            
             const rect = Event.target.getBoundingClientRect();
             const x = Event.clientX - rect.left;
             const y = Event.clientY - rect.top;
 
             const centerX = viewportWidth / 2;
             const centerY = viewportHeight / 2;
             const angle = Math.atan2(y - centerY, x - centerX);
             setAimAngle(angle);
             
            setAimAngle(angle);
        };

        const handleMouseDown = () => setIsShooting(true);
        const handleMouseUp = () => setIsShooting(false);

        window.addEventListener('mousemove', handleMouseMove);
        window.addEventListener('mousedown', handleMouseDown);
        window.addEventListener('mouseup', handleMouseUp);

        return () => {
            window.removeEventListener('mousemove', handleMouseMove);
            window.removeEventListener('mousedown', handleMouseDown);
            window.removeEventListener('mouseup', handleMouseUp);
        };
    }, []);
    
    useEffect(() => {
        if (socket) {
            const intervalId = setInterval(() => {
                sendPlayerInput(movement.x, movement.y, isShooting, aimAngle)
            }, 1000 / 60); //60 times a second

            return () => clearInterval(intervalId);
        }
    }, [socket, movement, isShooting, aimAngle]);

    useEffect(() => {

        if (!canvasRef.current) return;
        const canvas = canvasRef.current;
        const ctx = canvas.getContext('2d');

        const world = { width: 3000, height: 3000 };
        const viewport = { width: canvas.width, height: canvas.height };

        canvas.width = window.innerWidth;
        canvas.height = window.innerHeight;

        const drawBackground = () => {
            ctx.fillStyle = 'black';
            ctx.fillRect(0, 0, world.width, world.height);
        };

        const drawPlayers = () => {
            gameState.getPlayersList().forEach(player => {
                ctx.fillStyle = player.getId() == playerID ? 'blue' : 'red';
                const position = player.getPosition();
                if (position) {
                    ctx.fillRect(position.getX(), position.getY(), 20, 20);
                }
            });
        };

        const drawZombies = () => {
            gameState.getZombiesList().forEach(zombie => {
                ctx.fillStyle = 'green';
                const position = zombie.getPosition();
                if (position) {
                    ctx.fillRect(position.getX(), position.getY(), 20, 20);
                }
            });
        };

        const centerViewOnPlayer = () => {
            const player = playerRef.current;
            const xOffset = player.x - viewport.width / 2;
            const yOffset = player.y - viewport.height / 2;

            ctx.setTransform(1, 0, 0, 1, -xOffset, -yOffset);
        };

        const render = () => {
            ctx.clearRect(0, 0, canvas.width, canvas.height);
            drawBackground();
            drawPlayers();
            drawZombies();
            centerViewOnPlayer();
            requestAnimationFrame(render);
        };
        render();
    }, [allPlayersReady, gameState, playerID]);
    */
    /*
    useEffect(() => {
        if (gameState) {
            const currentPlayer = gameState.getPlayersList().find(p => p.getId() == playerID);
            if (currentPlayer) {
                playerRef.current = {
                    ...playerRef.current,
                    x: currentPlayer.getPosition().getX(),
                    y: currentPlayer.getPosition().getY()
                };
            }
        }
    }, [gameState]);
    */
    /*
    const update = () => {
        
        const player = playerRef.current;
        //player.x += movement.x * player.speed;
        //player.y += movement.y * player.speed;

        player.x = Math.max(0, Math.min(player.x, mapWidth - player.width));
        player.y = Math.max(0, Math.min(player.y, mapHeight - player.height));
        

    }
    */
    /*
    const draw = () => {
        const dynamicCanvas = dynamicCanvasRef.current;
        const dynamicCtx = dynamicCanvas.getContext('2d');
        //const backgroundCanvas = backgroundCanvasRef.current;

        dynamicCtx.clearRect(0, 0, dynamicCanvas.width, dynamicCanvas.height);
        const player = playerRef.current;
        const viewportX = Math.max(0, Math.min(player.x + player.width / 2 - dynamicCanvas.width / 2, mapWidth - dynamicCanvas.width));
        const viewportY = Math.max(0, Math.min(player.y + player.height / 2 - dynamicCanvas.height / 2, mapHeight - dynamicCanvas.height));

        //dynamicCtx.drawImage(backgroundCanvas, viewportX, viewportY, dynamicCanvas.width, dynamicCanvas.height, 0, 0, dynamicCanvas.width, dynamicCanvas.height);

        if (gameState) {
            gameState.getPlayersList().forEach(player => {
                dynamicCtx.fillStyle = player.getId() == playerID ? 'blue' : 'red';
                const position = player.getPosition();
                if (position) {
                    dynamicCtx.fillRect(position.getX() - viewportX, position.getY() - viewportY, 20, 20);
                }
            });

            gameState.getZombiesList().forEach(zombie => {
                dynamicCtx.fillStyle = 'green';
                const position = zombie.getPosition();
                if (position) {
                    dynamicCtx.fillRect(position.getX() - viewportX, position.getY() - viewportY, 20, 20);
                }
            });

            const map = gameState.getMap();
            if (map) {
                dynamicCtx.fillStyle = 'black';
                dynamicCtx.fillText(`Map: ${map.getName()}`, 10, 10);
            }
        }
    };

    const startGame = () => {
        const gameLoop = () => {
            update();
            draw();
            requestAnimationFrame(gameLoop);
        };
        gameLoop();
    };
    */
    /*
    useEffect(() => {
        if (allPlayersReady && map.complete && dynamicCanvasRef.current) {
            startGame();
        }
    }, [allPlayersReady, gameState, playerID]);
    */
    /*
    useEffect(() => {
        
        const startGame = () => {
            const canvas = canvasRef.current;
            const ctx = canvas.getContext('2d');

            const gameLoop = () => {
                ctx.clearRect(0, 0, canvas.width, canvas.height);

                if (gameState) {
                    gameState.getPlayersList().forEach(player => {
                        ctx.fillStyle = player.getId() == playerID ? 'blue' : 'red';
                        const position = player.getPosition();
                        if (position) {
                            ctx.fillRect(position.getX(), position.getY(), 20, 20);
                        }
                    });

                    gameState.getZombiesList().forEach(zombie => {
                        ctx.fillStyle = 'green';
                        const position = zombie.getPosition();
                        if (position) {
                            ctx.fillRect(position.getX(), position.getY(), 20, 20);
                        }
                    });

                    const map = gameState.getMap();
                    if (map) {
                        ctx.fillStyle = 'black';
                        ctx.fillText(`Map: ${gameState.getMap()}`, 10, 10);
                    }
                }
                requestAnimationFrame(gameLoop);
            };
            gameLoop();
        };
        if (allPlayersReady) {
            startGame();
        }

    }, [allPlayersReady, gameState, playerID]);
    */

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
            <GameCanvas playerID={playerID} gameState={gameState} socket={socket} />
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
                        <CharacterSprite character={player.character} />
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