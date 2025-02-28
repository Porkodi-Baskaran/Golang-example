import React, { useState, useEffect } from "react";
import axios from "axios";
import { QRCodeCanvas } from "qrcode.react"; // Corrected import
import { useNavigate, useParams } from "react-router-dom";

function EnableTOTP() {
  const { username } = useParams();
  const [qrCode, setQrCode] = useState("");
  const [secret, setSecret] = useState("");
  const [otp, setOtp] = useState("");
  const navigate = useNavigate();

  useEffect(() => {
    const fetchQRCode = async () => {
    //   const username = localStorage.getItem("user");
      if (!username) {
        alert("Please login first.");
        navigate("/login");
        return;
      }

      try {
        const res = await axios.get(`http://localhost:8080/api/enable-2fa?username=${username}`);
        setQrCode(res.data.qrCode);
        setSecret(res.data.secret);
      } catch (err) {
        alert("Error fetching QR Code.");
      }
    };

    fetchQRCode();
  }, []);

  const handleVerify = async () => {
    try {
    //   const username = localStorage.getItem("user");
      const res = await axios.post("http://localhost:8080/api/verify-2fa", { username, secret, totp: otp });
 console.log(res)
      if (res.status === 200) {
        alert("2FA enabled successfully!");
        navigate("/login");
      }
    } catch (err) {
      alert("Invalid OTP. Try again.");
    }
  };

  return (
    <div className="enable-2fa-container">
      <h2>Enable Google Authenticator</h2>
      {qrCode && <QRCodeCanvas value={secret} />} {/* Fixed QR Code rendering */}
      <p>Scan the QR Code in Google Authenticator and enter the OTP below:</p>
      <input type="text" placeholder="Enter OTP" value={otp} onChange={(e) => setOtp(e.target.value)} />
      <button onClick={handleVerify}>Verify & Enable 2FA</button>
    </div>
  );
}

export default EnableTOTP;
