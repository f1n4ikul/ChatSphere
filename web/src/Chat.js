import React, { useEffect, useState } from 'react';

const Chat = ({ userId }) => {
    const [chats, setChats] = useState([]);
    const [messages, setMessages] = useState([]);
    const [newMessage, setNewMessage] = useState('');
    const [newChatUsername, setNewChatUsername] = useState(''); // Состояние для нового имени пользователя
    const [activeChatId, setActiveChatId] = useState(null); // Для отслеживания активного чата

    const fetchChats = async () => {
        try {
            const response = await fetch('http://localhost:8080/api/chats');
            if (!response.ok) {
                throw new Error('Ошибка при загрузке чатов');
            }
            const data = await response.json();
            if (Array.isArray(data)) {
                setChats(data);
            } else {
                console.error('Полученные данные не являются массивом:', data);
            }
        } catch (error) {
            console.error('Ошибка:', error);
        }
    };

    useEffect(() => {
        fetchChats(); // Вызов функции при монтировании компонента
    }, []);

    const loadMessages = async (chatId) => {
        const response = await fetch(`http://localhost:8080/api/chats/${chatId}/messages`);
        const data = await response.json();
        setMessages(data);
        setActiveChatId(chatId); // Установите активный чат
    };

    const sendMessage = async (chatId) => {
        const response = await fetch(`http://localhost:8080/api/chats/${chatId}/messages`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({
                sender_id: userId,
                text: newMessage,
            }),
        });

        if (response.ok) {
            setNewMessage('');
            loadMessages(chatId); // Перезагрузить сообщения после отправки
        }
    };

    const createChat = async () => {
        try {
            const response = await fetch('http://localhost:8080/api/chats', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({
                    user1_id: userId, // Имя пользователя для нового чата
                    user2_id: newChatUsername, // Ваш ID пользователя
                }),
            });

            if (response.ok) {
                setNewChatUsername(''); // Очистить поле ввода
                fetchChats(); // Обновить список чатов
            } else {
                console.error('Ошибка при создании чата');
            }
        } catch (error) {
            console.error('Ошибка:', error);
        }
    };

    if (!Array.isArray(chats)) {
        return <div>Ошибка: ожидается массив чатов</div>;
    }

    return (
        <div>
            <h1>Чаты</h1>
            <ul>
                {chats.map(chat => (
                    <li key={chat.id} onClick={() => loadMessages(chat.id)}>
                        Чат с {chat.user2_id} {/* Предполагается, что у вас есть информация о пользователе */}
                    </li>
                ))}
            </ul>
            <div>
                {messages.map(message => (
                    <p key={message.id}>{message.text}</p>
                ))}
                {activeChatId && (
                    <div>
                        <input
                            type="text"
                            value={newMessage}
                            onChange={(e) => setNewMessage(e.target.value)}
                        />
                        <button onClick={() => sendMessage(activeChatId)}>Отправить</button>
                    </div>
                )}
            </div>
            <div>
                <input
                    type="text"
                    value={newChatUsername}
                    onChange={(e) => setNewChatUsername(e.target.value)}
                    placeholder="Введите имя пользователя"
                />
                <button onClick={createChat}>Создать чат</button>
            </div>
        </div>
    );
};

export default Chat;