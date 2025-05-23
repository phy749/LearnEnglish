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
        return response.data;
    } catch (error) {
        console.error("Error adding game:", error.response || error.message);
        return null; 
    }
}