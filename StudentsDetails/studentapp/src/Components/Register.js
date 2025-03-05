import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { QRCodeCanvas } from "qrcode.react"; 
import { FaEye, FaEyeSlash } from "react-icons/fa";
import axios from 'axios';


function Register() {
    const [username, setUsername] = useState('');
    const [password, setPassword] = useState('');
    const [confirmPassword, setConfirmPassword] = useState("");
    const [qrCode, setQrCode] = useState("");
  const [secret, setSecret] = useState("");
  const [otp, setOtp] = useState("");
  const[error,setError]=useState("");
  const [showPassword, setShowPassword] = useState(false); // State for password visibility
  const [showConfirmPassword, setShowConfirmPassword] = useState(false);
    const navigate = useNavigate();

    const handleRegister = async (e) => {
        e.preventDefault();
        if (password !== confirmPassword) {
          setError("Passwords do not match");
          return;
        }
        try {
          const res = await axios.post('http://localhost:8080/api/register', {
            username,
            password,
            confirmPassword
          });
          console.log(res.data);
          console.log(res.data.message);
          if (res.data.message==="User registered successfully") {
            alert("User registered! Scan the QR code for 2FA.");
            setQrCode(res.data.qr_code);
            setSecret(res.data.totp_secret);
            setError("");
            // navigate('/');
          }
          
          
        } catch (error) {
          console.error('Registration error:', error);
          
        }
      };
      const handleVerify = async () => {
        try {
          const res = await axios.post("http://localhost:8080/api/verify-2fa", { username, totp: otp });
          alert("2FA enabled! You can now log in.");
          navigate("/login");
        } catch (err) {
          alert("Invalid OTP. Try again.");
        }
      };
    return (
        <div className="register-container">
          <h2>Register</h2>
          {error && <p style={{ color: "red" }}>{error}</p>}
            <form onSubmit={handleRegister}>
            <div  className="input-container">
                <input 
                    type="text"
                    value={username}
                    onChange={(e) => setUsername(e.target.value)} required
                    placeholder="Username"
                />
                </div>
                <div  className="input-container">
                  <input
                    type={showPassword ? "text" : "password"}
                    placeholder="Password"
                    value={password}
                    onChange={(e) => setPassword(e.target.value)}
                  />
                  <span className='password-eye'
                      onClick={() => setShowPassword(!showPassword)}>
                  {showPassword ? <FaEyeSlash /> : <FaEye />}
                  </span>
                </div>
                {/* Confirm Password field with toggle */}
                <div className="input-container">
                  <input
                    type={showConfirmPassword ? "text" : "password"}
                    placeholder="Re-enter Password"
                    value={confirmPassword}
                    onChange={(e) => setConfirmPassword(e.target.value)}
                    
                  />
                  <span onClick={() => setShowConfirmPassword(!showConfirmPassword)}
                  >
                    {showConfirmPassword ? <FaEyeSlash /> : <FaEye />}
                  </span>
                </div>
                <button type="submit">Register</button>
            </form>
            {qrCode && (
        <div className='qr-container'>
          <h3>Scan QR Code in Google Authenticator</h3>
          <QRCodeCanvas value={qrCode} size={200} />
          <input className='qr-text' type="text" placeholder="Enter OTP" onChange={(e) => setOtp(e.target.value)} />
          <button onClick={handleVerify}>Verify & Enable 2FA</button>
        </div>
      )}
        </div>
    );
}

export default Register;
