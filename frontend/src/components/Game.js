import { useEffect, useState } from "react";
import { w3cwebsocket as W3CWebSocket } from "websocket";

const Game = () => {
    const [client, setClient] = useState(null);
    const [roomID, setRoomID] = useState("");
    const [isConnected, setIsConnected] = useState(false);

    useEffect(() => {
        const client = new W3CWebSocket("ws://localhost:8080/ws");
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

    return (
        <div>
            <h2>zombros</h2>
            {isConnected ? (
                <>
                    <button onClick={createRoom}>Create Room</button>
                    <input
                        type="text"
                        placeholder="Enter Room ID"
                        value={roomID}
                        onChange={(e) => setRoomID(e.target.value)}
                    />
                    <button onClick={() => joinRoom(roomID)}>Join Room</button>
                </>
            ) : (
                <p>Connecting to websocket...</p>
            )}
            <canvas id="gameCanvas"></canvas>
        </div>
    );
};

export default Game;
