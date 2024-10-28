document.addEventListener('DOMContentLoaded', () => {
    const registerButton = document.getElementById('register-button');
    const loginButton = document.getElementById('login-button');
    const sendButton = document.getElementById('send');
    const messageInput = document.getElementById('message');
    const messagesDiv = document.getElementById('messages');
    const chatDiv = document.getElementById('chat');

    // Регистрация пользователя
    // registerButton.onclick = function() {
    //     const username = document.getElementById('register-username').value;
    //     const password = document.getElementById('register-password').value;

    //     fetch('http://localhost:8080/register', {
    //         method: 'POST',
    //         headers: {
    //             'Content-Type': 'application/x-www-form-urlencoded',
    //         },
    //         body: new URLSearchParams({
    //             username: username,
    //             password: password
    //         })
    //     })
    //     .then(response => response.json())
    //     .then(data => {
    //         alert(data.message); // Показываем сообщение
    //     })
    //     .catch(error => console.error('Error:', error));
    // };


        registerButton.onclick = function() {
        const username = document.getElementById('register-email').value;
        const password = document.getElementById('register-password').value;
        const email = document.getElementById('register-username').value;
            

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
            window.location.href = `account.html?username=${encodeURIComponent(username)}`;
        })
        .catch(error => console.error('Error:', error));
    };





    // Вход пользователя
    // loginButton.onclick = function() {
    //     const username = document.getElementById('login-username').value;
    //     const password = document.getElementById('login-password').value;

    //     fetch('http://localhost:8080/login', {
    //         method: 'POST',
    //         headers: {
    //             'Content-Type': 'application/x-www-form-urlencoded',
    //         },
    //         body: new URLSearchParams({
    //             username: username,
    //             password: password
    //         })
    //     })
    //     .then(response => response.json())
    //     .then(data => {
    //         if (data.message === "User logged in") {
    //             alert(data.message); // Показываем сообщение
    //             chatDiv.style.display = 'block'; // Показываем чат
    //             connectToChat(username); // Подключаемся к чату
    //         } else {
    //             alert(data.error); // Показываем ошибку
    //         }
    //     })
    //     .catch(error => console.error('Error:', error));
    // };


        loginButton.onclick = function() {
            const email = document.getElementById('login-email').value;
            const password = document.getElementById('login-password').value;

        fetch('http://localhost:8080/login', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/x-www-form-urlencoded',
            },
            body: new URLSearchParams({
                email: email,
                password: password
            })
        })
        .then(response => response.json())
        .then(data => {
            if (data.message === "User logged in") {
                alert(data.message); // Показываем сообщение
                // Перенаправление на страницу аккаунта
                window.location.href = `account.html?username=${encodeURIComponent(username)}`;
            } else {
                alert(data.error); // Показываем ошибку
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

