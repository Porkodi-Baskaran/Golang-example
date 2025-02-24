import React from 'react';
import axios from 'axios';
import { useNavigate } from 'react-router-dom';


// const GoogleLoginButton = () => {
//   const handleGoogleLogin = () => {
//     window.location.href = 'http://localhost:8080/api/auth/google/login'; // Redirect to backend
//   };

//   return (
//     <button onClick={handleGoogleLogin} className="google-login">
//       Sign In With Google
//     </button>
//   );
// };

// export default GoogleLoginButton;

const GoogleLoginButton = () => {
  const navigate = useNavigate();

  const handleGoogleLogin = async () => {
    try {
      window.location.href = 'http://localhost:8080/api/auth/google/login';
    
    } catch (error) {
      console.error('Google login error:', error);
    }

  };
  

  return (
    <button onClick={handleGoogleLogin} className="google-login">
      Sign In With Google
    </button>
  );
};

export default GoogleLoginButton;

// 
// 
// import React, { useEffect } from 'react';
// import { useNavigate } from 'react-router-dom';
// import axios from 'axios';

// const GoogleLoginButton = () => {
//   const navigate = useNavigate();

//   useEffect(() => {
//     /* global google */
//     google.accounts.id.initialize({
//       client_id: '507015603058-e86p59t5irp55eq4f9btgbflg5n4gjfb.apps.googleusercontent.com',
//       callback: handleCredentialResponse,
//     });
//     google.accounts.id.renderButton(
//       document.getElementById('googleSignInButton'),
//       { theme: 'outline', size: 'large' }
//     );
//   }, []);

//   const handleCredentialResponse = async (response) => {
//     console.log("Response:",response);
//     try {
//       const res = await axios.post('http://localhost:8080/api/auth/google/callback', {
//         id_token: response.credential,
//       }, {
//         withCredentials: true, // This allows cookies to be sent and received
//       });
//       console.log(res.data);
//       navigate('/students');
//     } catch (error) {
//       console.error('Google login error:', error);
//     }
//   };

//   return (
//     <div>
//       <div id="googleSignInButton"></div>
//     </div>
//   );
// };

// export default GoogleLoginButton;
