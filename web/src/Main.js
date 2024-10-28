import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom'; // Импортируем useNavigate
import './reg.css'; 


const Main = () => {
    const navigate = useNavigate();


    const handleRegisterRedirect = () => {
        navigate('/register'); // Перенаправляем на страницу регистрации
    };

    const handleLoginRedirect = () => {
        navigate('/login'); // Перенаправляем на страницу входа
        };
    

    return (
        <div class="mainfile">
            <h1>Welcome</h1>
            <button onClick={handleRegisterRedirect}>Register</button>
            <button onClick={handleLoginRedirect}>Login</button>
        </div>
    )
}

export default Main;