import axios from 'axios';
import { getCookie } from '../../utils/cookies';

const token = getCookie("accesstoken");
const apiClient = axios.create({
    baseURL: 'http://localhost:8000',
    headers: {
        'Content-Type': 'application/json',
        'Authorization': `Bearer ${token}`
  
    },
  });
export const LoginFormRequest = async (DataLogin) => {
    try {
        const response = await apiClient.post('/auth/login', DataLogin); 
        sessionStorage.setItem('access_token', response.data.access_token);
        return response.data;
    } catch (error) {
        console.error("Error adding game:", error.response || error.message);
        return error; 
    }
}

export const GoogleLoginRequest = async (accessToken) => {
    try {
        const userInfoResponse = await axios.get('https://www.googleapis.com/oauth2/v3/userinfo', {
            headers: {
                Authorization: `Bearer ${accessToken}`
            }
        });
        
        const userInfo = userInfoResponse.data;
        
        // Store user info in localStorage
        sessionStorage.setItem('user', JSON.stringify({
            email: userInfo.email,
            name: userInfo.name,
            picture: userInfo.picture,
            googleId: userInfo.sub,
            accessToken: accessToken
        }));

        return userInfo;
    } catch (error) {
        console.error("Google login error:", error);
        throw error;
    }
}