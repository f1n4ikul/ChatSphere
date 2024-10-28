import React from 'react';
import ReactDOM from 'react-dom';
import App from './App'; // Импортируйте основной компонент App

// Рендеринг основного приложения
ReactDOM.render(
    <React.StrictMode>
        <App />
    </React.StrictMode>,
    document.getElementById('root') // Убедитесь, что в вашем HTML есть элемент с id="root"
);