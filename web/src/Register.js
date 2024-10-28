// src/Register.js
import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';

const Register = () => {
    const [username, setUsername] = useState('');
    const [password, setPassword] = useState('');
    const [email, setEmail] = useState('');
    const [message, setMessage] = useState('');
    const [rememberMe, setRememberMe] = useState(false);
    const [errorMessage, setErrorMessage] = useState('');
    const navigate = useNavigate();

    const handleRegister = async (e) => {
    e.preventDefault(); // предотвращаем перезагрузку страницы

    try {
        const response = await fetch('http://localhost:8080/api/register', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({
                username: username,
                email: email,
                password: password,
            }),
        });

        if (response.ok) {
            // Если сервер перенаправляет, переходим на страницу входа
            navigate('/login');
        } else {
            const data = await response.json(); // ждем завершения
            setErrorMessage(data.error || 'Произошла ошибка. Попробуйте еще раз.');
        }
    } catch (error) {
        console.error('Error:', error);
        setErrorMessage('Не удалось выполнить запрос. Попробуйте позже.');
    }
};

    const handleLoginRedirect = () => {
        
        navigate('/login');
    };

    

    return (
        <section> 
            <form id="app" onSubmit={handleRegister}>
                <h1>Register</h1>
                <div className="inputbox">
                    <ion-icon name="mail-outline"></ion-icon>
                    <input
                        id="register-email"
                        type="email"
                        required
                        value={email}
                        onChange={(e) => setEmail(e.target.value)}
                    />
                    <label>Email</label>
                </div>
                <div className="inputbox">
                    <ion-icon name="username"></ion-icon>
                    <input
                        id="register-username"
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
                        id="register-password"
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
                <button id="register-button" type="submit">Sign up</button>
                <div className="register">
                    <p>
                        Already have an account? 
                        <button type="button" onClick={handleLoginRedirect}>Log in</button>
                    </p>
                </div>
                {message && <p>{message}</p>} {/* Отображение сообщения */}
            </form>
        </section>
    );

};

export default Register;






