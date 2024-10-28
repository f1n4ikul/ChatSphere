// src/Account.js
import React, { useEffect, useState } from 'react';
import { useNavigate } from 'react-router-dom';

const Account = () => {
    const [userData, setUserData] = useState(null);
    const [errorMessage, setErrorMessage] = useState('');
    const navigate = useNavigate();

    useEffect(() => {
        const fetchUserData = async () => {
            const username = new URLSearchParams(window.location.search).get('username');

            if (!username) {
                setErrorMessage('Не удалось получить информацию о пользователе.');
                return;
            }

            try {
                const response = await fetch(`http://localhost:8080/api/user/${username}`);
                if (!response.ok) {
                    throw new Error('Ошибка получения данных пользователя');
                }
                const data = await response.json();
                setUserData(data); // Предполагается, что data содержит { username }
            } catch (error) {
                console.error('Error:', error);
                setErrorMessage('Ошибка загрузки данных пользователя.');
            }
        };

        fetchUserData();
    }, []);

    const handleLogout = () => {
        // Здесь можно добавить логику выхода из системы
        // Например, очистка токена или данных пользователя из состояния
        navigate('/login'); // Перенаправление на страницу входа
    };
    
    const handleChat = () => {
        navigate('/chats');
    }

    return (
        <section id='app'>
            <h1>Страница аккаунта</h1>
            {errorMessage && <p style={{ color: 'red' }}>{errorMessage}</p>}
            {userData ? (
                <div>
                    <h2>Добро пожаловать, {userData.username}!</h2>
                    <button onClick={handleChat} >Чаты</button>
                    {/* Дополнительная информация о пользователе может быть добавлена здесь */}
                    <button onClick={handleLogout}>Выйти</button>
                </div>
            ) : (
                <p>Загрузка данных...</p>
            )}
        </section>
    );
};

export default Account;
