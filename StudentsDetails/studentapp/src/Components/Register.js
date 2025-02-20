import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import axios from 'axios';


function Register() {
    const [username, setUsername] = useState('');
    const [password, setPassword] = useState('');
    const navigate = useNavigate();

    const handleRegister = async (e) => {
        e.preventDefault();
        try {
          const res = await axios.post('http://localhost:8080/api/register', {
            username,
            password
          });
          console.log(res.data);
          console.log(res.data.message);
          if (res.data.message==="User registered successfully") {
            navigate('/');
          }
          
          
        } catch (error) {
          console.error('Registration error:', error);
          
        }
      };

    return (
        <div className="register-container">
            <form onSubmit={handleRegister}>
                
                <input
                    type="text"
                    value={username}
                    onChange={(e) => setUsername(e.target.value)} required
                    placeholder="Username"
                />
                <input
                    type="password"
                    value={password}
                    onChange={(e) => setPassword(e.target.value)} required
                    placeholder="Password"
                />
                <button type="submit">Register</button>
            </form>
        </div>
    );
}

export default Register;
