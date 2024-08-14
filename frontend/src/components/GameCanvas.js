/* eslint-disable no-undef */
import React, { useRef, useEffect, useState, useCallback } from "react";
import * as PIXI from 'pixi.js';
import * as _ from '../utils/gamestate_pb';
import '../style/rootstyle.css';

//now there is no initial game state, need to correct logic
const GameCanvas = ({ playerID, initialGameState, socket }) => {
    const canvasRef = useRef(null);
    const appRef = useRef(null);
    const [assetsLoaded, setAssetsLoaded] = useState(false);
    const [movement, setMovement] = useState({ x: 0, y: 0 });
    const movementRef = useRef({ x: 0, y: 0 });
    const isShootingRef = useRef(false);
    const aimAngleRef = useRef(0);
    const playerSprites = useRef(new Map());
    const zombieSprites = useRef(new Map());
    const bulletSprites = useRef(new Map());
    const gameStateRef = useRef(null);
    const aimLineRef = useRef(null);
    const allowRenderRef = useRef(false);
    const [pixiReady, setPixiReady] = useState(false);
    const [gameStateReceived, setGameStateReceived] = useState(false);
    const animationsRef = useRef({});

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

    useEffect(() => {
        if (socket) {
            const updateGameState = async (e) => {
                try {
                    const arrayBuffer = await e.data.arrayBuffer();
                    const array = new Uint8Array(arrayBuffer);
                    const protoGameState = proto.gamestate.GameState.deserializeBinary(array);
                    gameStateRef.current = protoGameState;
                    allowRenderRef.current = true;
                    if (!gameStateReceived) {
                        setGameStateReceived(true);
                    }
                } catch (error) {
                    console.error("error deserializng binary message");
                }
            };

            socket.addEventListener('message', updateGameState);

            return () => {
                socket.removeEventListener('message', updateGameState);
            }
        }
    }, [socket, gameStateReceived]);

    const sendPlayerInput = () => {
        if (socket && socket.readyState === WebSocket.OPEN) {
            const input = new proto.gamestate.PlayerInput();
            input.setMoveX(movementRef.current.x);
            input.setMoveY(movementRef.current.y);
            input.setIsShooting(isShootingRef.current);
            input.setAimAngle(aimAngleRef.current);

            const buffer = input.serializeBinary();
            socket.send(buffer);
        } else {
            console.log("websocket busted");
        }
    };

    const updateAnimation = (moveX, Id) => {
        const playerSprite = playerSprites.current.get(parseInt(Id));
        if (!playerSprite) {
            console.log("player sprite not found");
            return;
        }
        let newTextures;
        if (moveX < 0) {
            newTextures = animationsRef.current.druggieRunLeft.textures;
            console.log("switching to run left");
            console.log("new textures", newTextures)
        } else if (moveX > 0) {
            newTextures = animationsRef.current.druggieRunRight.textures;
            console.log("switching to run right");
            console.log("new textures", newTextures)
        } else {
            newTextures = animationsRef.current.druggieIdleRight.textures;
            console.log("switching to idle right");
            console.log("new textures", newTextures)
        }

        if (playerSprite.textures !== newTextures) {
            playerSprite.textures = newTextures;
            playerSprite.play();
            console.log("textures updated and animation played");
        }
    };

    const updatePositions = () => {
        console.log("in update positions func");
        const gameState = gameStateRef.current;
        if (!gameState) return; // no gamestate

        const currentPlayerIds = new Set(gameState.getPlayersList().map(player => player.getId()));
        const currentZombieIds = new Set(gameState.getZombiesList().map(zombie => zombie.getId()));
        const currentBulletIds = new Set(gameState.getBulletsList().map(bullet => bullet.getId()));
        let playerPosition = null;

        //console.log("current zombies in gamestate: ", Array.from(currentZombieIds));

        playerSprites.current.forEach((sprite, playerId) => {
            if (!currentPlayerIds.has(playerId)) {
                appRef.current.stage.removeChild(sprite);
                playerSprites.current.delete(playerId);
            }
        });

        gameState.getPlayersList().forEach(player => {
            const position = player.getPosition();
            const sprite = playerSprites.current.get(player.getId());
            console.log("movement ref", movementRef.current.x, movementRef.current.y);
            if (position && sprite) {
                //console.log(`player x: ${position.getX()}, y: ${position.getY()}`);
                sprite.x = position.getX();
                sprite.y = position.getY();
                if (player.getId() === parseInt(playerID)) {
                    playerPosition = position;
                    playerRef.current.x = position.getX();
                    playerRef.current.y = position.getY();
                    updateAnimation(movementRef.current.x, playerID);
                    //updatePlayerSprites();
                } else {
                    updateAnimation(player.getMoveX(), player.getId());
                }
            }
        });

        zombieSprites.current.forEach((sprite, zombieId) => {
            if (!currentZombieIds.has(zombieId)) {
                appRef.current.stage.removeChild(sprite);
                zombieSprites.current.delete(zombieId);
            }
        });

        gameState.getZombiesList().forEach(zombie => {
            const position = zombie.getPosition();
            const sprite = zombieSprites.current.get(zombie.getId());
            if (position && sprite) {
                //console.log(`zombie x: ${position.getX()}, y: ${position.getY()}`);
                sprite.x = position.getX();
                sprite.y = position.getY();
            } else if (position) {
                const texture = PIXI.Texture.from("testZom");
                const newSprite = new PIXI.Sprite(texture);
                newSprite.width = 50;
                newSprite.height = 50;
                newSprite.anchor.set(0.5);
                newSprite.x = position.getX();
                newSprite.y = position.getY();
                appRef.current.stage.addChild(newSprite);
                zombieSprites.current.set(zombie.getId(), newSprite);
            }
        });

        bulletSprites.current.forEach((sprite, bulletId) => {
            if (!currentBulletIds.has(bulletId)) {
                appRef.current.stage.removeChild(sprite);
                bulletSprites.current.delete(sprite);
            }
        });

        gameState.getBulletsList().forEach(bullet => {
            const position = bullet.getPosition();
            const sprite = bulletSprites.current.get(bullet.getId());
            if (position && sprite) {
                sprite.x = position.getX();
                sprite.y = position.getY();
            } else if (position) {
                const bulletTexture = PIXI.Texture.from('bullet');
                const bulletSprite = new PIXI.Sprite(bulletTexture);
                bulletSprite.width = 10;
                bulletSprite.height = 10;
                bulletSprite.anchor.set(0.5);
                bulletSprite.x = position.getX();
                bulletSprite.y = position.getY();
                appRef.current.stage.addChild(bulletSprite);
                bulletSprites.current.set(bullet.getId(), bulletSprite);
            }
        });

        if (playerPosition) {
            const app = appRef.current;
            if (app) {
                app.stage.x = window.innerWidth / 2 - playerPosition.getX();
                app.stage.y = window.innerHeight / 2 - playerPosition.getY();

                //TESTING AIMING
                if (aimLineRef.current) {
                    aimLineRef.current.clear();
                    aimLineRef.current.moveTo(playerPosition.getX(), playerPosition.getY());
                    aimLineRef.current.lineTo(
                        playerPosition.getX() + Math.cos(aimAngleRef.current) * 1000,
                        playerPosition.getY() + Math.sin(aimAngleRef.current) * 1000
                    );
                }
                //TESTING
            }
        }
    };

    useEffect(() => {
        const initPixi = async () => {
            if (!canvasRef.current) return;

            const app = new PIXI.Application();
            await app.init({
                width: window.innerWidth,
                height: window.innerHeight,
                resolution: window.devicePixelRatio || 1,
            });
            canvasRef.current.appendChild(app.canvas);
            appRef.current = app;

            PIXI.Assets.add({ alias: 'testmap', src: '/testmap.png' });
            PIXI.Assets.add({ alias: 'testChar', src: '/soldiertestsprite.png' });
            PIXI.Assets.add({ alias: 'testZom', src: '/zombietestsprite.png' });
            PIXI.Assets.add({ alias: 'bullet', src: '/bullet.png' });
            PIXI.Assets.add({ alias: 'druggie_left_not_firing.png', src: '/sprites/druggie_left_not_firing.png' });
            PIXI.Assets.add({ alias: 'druggie_gun_run_left_sheet', src: '/sprites/druggie_left_not_firing_spritesheet.json' });
            PIXI.Assets.add({ alias: 'druggie_right_not_firing.png', src: '/sprites/druggie_right_not_firing.png' });
            PIXI.Assets.add({ alias: 'druggie_gun_run_right_sheet', src: '/sprites/druggie_right_not_firing_spritesheet.json' });
            PIXI.Assets.add({ alias: 'druggie_gun_idle_left.png', src: '/sprites/druggie_gun_idle_left.png' });
            PIXI.Assets.add({ alias: 'druggie_gun_idle_left_sheet', src: '/sprites/druggie_idle_gun_left_spritesheet.json' });
            PIXI.Assets.add({ alias: 'druggie_gun_idle_right.png', src: '/sprites/druggie_gun_idle_right.png' });
            PIXI.Assets.add({ alias: 'druggie_gun_idle_right_sheet', src: '/sprites/druggie_idle_gun_right_spritesheet.json' });

            await PIXI.Assets.load('testmap');
            await PIXI.Assets.load('testChar');
            await PIXI.Assets.load('testZom');
            await PIXI.Assets.load('bullet');
            await PIXI.Assets.load('druggie_left_not_firing.png');
            await PIXI.Assets.load('druggie_gun_run_left_sheet');
            await PIXI.Assets.load('druggie_right_not_firing.png');
            await PIXI.Assets.load('druggie_gun_run_right_sheet');
            await PIXI.Assets.load('druggie_gun_idle_left.png');
            await PIXI.Assets.load('druggie_gun_idle_left_sheet');
            await PIXI.Assets.load('druggie_gun_idle_right.png');
            await PIXI.Assets.load('druggie_gun_idle_right_sheet');
            //const sheet = PIXI.Assets.get('druggie_gun_run_left_sheet');
            //console.log("sprite sheet", sheet);
            //console.log("animation name", Object.keys(sheet.animations)[0]);
            const druggieIdleLeft = PIXI.Assets.get('druggie_gun_idle_left_sheet');
            const druggieIdleRight = PIXI.Assets.get('druggie_gun_idle_right_sheet');
            const druggieRunLeft = PIXI.Assets.get('druggie_gun_run_left_sheet');
            const druggieRunRight = PIXI.Assets.get('druggie_gun_run_right_sheet');

            const druggieIdleLeftSheet = new PIXI.Spritesheet(PIXI.Texture.from(druggieIdleLeft.data.meta.image), druggieIdleLeft.data);
            const druggieIdleRightSheet = new PIXI.Spritesheet(PIXI.Texture.from(druggieIdleRight.data.meta.image), druggieIdleRight.data);
            const druggieRunLeftSheet = new PIXI.Spritesheet(PIXI.Texture.from(druggieRunLeft.data.meta.image), druggieRunLeft.data);
            const druggieRunRightSheet = new PIXI.Spritesheet(PIXI.Texture.from(druggieRunRight.data.meta.image), druggieRunRight.data);

            await druggieIdleLeftSheet.parse();
            await druggieIdleRightSheet.parse();
            await druggieRunLeftSheet.parse();
            await druggieRunRightSheet.parse();

            const animations = {
                druggieIdleLeft: new PIXI.AnimatedSprite(druggieIdleLeftSheet.animations.animation),
                druggieIdleRight: new PIXI.AnimatedSprite(druggieIdleRightSheet.animations.animation),
                druggieRunLeft: new PIXI.AnimatedSprite(druggieRunLeftSheet.animations.animation),
                druggieRunRight: new PIXI.AnimatedSprite(druggieRunRightSheet.animations.animation)
            }
            animationsRef.current = animations;
            console.log("anim", animations);
            console.log("animref", animationsRef.current)

            setPixiReady(true);
        };
        if (gameStateReceived) {
            initPixi();
        }
        console.log("pixi init");
        return () => {
            if (appRef.current) {
                appRef.current.destroy(true, { children: true, texture: true, baseTexture: true });
                appRef.current = null;
            }
        };
    }, [gameStateReceived]);

    useEffect(() => {
        console.log("loading")
        if (pixiReady) {
            //console.log("pixi ready")
            const app = appRef.current;
            const gameState = gameStateRef.current;

            //const playerTexture = PIXI.Texture.from('testChar');
            const zombieTexture = PIXI.Texture.from('testZom');

            const mapTexture = PIXI.Texture.from('testmap');
            const mapSprite = new PIXI.Sprite(mapTexture);
            app.stage.addChild(mapSprite);

            gameState.getPlayersList().forEach(async (player) => {
                //const sprite = new PIXI.Sprite(playerTexture);
                //const sprite = await createAnimatedSprite('druggie_gun_idle_right_sheet');
                //const sprite = animationsRef.current.druggieIdleRight;
                const sprite = new PIXI.AnimatedSprite(animationsRef.current.druggieIdleRight.textures);
                sprite.width = 50;
                sprite.height = 50;
                const position = player.getPosition();
                if (position) {
                    sprite.x = position.getX();
                    sprite.y = position.getY();
                }
                sprite.anchor.set(0.5);
                sprite.animationSpeed = 0.1;
                sprite.play();
                app.stage.addChild(sprite);
                playerSprites.current.set(player.getId(), sprite);
            });

            gameState.getZombiesList().forEach(zombie => {
                const sprite = new PIXI.Sprite(zombieTexture);
                sprite.width = 50;
                sprite.height = 50;
                const position = zombie.getPosition();
                if (position) {
                    sprite.x = position.getX();
                    sprite.y = position.getY();
                }
                sprite.anchor.set(0.5);
                app.stage.addChild(sprite);
                zombieSprites.current.set(zombie.getId(), sprite);
            });
            setAssetsLoaded(true);
            app.ticker.add(() => {
                updatePositions();
                //updatePlayerSprites();
                sendPlayerInput();
            });
        }
    }, [pixiReady]);

    useEffect(() => {

        const handleKeyDown = (e) => {
            console.log("handling key down");
            switch (e.key.toLowerCase()) {
                case 'w': movementRef.current.y = -1; break;
                case 'a': movementRef.current.x = -1; break;
                case 's': movementRef.current.y = 1; break;
                case 'd': movementRef.current.x = 1; break;
            }
            //updateAnimation(movementRef.current.x, movementRef.current.y);
        };

        const handleKeyUp = (e) => {
            console.log("handling key up");
            switch (e.key.toLowerCase()) {
                case 'w':
                case 's': movementRef.current.y = 0; break;
                case 'a':
                case 'd': movementRef.current.x = 0; break;
            }
            //updateAnimation(movementRef.current.x, movementRef.current.y);
        };

        window.addEventListener('keydown', handleKeyDown);
        window.addEventListener('keyup', handleKeyUp);

        // Clean up on unmount
        return () => {
            window.removeEventListener('keydown', handleKeyDown);
            window.removeEventListener('keyup', handleKeyUp);
        };
    }, []);

    useEffect(() => {
        //if (!appRef.current) return;
        const handleMouseMove = (e) => {
            console.log("handling mouse move");
            const playerPos = playerRef.current;
            const canvasBounds = appRef.current.canvas.getBoundingClientRect();
            const scale = appRef.current.stage.scale.x;
            const mouseX = e.clientX - canvasBounds.left;
            const mouseY = e.clientY - canvasBounds.top;

            const worldMouseX = (mouseX - (window.innerWidth / 2 - playerPos.x * scale)) / scale;
            const worldMouseY = (mouseY - (window.innerHeight / 2 - playerPos.y * scale)) / scale;

            const dx = worldMouseX - playerPos.x;
            const dy = worldMouseY - playerPos.y;

            aimAngleRef.current = Math.atan2(dy, dx);

        };

        const handleMouseDown = () => isShootingRef.current = true; //setIsShooting(true); 
        const handleMouseUp = () => isShootingRef.current = false; //setIsShooting(false);

        window.addEventListener('mousemove', handleMouseMove);
        window.addEventListener('mousedown', handleMouseDown);
        window.addEventListener('mouseup', handleMouseUp);

        return () => {
            window.removeEventListener('mousemove', handleMouseMove);
            window.removeEventListener('mousedown', handleMouseDown);
            window.removeEventListener('mouseup', handleMouseUp);
        }

    }, []);

    return <div ref={canvasRef} />;
};

export default GameCanvas;