// src/Login.js
import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';

const Login = () => {
    const [username, setUsername] = useState('');
    const [password, setPassword] = useState('');
    const [rememberMe, setRememberMe] = useState(false);
    const [message, setMessage] = useState('');
    const [errorMessage, setErrorMessage] = useState('');
    const navigate = useNavigate();

    const handleLogin = async (e) => {
    e.preventDefault(); // предотвращаем перезагрузку страницы

    console.log(username, password);

    try {
        const response = await fetch('http://localhost:8080/api/login', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({
                username: username,
                password: password,
            }),
        });

        if (response.ok) {
            // Если сервер перенаправляет, переходим на страницу аккаунта
            navigate(`/account?username=${encodeURIComponent(username)}`);
        } else {
            const data = await response.json(); // ждем завершения
            setErrorMessage(data.error || 'Неверные учетные данные. Попробуйте еще раз.');
        }
    } catch (error) {
        console.error('Error:', error);
        setErrorMessage('Не удалось выполнить запрос. Попробуйте позже.');
    }
};
    const handleRegisterRedirect = () => {
        
        navigate('/register');
    }

    return (
        <section>
            <form id="app" onSubmit={handleLogin}>
                <h1>Login</h1>
                <div className="inputbox">
                    <ion-icon name="mail-outline"></ion-icon>
                    <input
                        id="login-username"
                        type="text"
                        required
                        value={username}
                        onChange={(e) => setUsername(e.target.value)}
                    />
                    <label>Username</label>
                </div>

                <div className="inputbox">
                    <ion-icon name="lock-closed-outline"></ion-icon>
                    <input
                        id="login-password"
                        type="password"
                        required
                        value={password}
                        onChange={(e) => setPassword(e.target.value)}
                    />
                    <label>Password</label>
                </div>
                <div className="forget">
                    <label>
                        <input
                            type="checkbox"
                            checked={rememberMe}
                            onChange={() => setRememberMe(!rememberMe)}
                        />
                        Remember me
                    </label>
                    <a href="#">Forget Password</a>
                </div>
                <button id="login-button" type="submit">Log in</button>
                <div className="register">
                    <p>
                        Don't have an account? 
                        <button type="button" onClick={handleRegisterRedirect}>Register</button>
                    </p>
                </div>
                {message && <p>{message}</p>} {/* Отображение сообщения */}
            </form>
        </section>
    );

};

export default Login;