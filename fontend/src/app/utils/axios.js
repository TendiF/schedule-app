import axios from 'axios'
let headers = {
}

let token = localStorage.getItem('token')

if(token){
    headers.Authorization = `Bearer ${token}`
}

let axiosInstance =  axios.create({
    baseURL: process.env.REACT_APP_API,
    headers
});

let errorTimeout = null

axiosInstance
.interceptors
.response
.use( 
    response => response, 
    error => {
    if(error.response && error.response.status === 401){
        window.location = '/login'
    }
    if (!error.response) {
        // network error
        if(errorTimeout) clearTimeout(errorTimeout)
      }
    return Promise.reject(error);
});

export default axiosInstance
