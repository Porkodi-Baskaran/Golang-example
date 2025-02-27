import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import axios from 'axios';
import '../styles.css';
import GoogleLoginButton from './GoogleLoginButton';

function Login({ onLogin }) {
    const [username, setUsername] = useState('');
    const [password, setPassword] = useState('');


    const navigate=useNavigate();

    const handleSubmit = async (e) => {
        e.preventDefault();

        if (!username || !password) {
            console.error('Username and password are required');
           alert("Username and password are required")
            return;
        }

        try {
            // await axios.post('http://localhost:8080/api/logout', {}, { withCredentials: true });
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
      <div className="login-page">
        <div className="login-container">
        <form onSubmit={handleSubmit} onKeyDown={(e) => e.key === 'Enter' && e.preventDefault()}>
            <input
                type="text"
                value={username}
                onChange={(e) => setUsername(e.target.value)}
                placeholder="Username"
                autoComplete="off"
            />
            <input
                type="password"
                value={password}
                onChange={(e) => setPassword(e.target.value)}
                placeholder="Password"
                autoComplete="off"
            />
            <button type="submit">Login</button>
            
            <button type="button" onClick={handleRegisterRedirect}>
                Register
            </button>
        </form>
        <br></br>
         <hr />
         {/* <GoogleLoginButton /> */}
      <GoogleLoginButton> Sign In With Google </GoogleLoginButton> 
       
        </div>
        </div>
    );
}

export default Login;


