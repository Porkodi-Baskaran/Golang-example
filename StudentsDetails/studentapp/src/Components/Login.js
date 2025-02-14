import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import '../styles.css';

function Login({ onLogin }) {
    const [username, setUsername] = useState('');
    const [password, setPassword] = useState('');

    const navigate=useNavigate();

    const handleSubmit = async (e) => {
        e.preventDefault();
        const response = await fetch('http://localhost:8080/login', {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ username, password }),
        });

        if (response.ok) {
            document.cookie = "token=your-authentication-token; path=/";
            onLogin(true);
        } else {
            alert('Login failed');
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
        </div>
    );
}

export default Login;





// import React, { useState } from "react";
// import axios from "axios";

// const Login = () => {
//     const [username, setUsername] = useState("");
//     const [password, setPassword] = useState("");

//     const handleLogin = async () => {
//         try {
//             const response = await axios.post("http://localhost:8080/login", {
//                 username,
//                 password,
//             });
//             console.log(response.data);
//             // Handle successful login
//         } catch (error) {
//             console.error("Error logging in:", error);
//         }
//     };

//     return (
//         <div>
//             <h2>Login</h2>
//             <input
//                 type="text"
//                 placeholder="Username"
//                 value={username}
//                 onChange={(e) => setUsername(e.target.value)}
//             />
//             <input
//                 type="password"
//                 placeholder="Password"
//                 value={password}
//                 onChange={(e) => setPassword(e.target.value)}
//             />
//             <button onClick={handleLogin}>Login</button>
//         </div>
//     );
// };

// export default Login;



// // import React, { useState } from 'react';
// // import axios from 'axios';

// // function Login({ setToken, setLoggedIn }) {
// //     const [username, setUsername] = useState('');
// //     const [password, setPassword] = useState('');

// //     const handleLogin = async (e) => {
// //         e.preventDefault();
// //         try {
// //             const response = await axios.post('http://localhost:8080/login', {
// //                 username,
// //                 password,
// //             });
// //             if (response.status === 200) {
// //                 setToken(response.data.token);
// //                 setLoggedIn(true);
// //             }
// //         } catch (error) {
// //             console.error("Login failed:", error);
// //         }
// //     };

// //     return (
// //         <form onSubmit={handleLogin}>
// //             <div>
// //                 <label>Username:</label>
// //                 <input
// //                     type="text"
// //                     value={username}
// //                     onChange={(e) => setUsername(e.target.value)}
// //                 />
// //             </div>
// //             <div>
// //                 <label>Password:</label>
// //                 <input
// //                     type="password"
// //                     value={password}
// //                     onChange={(e) => setPassword(e.target.value)}
// //                 />
// //             </div>
// //             <button type="submit">Login</button>
// //         </form>
// //     );
// // }

// // export default Login;




// // // // src/components/Login.js
// // // import React, { useState } from 'react';
// // // import axios from 'axios';
// // // // import Cookies from 'js-cookie';

// // // function Login({ setToken, setLoggedIn }) {
// // //     const [username, setUsername] = useState('');
// // //     const [password, setPassword] = useState('');

// // //     const handleLogin = async (e) => {
// // //         e.preventDefault();
// // //         try {
// // //             const response = await axios.post('http://localhost:8080/login', {
// // //                 username,
// // //                 password,
// // //             });
// // //             if (response.status === 200) {
// // //                 setToken(response.data.token);
// // //                 setLoggedIn(true);
// // //             }
// // //         } catch (error) {
// // //             console.error("Login failed:", error);
// // //         }
// // //     };

// // //     return (
// // //         <form onSubmit={handleLogin}>
// // //             <div>
// // //                 <label>Username:</label>
// // //                 <input
// // //                     type="text"
// // //                     value={username}
// // //                     onChange={(e) => setUsername(e.target.value)}
// // //                 />
// // //             </div>
// // //             <div>
// // //                 <label>Password:</label>
// // //                 <input
// // //                     type="password"
// // //                     value={password}
// // //                     onChange={(e) => setPassword(e.target.value)}
// // //                 />
// // //             </div>
// // //             <button type="submit">Login</button>
// // //         </form>
// // //     );
// // // }

// // // export default Login;


// // // // const Login = () => {
// // // //     const [username, setUsername] = useState('');
// // // //     const [password, setPassword] = useState('');

// // // //     const handleSubmit = async (e) => {
// // // //         e.preventDefault();
// // // //         const credentials = { username, password };
// // // //         try {
// // // //             const response = await axios.post('https://localhost:8080/login', credentials);
// // // //             console.log('Login successfully:', response.data);
// // // //             const token = Cookies.get('token');
// // // //             console.log('Token from cookie:', token);
// // // //         } catch (error) {
// // // //             console.error('Error logging in:', error);
// // // //         }
// // // //     };

// // // //     return (
// // // //         <form onSubmit={handleSubmit}>
// // // //             <input type="text" value={username} onChange={(e) => setUsername(e.target.value)} placeholder="Username" />
// // // //             <input type="password" value={password} onChange={(e) => setPassword(e.target.value)} placeholder="Password" />
// // // //             <button type="submit">Login</button>
// // // //         </form>
// // // //     );
// // // // };

// // // // export default Login;
