// account.js
document.addEventListener('DOMContentLoaded', () => {
    const usernameDisplay = document.getElementById('username-display');
    const sendButton = document.getElementById('send');
    const messageInput = document.getElementById('message');
    const messagesDiv = document.getElementById('messages');

    // Получаем имя пользователя из URL
    const urlParams = new URLSearchParams(window.location.search);
    const username = urlParams.get('username');
    usernameDisplay.textContent = `Logged in as: ${username}`;

    // Подключаемся к WebSocket
    const socket = new WebSocket('ws://localhost:8080/chat');

    socket.onopen = function() {
        console.log('Соединение установлено');
    };

    socket.onmessage = function(event) {
        const message = JSON.parse(event.data);
        const messageElement = document.createElement('div');
        messageElement.textContent = `${message.user}: ${message.content}`;
        messagesDiv.appendChild(messageElement);
    };

    sendButton.onclick = function() {
        const message = {
            user: username,
            content: messageInput.value
        };
        socket.send(JSON.stringify(message));
        messageInput.value = ''; // Очистка поля ввода
    };

    socket.onclose = function() {
        console.log('Соединение закрыто');
    };
});
