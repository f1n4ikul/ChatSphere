// src/App.js
import React from 'react';
import { BrowserRouter as Router, Route, Routes, useNavigate } from 'react-router-dom';
import Register from './Register';
import Login from './Login';
import Main from './Main';
import Account from './Account';
import Chat from './Chat';

// import  {useNavigate} from 'react-router-dom';


// const Nav = () => {
//     const navigate = useNavigate();

//     const handleRegister = async (event) => {
//         if (response.ok) {
//             navigate('/account');
//         }
//     }

//     const handleLogin = async (event) => {
//         if (response.ok) {
//             navigate('/account');
//         }

//     }
// }


function App() {
    return (
        <Router>
            <Routes>
                <Route path="/register" element={<Register />} />
                <Route path="/login" element={<Login />} />
                <Route path="/account" element={<Account />} />
                <Route path="/chats" element={<Chat />} />
                <Route path="/" element={<Main />} />
                {/* Добавьте другие маршруты по мере необходимости */}
            </Routes>
        </Router>
    );
}

export default App;
