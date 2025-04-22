import axios from "axios";
import { useState } from "react";
import { useNavigate } from "react-router-dom";

const LoginJWT = () => {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [totp, setTotp] = useState("");
  const navigate = useNavigate();

  const handleLogin = async () => {
    try {
      const loginResponse = await axios.post("http://localhost:8080/api/jwt/jwt-login", {
        username,
        password,
        totp_secret: totp,
      });

      const { token } = loginResponse.data;
      localStorage.setItem("token", token);
      console.log("Login successful. Token stored.");

      // Call protected route
      await axios.get("http://localhost:8080/api/jwt/jwt-protected", {
        headers: {
          Authorization: token,
        },
      });

      console.log("Protected access granted. Redirecting...");
      navigate("/students");

    } catch (error) {
      console.error("Error:", error.response ? error.response.data : error.message);
      alert("Login failed. Check credentials and TOTP.");
    }
  };

  return (
    <div>
      <h2>Login</h2>
      <input type="text" placeholder="Username" onChange={(e) => setUsername(e.target.value)} />
      <input type="password" placeholder="Password" onChange={(e) => setPassword(e.target.value)} />
      <input type="text" placeholder="TOTP Code" onChange={(e) => setTotp(e.target.value)} />
      <button onClick={handleLogin}>LoginJWT</button>
    </div>
  );
};

export default LoginJWT;
