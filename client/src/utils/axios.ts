import axios from "axios";

const instance = axios.create({
    baseURL: "/api",
    headers: {
        "Accepts": "application/json",
        "Access-Control-Allow-Origin": "*"
    }
});

export default instance
