// src/api.js
import axios from 'axios';

const API_URL = 'http://localhost:8080/api'; // Замените на ваш URL

export const registerUser = async (userData) => {
    try {
        const response = await axios.post(`${API_URL}/register`, userData);
        return response.data;
    } catch (error) {
        throw error.response.data; // Обработка ошибки
    }
};

export const loginUser = async (userData) => {
    try {
        const response = await axios.post(`${API_URL}/login`, userData);
        return response.data;
    } catch (error) {
        throw error.response.data; // Обработка ошибки
    }
};