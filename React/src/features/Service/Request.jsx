import axios from 'axios';
// const apiClient = axios.create({
//     baseURL: 'http://localhost:5084',
//     headers: {
//       'Content-Type': 'application/json',
//     'Authorization': `Bearer ${token}`
  
//     },
//   });
export const LoginFormRequest = async (DataLogin) => {
    try {
        const response = await axios.post('http://localhost:9000/api/auth/Login', DataLogin,{
            headers: {
                'Content-Type': 'multipart/form-data',
              }
        }); 
        console.log(response.data)
        return response.data;
    } catch (error) {
        console.error("Error adding game:", error.response || error.message);
        return null; 
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
        localStorage.setItem('user', JSON.stringify({
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