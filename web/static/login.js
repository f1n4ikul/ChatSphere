document.addEventListener('DOMContentLoaded', () => {
    const loginButton = document.getElementById('login-button');
    const sendButton = document.getElementById('send');
    const messageInput = document.getElementById('message');
    const messagesDiv = document.getElementById('messages');
    const chatDiv = document.getElementById('chat');




        loginButton.onclick = function() {
            // const email = document.getElementById('login-email').value;
            const password = document.getElementById('login-password').value;
            const username= document.getElementById('login-username').value;

            fetch('http://localhost:8080/login', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/x-www-form-urlencoded',
                },
                body: new URLSearchParams({
                    // email: email, 
                    username: username,
                    password: password
                })
            })
            .then(response => response.json())
            .then(data => {
                if (data.message === "User logged in") {
                    alert(data.message); // Показываем сообщение
                    // Перенаправление на страницу аккаунта
                    setTimeout(() => {
                        window.location.href = `account.html?username=${encodeURIComponent(username)}`;
                    }, 1500)
                    
                } else {
                    alert(data.error);// Показываем ошибку
                }
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

