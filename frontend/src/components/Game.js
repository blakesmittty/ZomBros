import { useEffect, useState } from "react";
import { w3cwebsocket as W3CWebSocket } from "websocket";
import { useHistory, useNavigate } from 'react-router-dom'


const Game = () => {
    //const [client, setClient] = useState(null);
    const [roomID, setRoomID] = useState("");
    //const [isConnected, setIsConnected] = useState(false);
    const playerId = localStorage.getItem('playerId');
    const navigate = useNavigate();


    const handleCreateRoom = async () => {
        try {
            const ws = new WebSocket(`ws://localhost:8080/ws?playerId=${playerId}`);
            ws.onopen = () => {
                ws.send('create');
                console.log('websocket connection opened');
            };
            ws.onmessage = (e) => {
                if (e.data.startsWith('Error:')) {
                    alert(e.data);
                } else {
                    const newRoomID = e.data;
                    console.log(`room created with ID: ${newRoomID}`);
                    ws.close();
                    navigate(`/room/${newRoomID}`);
                }
            };
            ws.onclose = (e) => {
                console.error("websocket closed unexpectedly:", e);
            };
            ws.onerror = (error) => {
                console.error('websocket error: ', error);
            };
        } catch (error) {
            console.error('Error creating room:', error);
        }
    };

    const handleJoinRoom = async () => {
        if (!roomID) {
            alert('please enter a room ID dude! for god\'s sake!');
            return;
        }
        try {
            const ws = new WebSocket(`ws://localhost:8080/ws?playerId=${playerId}`);
            ws.onopen = () => {
                ws.send(`join:${roomID}`);
                console.log('websocket connection opened');
            };
            ws.onmessage = (e) => {
                if (e.data.startsWith(`Error:`)) {
                    alert(e.data);
                } else {
                    console.log(`joined room with ID: ${roomID}`);
                    navigate(`/room/${roomID}`);
                }
            };
            ws.onclose = (e) => {
                console.error('websocket closed unexpectedly:', e);
            };
            ws.onerror = (error) => {
                console.error('websocket error:', error);
            };
        } catch (error) {
            console.error('Error joining room: ', error);
        }
    };



    /*
    useEffect(() => {
        if (!playerId) {
            console.error('No playerID found in local storage');
            return;
        }
        const client = new W3CWebSocket(`ws://localhost:8080/ws?playerId=${playerId}`);
        setClient(client);

        client.onopen = () => {
            console.log("websocket client connected");
            setIsConnected(true);
        };

        client.onmessage = (message) => {
            console.log(message.data);

        };

        client.onclose = () => {
            console.log("websocket client disconnected");
            setIsConnected(false);
        };

        client.onerror = (error) => {
            console.error('websocket error: ', error);
        }

        return () => {
            if (client.readyState === 1) {
                client.close();
            }
        };

    }, []);

    const createRoom = () => {
        if (client) {
            client.send('create');
        }
    };

    const joinRoom = () => {
        if (client) {
            client.send(`join:${roomID}`)
        }
    };
    */


    return (
        <div>
            <h2>game</h2>
            <h2>playerID: {playerId}</h2>
            <button onClick={handleCreateRoom}>Create Room</button>
            <input
                type="text"
                placeholder="Enter Room ID"
                value={roomID}
                onChange={(e) => setRoomID(e.target.value)}
            />
            <button onClick={handleJoinRoom}>Join Room</button>

        </div>
    );
};

export default Game;
