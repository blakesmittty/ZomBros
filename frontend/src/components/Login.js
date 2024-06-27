import React, { useState } from 'react';
import { Navigate, useNavigate } from 'react-router-dom';

function Login({ setLoggedIn }) {
    const [username, setUsername] = useState('');
    const [password, setPassword] = useState('');
    const [message, setMessage] = useState('');
    const playerId = localStorage.getItem('playerId');
    const navigate = useNavigate();

    const handleSubmit = async (e) => {
        e.preventDefault();
        try {
            const response = await fetch('http://localhost:8080/login', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/x-www-form-urlencoded',
                },
                body: new URLSearchParams({
                    username,
                    password,
                }),
            });
            const result = await response.json();
            if (response.ok) {
                setMessage(result.message);
                setLoggedIn(result.id);
                navigate('/game');
            } else {
                setMessage(result.message || 'login failed. please try again');
            }
        } catch (error) {
            setMessage('an error occurred, please try again');
            console.error('Failed to fetch:', error);
        }
    };

    return (
        <div>
            <h2>Login</h2>
            <h2>playerID: {playerId}</h2>
            <form onSubmit={handleSubmit}>
                <input
                    type='text'
                    placeholder='Username'
                    value={username}
                    onChange={(e) => setUsername(e.target.value)}
                    required
                />
                <input
                    type='password'
                    placeholder='Password'
                    value={password}
                    onChange={(e) => setPassword(e.target.value)}
                    required
                />
                <button type='submit'>Login</button>
            </form>
            {message && <p>{message}</p>}
        </div>
    );
}

export default Login;