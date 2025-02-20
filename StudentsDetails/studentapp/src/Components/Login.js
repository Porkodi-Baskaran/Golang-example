import React, { useState } from 'react';
import { useNavigate, useHistory } from 'react-router-dom';
import axios from 'axios';
import '../styles.css';
import GoogleLoginButton from './GoogleLoginButton';

function Login({ onLogin }) {
    const [username, setUsername] = useState('');
    const [password, setPassword] = useState('');

    const navigate=useNavigate();

    const handleSubmit = async (e) => {
        e.preventDefault();
        try {
            const res = await axios.post('http://localhost:8080/api/login', {
              username,
              password
        }, {
            withCredentials: true // This allows cookies to be sent and received
          });
          if (res.status === 200) {
            console.log("response:",res.data);
            navigate('/students');
          } else {
            console.error('Login failed:', res.data);
          }
        } catch (error) {
          console.error('Login error:', error);
        }
      };
    const handleRegisterRedirect = () => {
        navigate('/register');
    };
    
    return (
        <div className="login-container">
        <form onSubmit={handleSubmit}>
            <input
                type="text"
                value={username}
                onChange={(e) => setUsername(e.target.value)}
                placeholder="Username"
            />
            <input
                type="password"
                value={password}
                onChange={(e) => setPassword(e.target.value)}
                placeholder="Password"
            />
            <button type="submit">Login</button>
            <button type="button" onClick={handleRegisterRedirect}>
                Register
            </button>
        </form>
        {/* <hr />
      <GoogleLoginButton /> */}
       
        </div>
    );
}

export default Login;


