document.addEventListener('DOMContentLoaded', () => {
    const registerButton = document.getElementById('register-button');
    const sendButton = document.getElementById('send');
    const messageInput = document.getElementById('message');
    const messagesDiv = document.getElementById('messages');
    const chatDiv = document.getElementById('chat');


        registerButton.onclick = function() {
        const email = document.getElementById('register-email').value;
        const password = document.getElementById('register-password').value;
        const username = document.getElementById('register-username').value;
            

        fetch('http://localhost:8080/register', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/x-www-form-urlencoded',
            },
            body: new URLSearchParams({
                username: username,
                password: password,
                email: email
            })
        })
        .then(response => response.json())
        .then(data => {
            alert(data.message); // Показываем сообщение
            // Перенаправление на страницу аккаунта после регистрации
            window.location.href = `login.html`;
        })
        .catch(error => console.error('Error:', error));
    };



    

    // Функция для подключения к WebSocket
    function connectToChat(username) {
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
    }
});

