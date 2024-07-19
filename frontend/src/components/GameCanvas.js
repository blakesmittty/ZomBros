/* eslint-disable no-undef */
import React, { useRef, useEffect, useState } from "react";
import * as PIXI from 'pixi.js';
import * as _ from '../utils/gamestate_pb';

const GameCanvas = ({ playerID, gameState, socket }) => {
    const canvasRef = useRef(null);
    const appRef = useRef(null);
    const [assetsLoaded, setAssetsLoaded] = useState(false);
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

    const players = useRef({});
    const zombies = useRef({});

    const sendPlayerInput = () => {
        if (socket && socket.readyState === WebSocket.OPEN) {
            const input = new proto.gamestate.PlayerInput();
            input.setMoveX(playerRef.current.moveX);
            input.setMoveY(playerRef.current.moveY);
            input.setIsShooting(playerRef.current.isShooting);
            input.setAimAngle(playerRef.current.aimAngle);

            const buffer = input.serializeBinary();
            socket.send(buffer);
        }
    }

    useEffect(() => {
        if (!canvasRef.current) return;


        (async () => {
            const app = new PIXI.Application();
            await app.init({
                width: window.innerWidth,
                height: window.innerHeight,
                resolution: window.devicePixelRatio || 1,
            });
            canvasRef.current.appendChild(app.canvas);
            appRef.current = app;


            if (appRef.current) {
                PIXI.Assets.add({ alias: 'testmap', src: '/testmap.png' });
                PIXI.Assets.add({ alias: 'testChar', src: '/soldiertestsprite.png' });
                PIXI.Assets.add({ alias: 'testZom', src: '/zombietestsprite.png' });
                console.log('map loading');
                const map = await PIXI.Assets.load('testmap');
                console.log('map loaded', map);

                console.log('char loading');
                const char = await PIXI.Assets.load('testChar');
                console.log('char loaded', char);

                console.log('zom loading');
                const zom = await PIXI.Assets.load('testZom');
                console.log('zom loaded', zom);

                const playerTexture = PIXI.Texture.from('testChar');
                const playerSprite = new PIXI.Sprite(playerTexture);

                const zombieTexture = PIXI.Texture.from('testZom');
                const zombieSprite = new PIXI.Sprite(zombieTexture);

                const mapTexture = PIXI.Texture.from('testmap');
                const mapSprite = new PIXI.Sprite(mapTexture);

                appRef.current.stage.addChild(mapSprite);
                appRef.current.stage.addChild(playerSprite);
                appRef.current.stage.addChild(zombieSprite);
                setAssetsLoaded(true);
            }

        })();

        return () => {
            if (appRef.current) {
                appRef.current.destroy(true, { children: true, texture: true, baseTexture: true });
                appRef.current = null;
            }
        };

    }, []);

    useEffect(() => {
        const handleKeyDown = (Event) => {
            switch (Event.key.toLowerCase()) {
                case 'w': playerRef.current.moveY = -1; break;
                case 'a': playerRef.current.moveX = -1; break;
                case 's': playerRef.current.moveY = 1; break;
                case 'd': playerRef.current.moveX = 1; break;
                default: break;
            }
            sendPlayerInput();
        };

        const handleKeyUp = (Event) => {
            switch (Event.key.toLowerCase()) {
                case 'w': case 's': playerRef.current.moveY = 0; break;
                case 'a': case 'd': playerRef.current.moveX = 0; break;
                default: break;
            }
            sendPlayerInput();
        };

        window.addEventListener('keydown', handleKeyDown);
        window.addEventListener('keyup', handleKeyUp);

        // Clean up on unmount
        return () => {
            window.removeEventListener('keydown', handleKeyDown);
            window.removeEventListener('keyup', handleKeyUp);
        };
    }, []);

    return <div ref={canvasRef} />;
};

export default GameCanvas;