import axios from "axios";
axios.defaults.baseURL = "";
axios.interceptors.request.use(
    config => {
        const token = localStorage.getItem("token");
        if (token) {
            config.headers['token'] = token;
        }
        return config;
    },
    error => {
        return Promise.reject(error);
    }
)
axios.interceptors.response.use(
    config => {
        if (config.status === 200) {
            return config.data;
        }
    },
    error => {
        error.message = 'overtime'
        return Promise.reject(error)
    }
)

export const Token = ''

export default axios;
